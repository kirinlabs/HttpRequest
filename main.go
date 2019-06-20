package HttpRequest

import (
	"crypto/tls"
	"time"
)

func NewRequest() *Request {
	r := &Request{
		timeout: 30,
		headers: map[string]string{},
		cookies: map[string]string{},
	}
	return r
}

// Debug model
func Debug(v bool) *Request {
	r := NewRequest()
	return r.Debug(v)
}

func DisableKeepAlives(v bool) *Request {
	r := NewRequest()
	return r.DisableKeepAlives(v)
}

func SetTLSClient(v *tls.Config) *Request {
	r := NewRequest()
	return r.SetTLSClient(v)
}

func SetHeaders(headers map[string]string) *Request {
	r := NewRequest()
	return r.SetHeaders(headers)
}

func SetCookies(cookies map[string]string) *Request {
	r := NewRequest()
	return r.SetCookies(cookies)
}

func JSON() *Request {
	r := NewRequest()
	return r.JSON()
}

func SetTimeout(d time.Duration) *Request {
	r := NewRequest()
	return r.SetTimeout(d)
}

// Get is a get http request
func Get(url string, data ...interface{}) (*Response, error) {
	r := NewRequest()
	return r.Get(url, data...)
}

func Post(url string, data ...interface{}) (*Response, error) {
	r := NewRequest()
	return r.Post(url, data...)
}

// Put is a put http request
func Put(url string, data ...interface{}) (*Response, error) {
	r := NewRequest()
	return r.Put(url, data...)
}

// Delete is a delete http request
func Delete(url string, data ...interface{}) (*Response, error) {
	r := NewRequest()
	return r.Delete(url, data...)
}

// Upload file
func Upload(url, filename, fileinput string) (*Response, error) {
	r := NewRequest()
	return r.Upload(url, filename, fileinput)
}
