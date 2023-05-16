# goada: Fast WHATWG URL library in Go
[![Go-CI](https://github.com/ada-url/goada/actions/workflows/ubuntu.yml/badge.svg)](https://github.com/ada-url/goada/actions/workflows/ubuntu.yml) 
[![GoDoc](https://godoc.org/github.com/ada-url/goada?status.svg)](https://godoc.org/github.com/ada-url/goada)

The goada library provides support for the WHATWG URL standard in Go.

## Requirements

- Go 1.20 or better.


### Examples

```Go
url, nil := New("https://	www.GOoglé.com")
fmt.Println(url.Href()) // "https://www.xn--googl-fsa.com/"
```

The standard `net/url` `Parse` function would refuse to parse the URL `"https://	www.GOoglé.com"` because it 
contains a tabulation character. Even if we remove the tabulation character, it would still parse it to an incorrect 
string as per the WHATGL URL standard (`https://www.GOogl%C3%A9.com`).

### Usage

```Go
import (
   "github.com/ada-url/goada"
   "fmt"
)

url, err := goada.New("https://www.GOogle.com")
if err != nil {
    t.Error("Expected no error")
}
fmt.Println(url.Href()) // prints https://www.google.com/
url.SetProtocol("http:")
url.SetHash("goada")
fmt.Println(url.Hash()) // prints #goada
fmt.Println(url.Href()) // prints http://www.google.com/#goada
```
