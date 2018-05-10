package HttpRequest

import (
	"io"
	"net/http"
	"strings"

	"bytes"

	"errors"

	"fmt"

	"time"

	"net"

	"github.com/gin-gonic/gin/json"
)

type Request struct {
	cli     *http.Client
	req     *http.Request
	method  string
	timeout time.Duration
	headers map[string]string
	cookies map[string]string
	data    map[string]interface{}
}

// Create an instance of the Request
func NewRequest() *Request {
	r := &Request{timeout: 30}
	return r
}

// Set headers
func (r *Request) SetHeaders(h map[string]string) {
	r.headers = h
	return
}

// Init headers
func (r *Request) initHeaders() {
	r.req.Header.Set("Content-Type", "x-www-form-urlencoded")
	for k, v := range r.headers {
		r.req.Header.Set(k, v)
	}
}

// Set cookies
func (r *Request) SetCookies(c map[string]string) {
	r.cookies = c
	return
}

// Init cookies
func (r *Request) initCookies() {
	for k, v := range r.cookies {
		r.req.AddCookie(&http.Cookie{
			Name:  k,
			Value: v,
		})
	}
}

// Check application/json
func (r *Request) isJson() bool {
	if len(r.headers) > 0 {
		for _, v := range r.headers {
			if strings.Contains(strings.ToLower(v), "application/json") {
				return true
			}
		}
	}
	return false
}

// Build query data
func (r *Request) buildBody(d map[string]interface{}) (io.Reader, error) {
	// Get request dose not send body
	if r.method == "GET" {
		return nil, nil
	}

	if d != nil && len(d) > 0 {
		if r.isJson() {
			if b, err := json.Marshal(d); err != nil {
				return nil, err
			} else {
				return bytes.NewReader(b), nil
			}
		}
		data := make([]string, 0)
		for k, v := range d {
			b, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			data = append(data, fmt.Sprintf("%s=%s", k, string(b)))
		}
		return strings.NewReader(strings.Join(data, "&")), nil
	}

	return nil, errors.New("data is empty.")
}

func (r *Request) SetTimeout(d time.Duration) {
	r.timeout = d
	return
}

// Build client
func (r *Request) buildClient() *http.Client {
	if r.cli == nil {
		r.cli = &http.Client{
			Transport: &http.Transport{
				Dial: func(network, addr string) (net.Conn, error) {
					conn, err := net.DialTimeout(network, addr, time.Second*r.timeout)
					if err != nil {
						return nil, err
					}
					conn.SetDeadline(time.Now().Add(time.Second * r.timeout))
					return conn, nil
				},
				ResponseHeaderTimeout: time.Second * r.timeout,
				//DisableKeepAlives:     true,
			},
		}
	}
	return r.cli
}

// Get is a get http request
func (r *Request) Get(url string) (*Response, error) {
	return r.send(http.MethodGet, url, nil)
}

// Post is a post http request
func (r *Request) Post(url string, data map[string]interface{}) (*Response, error) {
	return r.send(http.MethodPost, url, data)
}

// Put is a put http request
func (r *Request) Put(url string, data map[string]interface{}) (*Response, error) {
	return r.send(http.MethodPut, url, data)
}

// Delete is a delete http request
func (r *Request) Delete(url string) (*Response, error) {
	return r.send(http.MethodDelete, url, nil)
}

// Send http request
func (r *Request) send(method, url string, data map[string]interface{}) (*Response, error) {
	response := Response{}
	r.cli = r.buildClient()
	var (
		err  error
		body io.Reader
	)
	method = strings.ToUpper(method)
	r.method = method

	body, err = r.buildBody(data)

	if err != nil {
		return nil, err
	}

	r.req, err = http.NewRequest(method, url, body)

	if err != nil {
		return nil, err
	}

	r.initHeaders()
	r.initCookies()

	resp, err := r.cli.Do(r.req)

	if err != nil {
		return nil, err
	}

	response.resp = resp

	return &response, nil
}
