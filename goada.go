package goada

/*
#cgo CFLAGS: -O3  -std=c11
#cgo CXXFLAGS: -O3 -std=c++17
#include "ada_c.h"

const char empty_string[] = "";
*/
import "C"
import (
	"errors"
	"runtime"
	"unsafe"
)

func free(u *Url) {
	C.ada_free(u.cpointer)
}

var ErrEmptyString = errors.New("empty url string")
var ErrInvalidUrl = errors.New("invalid url")

type Url struct {
	cpointer C.ada_url
}

// parse the given string into a URL, a finalizer
// will be set to free the URL when it is no longer needed.
func New(urlstring string) (*Url, error) {
	if len(urlstring) == 0 {
		return nil, ErrEmptyString
	}
	var answer *Url
	answer = &Url{C.ada_parse((*C.char)(unsafe.Pointer(unsafe.StringData(urlstring))), C.size_t(len(urlstring)))}
	runtime.KeepAlive(urlstring)
	if !C.ada_is_valid(answer.cpointer) {
		C.ada_free(answer.cpointer)
		return nil, ErrInvalidUrl
	}
	runtime.SetFinalizer(answer, free)
	return answer, nil
}

// parse the given strings into a URL, a finalizer
// will be set to free the URL when it is no longer needed.
func NewWithBase(urlstring string, basestring string) (*Url, error) {
	if len(urlstring) == 0 || len(basestring) == 0 {
		return nil, ErrEmptyString
	}
	var answer *Url
	answer = &Url{C.ada_parse_with_base((*C.char)(unsafe.Pointer(unsafe.StringData(urlstring))), C.size_t(len(urlstring)), (*C.char)(unsafe.Pointer(unsafe.StringData(basestring))), C.size_t(len(basestring)))}
	runtime.KeepAlive(urlstring)
	runtime.KeepAlive(basestring)
	if !C.ada_is_valid(answer.cpointer) {
		C.ada_free(answer.cpointer)
		return nil, ErrInvalidUrl
	}
	runtime.SetFinalizer(answer, free)
	return answer, nil
}

func (rb *Url) Free() {
	// Clear the finalizer to avoid double frees
	runtime.SetFinalizer(rb, nil)
	free(rb)
}

func (u *Url) Valid() bool {
	answer := bool(C.ada_is_valid(u.cpointer))
	runtime.KeepAlive(u)
	return answer
}

func (u *Url) HasCredentials() bool {
	answer := bool(C.ada_has_credentials(u.cpointer))
	runtime.KeepAlive(u)
	return answer
}
func (u *Url) HasEmptyHostname() bool {
	answer := bool(C.ada_has_empty_hostname(u.cpointer))
	runtime.KeepAlive(u)
	return answer
}

func (u *Url) HasHostname() bool {
	answer := bool(C.ada_has_hostname(u.cpointer))
	runtime.KeepAlive(u)
	return answer
}

func (u *Url) HasNonEmptyUsername() bool {
	answer := bool(C.ada_has_non_empty_username(u.cpointer))
	runtime.KeepAlive(u)
	return answer
}

func (u *Url) HasNonEmptyPassword() bool {
	answer := bool(C.ada_has_non_empty_password(u.cpointer))
	runtime.KeepAlive(u)
	return answer
}

func (u *Url) HasPort() bool {
	answer := bool(C.ada_has_port(u.cpointer))
	runtime.KeepAlive(u)
	return answer
}

func (u *Url) HasPassword() bool {
	answer := bool(C.ada_has_port(u.cpointer))
	runtime.KeepAlive(u)
	return answer
}

func (u *Url) HasHash() bool {
	answer := bool(C.ada_has_port(u.cpointer))
	runtime.KeepAlive(u)
	return answer
}

func (u *Url) HasSearch() bool {
	answer := bool(C.ada_has_search(u.cpointer))
	runtime.KeepAlive(u)
	return answer
}

func (u *Url) Href() string {
	ada_string := C.ada_get_href(u.cpointer)
	answer := C.GoStringN(ada_string.data, C.int(ada_string.length))
	runtime.KeepAlive(u)
	return answer
}

func (u *Url) Username() string {
	ada_string := C.ada_get_username(u.cpointer)
	answer := C.GoStringN(ada_string.data, C.int(ada_string.length))
	runtime.KeepAlive(u)
	return answer
}

func (u *Url) Password() string {
	ada_string := C.ada_get_password(u.cpointer)
	answer := C.GoStringN(ada_string.data, C.int(ada_string.length))
	runtime.KeepAlive(u)
	return answer
}

