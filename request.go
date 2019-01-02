package HttpRequest

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"reflect"
	"strings"
	"time"
)

// Request is the base type for this package
type Request struct {
	cli               *http.Client
	req               *http.Request
	debug             bool
	url               string
	method            string
	time              int64
	timeout           time.Duration
	headers           map[string]string
	cookies           map[string]string
	data              map[string]interface{}
	disableKeepAlives bool
	tlsClientConfig   *tls.Config
}

// NewRequest creates an instance of the Request
func NewRequest() *Request {
	r := &Request{timeout: 30}
	return r
}

// DisableKeepAlives sets the disableKeepAlives property on Request
func (r *Request) DisableKeepAlives(v bool) *Request {
	r.disableKeepAlives = v
	return r
}

// SetTLSClient sets the tlsClientConfig property on Request
func (r *Request) SetTLSClient(v *tls.Config) *Request {
	r.tlsClientConfig = v
	return r
}

// Debug sets the debug property on Request
func (r *Request) Debug(v bool) *Request {
	r.debug = v
	return r
}

// buildClient builds the client
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
				TLSClientConfig:       r.tlsClientConfig,
				DisableKeepAlives:     r.disableKeepAlives,
			},
		}
	}
	return r.cli
}

// SetHeaders Sets the headers
func (r *Request) SetHeaders(h map[string]string) *Request {
	r.headers = h
	return r
}

// initHeaders Inits the headers
func (r *Request) initHeaders() {
	r.req.Header.Set("Content-Type", "x-www-form-urlencoded")
	for k, v := range r.headers {
		r.req.Header.Set(k, v)
	}
}

// SetCookies can set a cookies
func (r *Request) SetCookies(c map[string]string) *Request {
	r.cookies = c
	return r
}

// initCookies Inits the cookies
func (r *Request) initCookies() {
	for k, v := range r.cookies {
		r.req.AddCookie(&http.Cookie{
			Name:  k,
			Value: v,
		})
	}
}

// isJson Checks if the application/json headers exists
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

// buildBody Builds the query data
func (r *Request) buildBody(d map[string]interface{}) (io.Reader, error) {
	// GET and DELETE request dose not send body
	if r.method == "GET" || r.method == "DELETE" {
		return nil, nil
	}

	if d == nil || len(d) == 0 {
		return strings.NewReader(""), nil
	}

	if r.isJson() {
		if b, err := json.Marshal(d); err != nil {
			return nil, err
		} else {
			return bytes.NewReader(b), nil
		}
	}

	data := make([]string, 0)
	for k, v := range d {
		if s, ok := v.(string); ok {
			data = append(data, fmt.Sprintf("%s=%v", k, s))
			continue
		}
		b, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		data = append(data, fmt.Sprintf("%s=%s", k, string(b)))
	}

	return strings.NewReader(strings.Join(data, "&")), nil
}

// SetTimeout sets the timeout property of the Request
func (r *Request) SetTimeout(d time.Duration) *Request {
	r.timeout = d
	return r
}

// parseQuery parses a query for a GET request
func parseQuery(url string) ([]string, error) {
	urlList := strings.Split(url, "?")
	if len(urlList) < 2 {
		return make([]string, 0), nil
	}
	query := make([]string, 0)
	for _, val := range strings.Split(urlList[1], "&") {
		v := strings.Split(val, "=")
		if len(v) < 2 {
			return make([]string, 0), errors.New("query parameter error")
		}
		query = append(query, fmt.Sprintf("%s=%s", v[0], v[1]))
	}
	return query, nil
}

// buildUrl builds a GET request url
func buildUrl(url string, data map[string]interface{}) (string, error) {
	query, err := parseQuery(url)
	if err != nil {
		return url, err
	}

	if data != nil {
		for k, v := range data {
			vv := ""
			if reflect.TypeOf(v).String() == "string" {
				vv = v.(string)
			} else {
				b, err := json.Marshal(v)
				if err != nil {
					return url, err
				}
				vv = string(b)
			}
			query = append(query, fmt.Sprintf("%s=%s", k, vv))
		}
	}
	list := strings.Split(url, "?")

	if len(query) > 0 {
		return fmt.Sprintf("%s?%s", list[0], strings.Join(query, "&")), nil
	}

	return list[0], nil
}

// elapsedTime calculates the elapsed time
func (r *Request) elapsedTime(n int64, resp *Response) {
	end := time.Now().UnixNano() / 1e6
	resp.time = end - n
}

// log logs a request when in debug mode
func (r *Request) log() {
	if r.debug {
		fmt.Printf("[HttpRequest]\n")
		fmt.Printf("-------------------------------------------------------------------\n")
		fmt.Printf("Request: %s %s\nHeaders: %v\nCookies: %v\nTimeout: %ds\nBodyMap: %v\n", r.method, r.url, r.headers, r.cookies, r.timeout, r.data)
		fmt.Printf("-------------------------------------------------------------------\n\n")
	}
}

// Get is a get http request
func (r *Request) Get(url string, data map[string]interface{}) (*Response, error) {
	return r.request(http.MethodGet, url, data)
}

// Post is a post http request
func (r *Request) Post(url string, data map[string]interface{}) (*Response, error) {
	return r.request(http.MethodPost, url, data)
}

// Put is a put http request
func (r *Request) Put(url string, data map[string]interface{}) (*Response, error) {
	return r.request(http.MethodPut, url, data)
}

// Delete is a delete http request
func (r *Request) Delete(url string, data map[string]interface{}) (*Response, error) {
	return r.request(http.MethodDelete, url, data)
}

// request sends a http request
func (r *Request) request(method, url string, data map[string]interface{}) (*Response, error) {
	// Build Response
	response := &Response{}

	// Start time
	start := time.Now().UnixNano() / 1e6
	// Count elapsed time
	defer r.elapsedTime(start, response)

	if method == "" || url == "" {
		return nil, errors.New("parameter method and url is required")
	}

	// Debug infomation
	defer r.log()

	r.url = url
	r.data = data

	var (
		err  error
		body io.Reader
	)
	r.cli = r.buildClient()

	method = strings.ToUpper(method)
	r.method = method

	if method == "GET" || method == "DELETE" {
		r.url, err = buildUrl(url, data)
		if err != nil {
			return nil, err
		}
	}

	body, err = r.buildBody(data)

	if err != nil {
		return nil, err
	}

	r.req, err = http.NewRequest(method, r.url, body)

	if err != nil {
		return nil, err
	}

	r.initHeaders()
	r.initCookies()

	resp, err := r.cli.Do(r.req)

	if err != nil {
		return nil, err
	}

	response.url = r.url
	response.resp = resp

	return response, nil
}
