<!-- template:begin:header -->
<!-- do not edit anything in this "template" block, its auto-generated -->

<p align="center">go-bogon -- Go package to check if an IP address is a bogon (internal) IP address</p>
<p align="center">
  <a href="https://github.com/lrstanley/go-bogon/tags">
    <img title="Latest Semver Tag" src="https://img.shields.io/github/v/tag/lrstanley/go-bogon?style=flat-square">
  </a>
  <a href="https://github.com/lrstanley/go-bogon/commits/master">
    <img title="Last commit" src="https://img.shields.io/github/last-commit/lrstanley/go-bogon?style=flat-square">
  </a>


  <a href="https://github.com/lrstanley/go-bogon/actions?query=workflow%3Atest+event%3Apush">
    <img title="GitHub Workflow Status (test @ master)" src="https://img.shields.io/github/actions/workflow/status/lrstanley/go-bogon/test.yml?branch=master&label=test&style=flat-square">
  </a>



  <a href="https://codecov.io/gh/lrstanley/go-bogon">
    <img title="Code Coverage" src="https://img.shields.io/codecov/c/github/lrstanley/go-bogon/master?style=flat-square">
  </a>

  <a href="https://pkg.go.dev/github.com/lrstanley/go-bogon">
    <img title="Go Documentation" src="https://pkg.go.dev/badge/github.com/lrstanley/go-bogon?style=flat-square">
  </a>
  <a href="https://goreportcard.com/report/github.com/lrstanley/go-bogon">
    <img title="Go Report Card" src="https://goreportcard.com/badge/github.com/lrstanley/go-bogon?style=flat-square">
  </a>
</p>
<p align="center">
  <a href="https://github.com/lrstanley/go-bogon/issues?q=is:open+is:issue+label:bug">
    <img title="Bug reports" src="https://img.shields.io/github/issues/lrstanley/go-bogon/bug?label=issues&style=flat-square">
  </a>
  <a href="https://github.com/lrstanley/go-bogon/issues?q=is:open+is:issue+label:enhancement">
    <img title="Feature requests" src="https://img.shields.io/github/issues/lrstanley/go-bogon/enhancement?label=feature%20requests&style=flat-square">
  </a>
  <a href="https://github.com/lrstanley/go-bogon/pulls">
    <img title="Open Pull Requests" src="https://img.shields.io/github/issues-pr/lrstanley/go-bogon?label=prs&style=flat-square">
  </a>
  <a href="https://github.com/lrstanley/go-bogon/discussions/new?category=q-a">
    <img title="Ask a Question" src="https://img.shields.io/badge/support-ask_a_question!-blue?style=flat-square">
  </a>
  <a href="https://liam.sh/chat"><img src="https://img.shields.io/badge/discord-bytecord-blue.svg?style=flat-square" title="Discord Chat"></a>
</p>
<!-- template:end:header -->

<!-- template:begin:toc -->
<!-- do not edit anything in this "template" block, its auto-generated -->
## :link: Table of Contents

  - [What](#what)
  - [Usage](#usage)
  - [Support &amp; Assistance](#raising_hand_man-support--assistance)
  - [Contributing](#handshake-contributing)
  - [License](#balance_scale-license)
<!-- template:end:toc -->

## What

go-bogon is a basic package to check if an IP is within the bogon IP ranges or not.

- No-frills check for the common default bogon IP ranges: `bogon.Is("127.0.0.1")`
- You can supply your own CIDR's if you wish: `bogon.New([]string{"0.0.0.0/8"})`


## Usage

<!-- template:begin:goget -->
<!-- do not edit anything in this "template" block, its auto-generated -->
```console
go get -u github.com/lrstanley/go-bogon@latest
```
<!-- template:end:goget -->

<!-- template:begin:support -->
<!-- do not edit anything in this "template" block, its auto-generated -->
## :raising_hand_man: Support & Assistance

* :heart: Please review the [Code of Conduct](.github/CODE_OF_CONDUCT.md) for
     guidelines on ensuring everyone has the best experience interacting with
     the community.
* :raising_hand_man: Take a look at the [support](.github/SUPPORT.md) document on
     guidelines for tips on how to ask the right questions.
* :lady_beetle: For all features/bugs/issues/questions/etc, [head over here](https://github.com/lrstanley/go-bogon/issues/new/choose).
<!-- template:end:support -->

<!-- template:begin:contributing -->
<!-- do not edit anything in this "template" block, its auto-generated -->
## :handshake: Contributing

* :heart: Please review the [Code of Conduct](.github/CODE_OF_CONDUCT.md) for guidelines
     on ensuring everyone has the best experience interacting with the
    community.
* :clipboard: Please review the [contributing](.github/CONTRIBUTING.md) doc for submitting
     issues/a guide on submitting pull requests and helping out.
* :old_key: For anything security related, please review this repositories [security policy](https://github.com/lrstanley/go-bogon/security/policy).
<!-- template:end:contributing -->

<!-- template:begin:license -->
<!-- do not edit anything in this "template" block, its auto-generated -->
## :balance_scale: License

```
MIT License

Copyright (c) 2017 Liam Stanley <liam@liam.sh>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

_Also located [here](LICENSE)_
<!-- template:end:license -->
