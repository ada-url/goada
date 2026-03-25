# goada: Fast WHATWG URL library in Go
[![Go-CI](https://github.com/ada-url/goada/actions/workflows/ubuntu.yml/badge.svg)](https://github.com/ada-url/goada/actions/workflows/ubuntu.yml) 
[![GoDoc](https://godoc.org/github.com/ada-url/goada?status.svg)](https://godoc.org/github.com/ada-url/goada)

The goada library provides support for the WHATWG URL standard in Go.

## Requirements

- Go 1.24 or better.


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

## Performance

Benchmarks comparing URL parsing performance using the top 100k URLs dataset:

| Library      | ns/op (full dataset) | ns/op per URL | WHATWG URL compliant ? |
|--------------|----------------------|---------------|------------------------|
| github.com/ada-url/goada   | 2,206,395            | 22.0          |   YES                  |
| Go `net/url` | 1,937,154            | 19.4          |   NO (sad face)        | 
| github.com/nlnwa/whatwg-url   | 13,916,373           | 139.0         |   Yes                  | 

Run `go test -bench BenchmarkTop100 -run -` to reproduce these results.

### Running Benchmarks

To run the URL parsing benchmarks:

1. **Clone the repository and navigate to the directory:**
   ```bash
   git clone https://github.com/ada-url/goada.git
   cd goada
   ```

2. **Install dependencies:**
   ```bash
   go mod download
   ```

3. **Run the benchmarks:**
   ```bash
   # Run all benchmarks
   go test -bench=. -run=^$
   
   # Run specific benchmark for top 100k URLs
   go test -bench=BenchmarkTop100 -run=^
   
   # Run with more detailed output
   go test -bench=BenchmarkTop100 -benchmem -run=^
   ```

4. **Compare with other libraries:**
   The benchmarks compare goada (Ada), Go's standard `net/url`, and `whatwg-url` library performance.

### Dataset

The benchmarks use the top 100k URLs dataset from [ada-url/url-various-datasets](https://github.com/ada-url/url-various-datasets), which contains the most popular URLs collected from real web traffic.

### Compliance Top 100

Testing with the top 100k URLs dataset shows that goada and whatwg-url produce identical normalized output for all URLs they both successfully parse. Go's standard `net/url` library produces different results in 1398 cases. The main differences are:

- **Query parameter encoding**: Go net/url does not encode spaces and special characters in query parameters, while Ada/whatwg-url do
- **Path normalization**: Different handling of path components

Run `go test -run TestParserComparison -v` to see some differences.

## Citation

Yagiz Nizipli, Daniel Lemire, [Parsing Millions of URLs per Second](https://doi.org/10.1002/spe.3296), Software: Practice and Experience 54(5) May 2024.