func (u *Url) Port() string {
	ada_string := C.ada_get_port(u.cpointer)
	answer := C.GoStringN(ada_string.data, C.int(ada_string.length))
	runtime.KeepAlive(u)
	return answer
}

func (u *Url) Hash() string {
	ada_string := C.ada_get_hash(u.cpointer)
	answer := C.GoStringN(ada_string.data, C.int(ada_string.length))
	runtime.KeepAlive(u)
	return answer
}

func (u *Url) Host() string {
	ada_string := C.ada_get_host(u.cpointer)
	answer := C.GoStringN(ada_string.data, C.int(ada_string.length))
	runtime.KeepAlive(u)
	return answer
}

func (u *Url) Hostname() string {
	ada_string := C.ada_get_hostname(u.cpointer)
	answer := C.GoStringN(ada_string.data, C.int(ada_string.length))
	runtime.KeepAlive(u)
	return answer
}

func (u *Url) Pathname() string {
	ada_string := C.ada_get_pathname(u.cpointer)
	answer := C.GoStringN(ada_string.data, C.int(ada_string.length))
	runtime.KeepAlive(u)
	return answer
}

func (u *Url) Search() string {
	ada_string := C.ada_get_search(u.cpointer)
	answer := C.GoStringN(ada_string.data, C.int(ada_string.length))
	runtime.KeepAlive(u)
	return answer
}

func (u *Url) Protocol() string {
	ada_string := C.ada_get_protocol(u.cpointer)
	answer := C.GoStringN(ada_string.data, C.int(ada_string.length))
	runtime.KeepAlive(u)
	return answer
}

func (u *Url) SetHref(s string) bool {
	answer := C.ada_set_href(u.cpointer, (*C.char)(unsafe.Pointer(unsafe.StringData(s))), C.size_t(len(s)))
	runtime.KeepAlive(u)
	runtime.KeepAlive(s)
	return bool(answer)
}

func (u *Url) SetHost(s string) bool {
	answer := C.ada_set_host(u.cpointer, (*C.char)(unsafe.Pointer(unsafe.StringData(s))), C.size_t(len(s)))
	runtime.KeepAlive(u)
	runtime.KeepAlive(s)
	return bool(answer)
}

func (u *Url) SetHostname(s string) bool {
	answer := C.ada_set_hostname(u.cpointer, (*C.char)(unsafe.Pointer(unsafe.StringData(s))), C.size_t(len(s)))
	runtime.KeepAlive(u)
	runtime.KeepAlive(s)
	return bool(answer)
}

func (u *Url) SetProtocol(s string) bool {
	answer := C.ada_set_protocol(u.cpointer, (*C.char)(unsafe.Pointer(unsafe.StringData(s))), C.size_t(len(s)))
	runtime.KeepAlive(u)
	runtime.KeepAlive(s)
	return bool(answer)
}

func (u *Url) SetUsername(s string) bool {
	answer := C.ada_set_username(u.cpointer, (*C.char)(unsafe.Pointer(unsafe.StringData(s))), C.size_t(len(s)))
	runtime.KeepAlive(u)
	runtime.KeepAlive(s)
	return bool(answer)
}

func (u *Url) SetPassword(s string) bool {
	answer := C.ada_set_password(u.cpointer, (*C.char)(unsafe.Pointer(unsafe.StringData(s))), C.size_t(len(s)))
	runtime.KeepAlive(u)
	runtime.KeepAlive(s)
	return bool(answer)
}

func (u *Url) SetPort(s string) bool {
	answer := C.ada_set_port(u.cpointer, (*C.char)(unsafe.Pointer(unsafe.StringData(s))), C.size_t(len(s)))
	runtime.KeepAlive(u)
	runtime.KeepAlive(s)
	return bool(answer)
}

func (u *Url) SetPathname(s string) bool {
	answer := C.ada_set_pathname(u.cpointer, (*C.char)(unsafe.Pointer(unsafe.StringData(s))), C.size_t(len(s)))
	runtime.KeepAlive(u)
	runtime.KeepAlive(s)
	return bool(answer)
}

func (u *Url) SetSearch(s string) {
	C.ada_set_search(u.cpointer, (*C.char)(unsafe.Pointer(unsafe.StringData(s))), C.size_t(len(s)))
	runtime.KeepAlive(u)
	runtime.KeepAlive(s)
}

func (u *Url) SetHash(s string) {
	C.ada_set_hash(u.cpointer, (*C.char)(unsafe.Pointer(unsafe.StringData(s))), C.size_t(len(s)))
	runtime.KeepAlive(u)
	runtime.KeepAlive(s)
}
