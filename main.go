// Copyright 2018 Christopher Eaton
// Use of this source code is governed by the Mozilla Public License v2
// You can find the license text in the LICENSE file, or online at:
// https://www.mozilla.org/en-US/MPL/2.0/

// This app looks for Stellar keypairs with any of the given suffixes (command-line args).
package main

import (
	"fmt"
	"os"
	"strings"
	"sync/atomic"

	"github.com/stellar/go/keypair"
)

var count uint32 = 0

func HasAnySuffix(input string, words []string) bool {
	for x := 0; x < len(words); x++ {
		if strings.HasSuffix(input, strings.ToUpper(words[x])) {
			return true
		}
	}
	return false
}

func scan(found chan bool, instance int, words []string) {
	for true {
		kp, err := keypair.Random()
		if err != nil {
			panic(err)
		}
		pubKey := kp.Address()
		if HasAnySuffix(pubKey, words) {
			fmt.Printf("Public Key: %s", pubKey)
			fmt.Println()
			fmt.Printf("Private Key: %s", kp.Seed())
			fmt.Println()
			found <- true
			break
		}
		atomic.AddUint32(&count, 1)
		countTotal := atomic.LoadUint32(&count)
		if countTotal%100000 == 0 {
			fmt.Printf("Generated %d random keys...\n", countTotal)
		}
		select {
		case <-found:
			break
		default:
		}
	}
}

func manualExit(found chan bool) {
	//cancel automatically
	var input string
	fmt.Scanln(&input)
	found <- true
	fmt.Println("Ended.")
}

func main() {
	found := make(chan bool, 1)
	words := os.Args[1:]
	if len(words) > 0 {

		for x := 1; x <= 4; x++ {
			go scan(found, x, words)
		}
		go manualExit(found)
		<-found
	} else {
		fmt.Println("No words specified in command line arguments.")
	}
}
