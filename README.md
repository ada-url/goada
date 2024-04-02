# goada: Fast WHATWG URL library in Go
[![Go-CI](https://github.com/ada-url/goada/actions/workflows/ubuntu.yml/badge.svg)](https://github.com/ada-url/goada/actions/workflows/ubuntu.yml) 
[![GoDoc](https://godoc.org/github.com/ada-url/goada?status.svg)](https://godoc.org/github.com/ada-url/goada)

The goada library provides support for the WHATWG URL standard in Go.

## Requirements

- Go 1.21 or better.


### Examples

```Go
url, nil := New("https://	www.GOoglé.com/./path/../path2/")
fmt.Println(url.Href()) // "https://www.xn--googl-fsa.com/path2/"
```


A common use of a URL parser is to take a URL string and normalize it. 
The WHATWG URL specification has been adopted by most browsers.  Other tools, such as the Go runtime, follow the RFC 3986. 
The following table illustrates possible differences in practice (encoding of the host, encoding of the path):

| string source | string value |
|:--------------|:--------------|
| input string | https://www.7-Eleven.com/Home/../Privacy/Montréal |
| ada's normalized string | https://www.xn--7eleven-506c.com/Home/Privacy/Montr%C3%A9al |
| curl 7.87 | https://www.7-Eleven.com/Privacy/Montr%C3%A9al |
| Go runtime (`net/url`) | https://www.7-Eleven.com/Home/../Privacy/Montr%C3%A9al |

The Go runtime (`net/url`) does not normalize hostnames, and it does not process pathnames properly.

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
