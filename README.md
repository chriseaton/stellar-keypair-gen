# Stellar Keypair Generator
This command-line interface (CLI) application randomly generates Stellar compatible key-pairs (a public and private key) and searches for a suffix
of the public key (AKA the "Account ID") matching any of the given arguments. The more characters, the longer it will
take unless you are very lucky. A suffix of 5 characters might take 5 minutes, but 10 characters might take a year.

It's recommended to pass every conceivable suffix you're willing to accept, as it's easier to look for multiple matches
at once than running multiple times.

# Example
To generate and scan for a key-pair with a public key ending in "ABC" or "APPLE", just run:
```
skpg ABC APPLE
```

# Installation
You have two options available to you for installation: Either use our pre-built binaries, or build and install from the source code.

## Option 1: Pre-Built Binaries
You can find the binaries on the Releases page.

## Option 2: Build & Install from Source
Follow the instructions below to build the application.

# Building
This is a straight-forward go application built and tested on Go 1.7 (although earlier versions are likely supported).

1. Clone the repo into your ```GOPATH```.
```
git clone 
cd stellar-keypair-gen
```

2. Get the dependancies.
```
go get -d ./...
```

3. Build
```
go build -o skpg
```
