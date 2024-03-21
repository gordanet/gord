
Gord
====

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/gordanet/gord)

Gord is the reference full node GOR implementation written in Go (golang).

## What is Gor Network

GOR is decentralized cryptocurrency, that is making waves in the world of BlockChain.
This innovative project is based on the Proof of Work consensus algorithm and uses the Blake3 hashing function to ensure the security and efficiency of its network.

## Requirements

Go 1.18 or later.

## Installation

#### Build from Source

- Install Go according to the installation instructions here:
  http://golang.org/doc/install

- Ensure Go was installed properly and is a supported version:

```bash
$ go version
```

- Run the following commands to obtain and install gord including all dependencies:

```bash
$ git clone https://github.com/gordanet/gord
$ cd gord
$ [go install . ./cmd/...]
$ build.sh

```

- Gord (and utilities) should now be installed in `$(go env GOPATH)/bin`. If you did
  not already add the bin directory to your system path during Go installation,
  you are encouraged to do so now.

# postgres

```bash
$ sudo apt install postgresql postgresql-contrib
$ sudo -u postgres psql
$ CREATE ROLE GOR WITH LOGIN ENCRYPTED PASSWORD '1';
$ CREATE DATABASE gor OWNER GOR;
$ Quit psql with \q
```




## Getting Started

Gord has several configuration options available to tweak how it runs, but all
of the basic operations work with zero configuration.

```bash
$ cd ~/go/bin
$ gord --utxoindex
```


## Stratum server
$ git clone https://github.com/gordanet/gor-bridge
...

## Website
Join our website server using the following link: https://GOR.com/

## Twitter
Join our twitter server using the following link: https://twitter.com/GorCurrency

## Discord
Join our discord server using the following link: https://discord.gg/YNYnNN5Pf2

## Telegram
Join our telegram server using the following link: https://t.me/gorcurrency