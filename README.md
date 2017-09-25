# go-bogon [![godoc](https://godoc.org/github.com/lrstanley/go-bogon?status.png)](https://godoc.org/github.com/lrstanley/go-bogon) [![goreport](https://goreportcard.com/badge/github.com/lrstanley/go-bogon)](https://goreportcard.com/report/github.com/lrstanley/go-bogon) [![build](https://travis-ci.org/lrstanley/go-bogon.svg?branch=master)](https://travis-ci.org/lrstanley/go-bogon) [![gocover](http://gocover.io/_badge/github.com/lrstanley/go-bogon)](https://gocover.io/github.com/lrstanley/go-bogon)

go-bogon is a basic package to check if an IP is within the bogon IP ranges or not.

- No-frills check for the common default bogon IP ranges: `bogon.Is("127.0.0.1")`
- You can supply your own CIDR's if you wish: `bogon.New([]string{"0.0.0.0/8"})`

## Contributing

Please review the [CONTRIBUTING](https://github.com/lrstanley/go-bogon/blob/master/CONTRIBUTING.md)
doc for submitting issues/a guide on submitting pull requests and helping out.

## License

    LICENSE: The MIT License (MIT)
    Copyright (c) 2017 Liam Stanley <me@liamstanley.io>

    Permission is hereby granted, free of charge, to any person obtaining a copy
    of this software and associated documentation files (the "Software"), to deal
    in the Software without restriction, including without limitation the rights
    to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
    copies of the Software, and to permit persons to whom the Software is
    furnished to do so, subject to the following conditions:

    The above copyright notice and this permission notice shall be included in
    all copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
    AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
    OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
    SOFTWARE.
