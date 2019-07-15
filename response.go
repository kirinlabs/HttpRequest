package HttpRequest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	time int64
	url  string
	resp *http.Response
	body []byte
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
	return fmt.Sprintf("%dms", r.time)
}

func (r *Response) Url() string {
	return r.url
}

func (r *Response) Headers() map[string]string {
	headers := make(map[string]string)
	for k, v := range r.resp.Header {
		if len(v) > 0 {
			headers[k] = v[len(v)-1]
		} else {
			headers[k] = ""
		}
	}
	return headers
}

func (r *Response) Body() ([]byte, error) {
	defer r.resp.Body.Close()

	if len(r.body) > 0 {
		return r.body, nil
	}

	if r.resp == nil || r.resp.Body == nil {
		return nil, errors.New("response or body is nil")
	}

	b, err := ioutil.ReadAll(r.resp.Body)
	if err != nil {
		return nil, err
	}
	r.body = b

	return b, nil
}

func (r *Response) Json(v interface{}) error {
	b, err := r.Body()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	return nil
}

func (r *Response) Export() (string, error) {
	b, err := r.Body()
	if err != nil {
		return "", err
	}

	var i interface{}
	if err := json.Unmarshal(b, &i); err != nil {
		return "", errors.New("illegal json: " + err.Error())
	}

	return Export(i), nil
}
