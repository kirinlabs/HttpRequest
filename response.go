package HttpRequest

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Response struct {
	time int64
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

func (r *Response) Body() ([]byte, error) {
	defer r.resp.Body.Close()

	if r.resp == nil {
		return nil, errors.New("response is nil")
	}

	if r.resp.Body == nil {
		return nil, errors.New("response body is nil")
	}

	b, err := ioutil.ReadAll(r.resp.Body)

	if err != nil {
		return nil, err
	}

	return b, nil
}
