# go-bogon basic library to check if an IP is within the bogon IP ranges or not.

[![GoDoc](https://godoc.org/github.com/lrstanley/go-bogon?status.png)](https://godoc.org/github.com/lrstanley/go-bogon)
[![Go Report Card](https://goreportcard.com/badge/github.com/lrstanley/go-bogon)](https://goreportcard.com/report/github.com/lrstanley/go-bogon)

- No-frills check for the common default bogon IP ranges: `bogon.Is("127.0.0.1")`
- You can supply your own CIDR's if you wish: `bogon.New([]string{"0.0.0.0/8"})`
