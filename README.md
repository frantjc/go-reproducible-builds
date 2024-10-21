# go-reproducible-builds [![CI](https://github.com/frantjc/go-reproducible-builds/actions/workflows/ci.yml/badge.svg?branch=main&event=push)](https://github.com/frantjc/go-reproducible-builds/actions) [![godoc](https://pkg.go.dev/badge/github.com/frantjc/go-reproducible-builds.svg)](https://pkg.go.dev/github.com/frantjc/go-reproducible-builds) [![goreportcard](https://goreportcard.com/badge/github.com/frantjc/go-reproducible-builds)](https://goreportcard.com/report/github.com/frantjc/go-reproducible-builds) ![license](https://shields.io/github/license/frantjc/go-reproducible-builds)

[Go](https://go.dev) module to expose the [reproducible-builds.org](https://reproducible-builds.org/)'s `SOURCE_DATE_EPOCH` environment variable.

## install

```sh
go get github.com/frantjc/go-reproducible-builds
```

## use

```go
import (
    "fmt"

    reproduciblebuilds "github.com/frantjc/go-reproducible-builds"
)

func main() {
    fmt.Println(reproduciblebuilds.SourceDateEpoch)
}
```
