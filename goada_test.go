package goada

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"testing"

	whatwgurl "github.com/nlnwa/whatwg-url/url"
)

var benchmarkURLs []string

func init() {
	if _, err := os.Stat("top100.txt"); os.IsNotExist(err) {
		// Download the file
		resp, err := http.Get("https://raw.githubusercontent.com/ada-url/url-various-datasets/refs/heads/main/top100/top100.txt")
		if err != nil {
			fmt.Println("Failed to download benchmark dataset, benchmarks will use empty dataset")
			return
		}
		defer resp.Body.Close()
		file, err := os.Create("top100.txt")
		if err != nil {
			fmt.Println("Failed to create benchmark dataset file, benchmarks will use empty dataset")
			return
		}
		defer file.Close()
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			fmt.Println("Failed to write benchmark dataset, benchmarks will use empty dataset")
			return
		}
	}
	// Read the file
	file, err := os.Open("top100.txt")
	if err != nil {
		fmt.Println("Failed to open benchmark dataset file, benchmarks will use empty dataset")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		benchmarkURLs = append(benchmarkURLs, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read benchmark dataset, benchmarks will use empty dataset")
		return
	}
}

func compareString(t *testing.T, expected, actual, message string) {
	if expected != actual {
		t.Errorf("Expected %s, but got %s. %s", expected, actual, message)
	}
}

func TestBadUrl(t *testing.T) {
	url, err := New("some bad url")
	if err == nil {
		t.Error("Expected error")
	}
	if url != nil {
		t.Error("Expected no URL")
	}
}

func TestGoodUrl(t *testing.T) {
	url, err := New("https://www.GOogle.com")
	if err != nil {
		t.Error("Expected no error")
	}
	fmt.Println(url.Href())

	compareString(t, "https://www.google.com/", url.Href(), "Expected normalized url")
	compareString(t, "https:", url.Protocol(), "Expected https protocol")
}

func TestGoodUrlSet(t *testing.T) {
	url, err := New("https://www.GOogle.com")
	if err != nil {
		t.Error("Expected no error")
	}
	fmt.Println(url.Href())

	compareString(t, "https://www.google.com/", url.Href(), "Expected normalized url")
	compareString(t, "https:", url.Protocol(), "Expected https protocol")
	url.SetProtocol("http:")
	compareString(t, "http:", url.Protocol(), "Expected http protocol")
	url.SetHash("goada")
	fmt.Println(url.Hash())

	compareString(t, "#goada", url.Hash(), "Expected goada hash")
	fmt.Println(url.Href())
	compareString(t, "http://www.google.com/#goada", url.Href(), "Expected normalized url")
}

// Parse results are stored so the compiler cannot dead-code the calls.
var (
	benchSinkAda    *Url
	benchSinkStd    *url.URL
	benchSinkWhatWG *whatwgurl.Url
)

func requireBenchmarkURLs(b *testing.B) []string {
	b.Helper()
	if len(benchmarkURLs) == 0 {
		b.Fatal("empty benchmark dataset")
	}
	return benchmarkURLs
}

// go test -bench BenchmarkTop100 -run -
func BenchmarkTop100Ada(b *testing.B) {
	urls := requireBenchmarkURLs(b)
	var out *Url
	for j := 0; j < b.N; j++ {
		// Parse every URL, including ones that fail. Breaking on error
		// would let each library stop at a different input.
		for _, u := range urls {
			out, _ = New(u)
		}
	}
	benchSinkAda = out
}

func BenchmarkTop100Standard(b *testing.B) {
	urls := requireBenchmarkURLs(b)
	var out *url.URL
	for j := 0; j < b.N; j++ {
		for _, u := range urls {
			out, _ = url.Parse(u)
		}
	}
	benchSinkStd = out
}

func BenchmarkTop100WhatWG(b *testing.B) {
	urls := requireBenchmarkURLs(b)
	var out *whatwgurl.Url
	for j := 0; j < b.N; j++ {
		for _, u := range urls {
			out, _ = whatwgurl.Parse(u)
		}
	}
	benchSinkWhatWG = out
}

func TestStandard(t *testing.T) {
	s1 := "https://	www.GOoglé.com/./path/../path2/"
	url, err := New(s1)
	if err != nil {
		t.Error("Expected no error")
	}
	fmt.Println(url.Href())
	compareString(t, "https://www.xn--googl-fsa.com/path2/", url.Href(), "Expected normalized url")
}

func TestStandardGP(t *testing.T) {
	s1 := "https://	www.GOoglé.com"
	_, err := url.Parse(s1)
	if err == nil {
		t.Error("Go url should fail")
	}
	s2 := "https://www.GOoglé.com/./path/../path2/"
	url, err2 := url.Parse(s2)
	if err2 != nil {
		t.Error("Go url should hot fail")
	}
	fmt.Println(url.String())
	compareString(t, "https://www.GOogl%C3%A9.com/./path/../path2/", url.String(), "Expected invalid normalized url")
}

func TestParserComparison(t *testing.T) {
	disagreeCount := 0
	totalURLs := len(benchmarkURLs)
	var disagreements []string

	for i, u := range benchmarkURLs {
		if i%10000 == 0 {
			fmt.Printf("Processed %d/%d URLs\n", i, totalURLs)
		}

		// Parse with Ada
		adaURL, adaErr := New(u)
		if adaErr != nil {
			// Ada failed, skip this URL
			continue
		}
		adaHref := adaURL.Href()

		// Parse with whatwg-url
		whatwgParsed, whatwgErr := whatwgurl.Parse(u)
		if whatwgErr != nil {
			// whatwg-url failed, but Ada succeeded - this is a difference
			disagreeCount++
			if len(disagreements) < 10 {
				disagreements = append(disagreements, fmt.Sprintf("whatwg-url failed on: %s", u))
			}
			continue
		}
		whatwgHref := whatwgParsed.Href(false)

		// Check Ada agrees with whatwg-url
		if adaHref != whatwgHref {
			t.Errorf("Ada and whatwg-url disagree on URL %q: Ada=%q, whatwg=%q", u, adaHref, whatwgHref)
		}

		// Parse with Go net/url
		goURL, goErr := url.Parse(u)
		if goErr != nil {
			// Go failed, but Ada succeeded - this is a difference
			disagreeCount++
			if len(disagreements) < 10 {
				disagreements = append(disagreements, fmt.Sprintf("net/url failed on: %s", u))
			}
			continue
		}
		goHref := goURL.String()

		// Check if Go disagrees with Ada
		if goHref != adaHref {
			disagreeCount++
			// Skip cases where the difference is just one character (e.g. trailing slash)
			if len(adaHref) == len(goHref)+1 || len(goHref) == len(adaHref)+1 || goHref == adaHref+"%20" {
				continue
			}
			disagreements = append(disagreements, fmt.Sprintf("URL: '%s'\n  net/url: '%s'\n  Ada:     '%s'", u, goHref, adaHref))
		}
	}

	fmt.Printf("Go net/url disagreed with Ada/whatwg-url on %d out of %d URLs\n", disagreeCount, totalURLs)
	if len(disagreements) > 0 {
		fmt.Println("Sample disagreements:")
		count := 0
		for _, d := range disagreements {
			fmt.Println(" -", d)
			count++
			if count >= 10 {
				break
			}
		}
	}
}
