package HttpRequest

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Response struct {
	time int64
	url string
	resp *http.Response
}

func (r *Response) Response() *http.Response {
	return r.resp
}

func (r *Response) StatusCode() int {
	if r.resp == nil {
		return 0
	}
	return r.resp.StatusCode
}

func (r *Response) Time() string {
	return strconv.Itoa(int(r.time)) + "ms"
}

func (r *Response) Url() string {
	return r.url
}

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
		return "", errors.New("Illegal json: "+err.Error())
	}

	return Json(i), nil
}
