# go-bogon [![GoDoc](https://godoc.org/github.com/lrstanley/go-bogon?status.png)](https://godoc.org/github.com/lrstanley/go-bogon) [![Go Report Card](https://goreportcard.com/badge/github.com/lrstanley/go-bogon)](https://goreportcard.com/report/github.com/lrstanley/go-bogon) [![Coverage Status](https://coveralls.io/repos/github/lrstanley/go-bogon/badge.svg?branch=master)](https://coveralls.io/github/lrstanley/go-bogon?branch=master) [![Build Status](https://travis-ci.org/lrstanley/go-bogon.svg?branch=master)](https://travis-ci.org/lrstanley/go-bogon)

go-bogon is a basic package to check if an IP is within the bogon IP ranges or not.

- No-frills check for the common default bogon IP ranges: `bogon.Is("127.0.0.1")`
- You can supply your own CIDR's if you wish: `bogon.New([]string{"0.0.0.0/8"})`
