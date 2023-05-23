package goada

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
)

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

	if strings.Compare(url.Href(), "https://www.google.com/") != 0 {
		t.Error("Expected normalized url")
	}
	if strings.Compare(url.Protocol(), "https:") != 0 {
		t.Error("Expected https protocol")
	}
}

func TestGoodUrlSet(t *testing.T) {
	url, err := New("https://www.GOogle.com")
	if err != nil {
		t.Error("Expected no error")
	}
	fmt.Println(url.Href())

	if strings.Compare(url.Href(), "https://www.google.com/") != 0 {
		t.Error("Expected normalized url")
	}
	if strings.Compare(url.Protocol(), "https:") != 0 {
		t.Error("Expected https protocol")
	}
	url.SetProtocol("http:")
	if strings.Compare(url.Protocol(), "http:") != 0 {
		t.Error("Expected http protocol")
	}
	url.SetHash("goada")
	fmt.Println(url.Hash())

	if strings.Compare(url.Hash(), "#goada") != 0 {
		t.Error("Expected goada hash")
	}
	fmt.Println(url.Href())
	if strings.Compare(url.Href(), "http://www.google.com/#goada") != 0 {
		t.Error("Expected normalized url")
	}
}

// go test -bench Benchmark -run -
func BenchmarkSillyAda(b *testing.B) {
	for j := 0; j < b.N; j++ {
		_, err := New("https://www.Googlé.com")
		if err != nil {
			break
		}
	}
}

func TestStandard(t *testing.T) {
	s1 := "https://	www.GOoglé.com/./path/../path2/"
	url, err := New(s1)
	if err != nil {
		t.Error("Expected no error")
	}
	fmt.Println(url.Href())
	if strings.Compare(url.Href(), "https://www.xn--googl-fsa.com/path2/") != 0 {
		t.Error("Expected normalized url")
	}
	url.Free()
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
	if strings.Compare(url.String(), "https://www.GOogl%C3%A9.com/./path/../path2/") != 0 {
		t.Error("Expected invalid normalized url")
	}
}
