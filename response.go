package HttpRequest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Response is a wrapper around *http.Response with more info
type Response struct {
	time int64
	url  string
	resp *http.Response
}

// Response returns the http.Response of a request
func (r *Response) Response() *http.Response {
	return r.resp
}

// StatusCode returns the status code of a request
func (r *Response) StatusCode() int {
	if r.resp == nil {
		return 0
	}
	return r.resp.StatusCode
}

// Time returns the duration of a request
func (r *Response) Time() string {
	return strconv.Itoa(int(r.time)) + "ms"
}

// Url returns the requested url
func (r *Response) Url() string {
	return r.url
}

// Body returns the body of a request
func (r *Response) Body() ([]byte, error) {
	defer r.resp.Body.Close()

	if r.resp == nil || r.resp.Body == nil {
		return nil, errors.New("response or body is nil")
	}

	b, err := ioutil.ReadAll(r.resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Json returns the json output of a request
func (r *Response) Json() (string, error) {
	defer r.resp.Body.Close()

	if r.resp == nil || r.resp.Body == nil {
		return "", errors.New("response or body is nil")
	}

	b, err := ioutil.ReadAll(r.resp.Body)
	if err != nil {
		return "", err
	}

	var i interface{}

	err = json.Unmarshal(b, &i)
	if err != nil {
		return "", errors.New("Illegal json: " + err.Error())
	}

	return Json(i), nil
}
