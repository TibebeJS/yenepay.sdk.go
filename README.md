# YenePaySDK - GoLang #
[![Go Reference](https://pkg.go.dev/badge/github.com/TibebeJs/yenepay.sdk.go.svg)](https://pkg.go.dev/github.com/TibebeJs/yenepay.sdk.go) 
![tests](https://github.com/TibebeJS/yenepay.sdk.go/workflows/tests/badge.svg)
[![codecov](https://codecov.io/gh/TibebeJS/yenepay.sdk.go/branch/main/graph/badge.svg?token=8M2G27NVA5)](https://codecov.io/gh/TibebeJS/yenepay.sdk.go)
![linter](https://github.com/TibebeJS/yenepay.sdk.go/workflows/linter/badge.svg)
![Go Report Card](https://goreportcard.com/badge/github.com/tibebejs/yenepay.sdk.go)

A Go library to integrate YenePay as a Payment method in Go Web Applications.

## Getting Started ##

To add YenePay to your application and start collecting payments, you will first need to register on YenePay as a merchant and get your seller code. You can do that from https://www.yenepay.com/merchant.

Also make sure to enable "sandbox" mode to not incur any charges during developing/testing.

## Installation ##

```
$ go get -u github.com/TibebeJs/yenepay.sdk.go
```

Alternatively, If you are using go-mod to manage your project, simply add yenepay to your `go.mod` file:

```
module github.com/x/y

go 1.15

require (
    github.com/TibebeJs/yenepay.sdk.go
)
```

And simply import in your source code.

## Examples/Demos ##

 - [Sample Shop Site](github.com/TibebeJs/yenepay.sample-shop.go/tree/main/).


## Bugs ##

Bugs or suggestions? Visit the [issue tracker](https://github.com/TibebeJS/yenepay.sdk.go/issues) 

## Contribution

Any contribution is highly appreciated (bug fixes, feature implementation, etc..)

Please check this section to [[Contribute]](CONTRIBUTING.md).


## Deployment ##

> To Be Documented Soon




