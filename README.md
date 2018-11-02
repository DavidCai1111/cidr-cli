# cidr-cli
[![Build Status](https://travis-ci.org/DavidCai1993/cidr-cli.svg?branch=master)](https://travis-ci.org/DavidCai1993/cidr-cli)
[![Coverage Status](https://coveralls.io/repos/github/DavidCai1993/cidr-cli/badge.svg?branch=master)](https://coveralls.io/github/DavidCai1993/cidr-cli?branch=master)

A tiny cli tool to get useful IP address information from a CIDR.

## Installation

```
go get -u -v github.com/DavidCai1993/cidr-cli
```

## Usage

```
cidr-cli 192.168.23.35/21
// output:
// {"IPAvailable":2046,"netID":"192.168.16.0","mask":"fffff800"}
```
