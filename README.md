# figma
A Golang package for interacting with the [Figma](https://figma.com) APIs

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
