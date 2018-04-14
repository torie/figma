# figma
A Golang package for interacting with the [Figma](https://figma.com) APIs

[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/torie/figma) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/torie/figma/master/LICENSE) [![Go Report Card](https://goreportcard.com/badge/github.com/torie/figma)](https://goreportcard.com/report/github.com/torie/figma)

## TODO

- [x] [GET] Files
- [x] [GET] Images
- [x] [POST] Comments
- [x] [GET] Team projects
- [x] [GET] Project files
- [ ] Testing
- [ ] CI integration

## Getting started
Before using this package, you will need a Figma account and a personal access token.

## Installation

```bash
> go get -u github.com/torie/figma
```

## Usage

### Create a Figma client
```go
c := figma.New("access-token")
```

### Get a Figma document
```go
f, err := c.File("document-key")
```

### Render a node as PNG
```go
imgs, err := c.Images("document-key", 2, figma.ImageFormatPNG, "node-id")
```

### Examples
Examples can be found in the [examples folder](examples)
