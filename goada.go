package goada

/*
#cgo CFLAGS: -O3  -std=c11
#cgo CXXFLAGS: -O3 -std=c++20
#include "ada_c.h"


// ada_parse_and_validate combines parse + validity check in a single CGo call.
// Returns the parsed URL, or NULL if invalid.
static inline ada_url ada_parse_and_validate(const char* input, size_t length) {
    ada_url url = ada_parse(input, length);
    if (!ada_is_valid(url)) {
        ada_free(url);
        return NULL;
    }
    return url;
}
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
	cleanup  runtime.Cleanup
}

// parse the given string into a URL, a finalizer
// will be set to free the URL when it is no longer needed.
func New(urlstring string) (*Url, error) {
	if len(urlstring) == 0 {
		return nil, ErrEmptyString
	}
	cptr := C.ada_parse_and_validate((*C.char)(unsafe.Pointer(unsafe.StringData(urlstring))), C.size_t(len(urlstring)))
	runtime.KeepAlive(urlstring)
	if cptr == nil {
		return nil, ErrInvalidUrl
	}
	answer := &Url{cpointer: cptr}
	answer.cleanup = runtime.AddCleanup(answer, func(cptr C.ada_url) {
		C.ada_free(cptr)
	}, cptr)
	return answer, nil
}

// parse the given strings into a URL, a finalizer
// will be set to free the URL when it is no longer needed.
func NewWithBase(urlstring string, basestring string) (*Url, error) {
	if len(urlstring) == 0 || len(basestring) == 0 {
		return nil, ErrEmptyString
	}
	cptr := C.ada_parse_with_base((*C.char)(unsafe.Pointer(unsafe.StringData(urlstring))), C.size_t(len(urlstring)), (*C.char)(unsafe.Pointer(unsafe.StringData(basestring))), C.size_t(len(basestring)))
	runtime.KeepAlive(urlstring)
	runtime.KeepAlive(basestring)
	if !C.ada_is_valid(cptr) {
		C.ada_free(cptr)
		return nil, ErrInvalidUrl
	}
	answer := &Url{cpointer: cptr}
	answer.cleanup = runtime.AddCleanup(answer, func(cptr C.ada_url) {
		C.ada_free(cptr)
	}, cptr)
	return answer, nil
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
	answer := bool(C.ada_has_password(u.cpointer))
	runtime.KeepAlive(u)
	return answer
}

func (u *Url) HasHash() bool {
	answer := bool(C.ada_has_hash(u.cpointer))
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
