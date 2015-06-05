## go-image

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/ferhatelmas/go-image)
[![Build Status](https://travis-ci.org/ferhatelmas/go-image.png?branch=master)](https://travis-ci.org/ferhatelmas/go-image)

> Check if a filepath is an image file.

### Install

```
go get github.com/ferhatelmas/go-image
```

### Usage

```go
import "github.com/ferhatelmas/go-image"

image.Is("src/unicorn.png")
//=> true

image.Is("src/unicorn.go")
//=> false
```

### License

MIT © [ferhat elmas](http://ferhatelmas.com)
