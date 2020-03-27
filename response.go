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
	if r!=nil{
		return r.resp
	}
	return nil
}

func (r *Response) StatusCode() int {
	if r.resp == nil {
		return 0
	}
	return r.resp.StatusCode
}

func (r *Response) Time() string {
	if r != nil {
		return fmt.Sprintf("%dms", r.time)
	}
	return "0ms"
}

func (r *Response) Url() string {
	if r != nil {
		return r.url
	}
	return ""
}

func (r *Response) Headers() http.Header {
	if r != nil {
		return r.resp.Header
	}
	return nil
}

func (r *Response) Cookies() []*http.Cookie {
	if r != nil {
		return r.resp.Cookies()
	}
	return []*http.Cookie{}
}

func (r *Response) Body() ([]byte, error) {
	if r == nil {
		return []byte{}, errors.New("HttpRequest.Response is nil.")
	}

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

func (r *Response) Content() (string, error) {
	b, err := r.Body()
	if err != nil {
		return "", nil
	}
	return string(b), nil
}

func (r *Response) Json(v interface{}) error {
	return r.Unmarshal(v)
}

func (r *Response) Unmarshal(v interface{}) error {
	b, err := r.Body()
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	return nil
}

func (r *Response) Close() error {
	if r != nil {
		return r.resp.Body.Close()
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
