package HttpRequest

import (
	"bytes"
	"fmt"
	"testing"
)

const localUrl = "http://localhost:8000"

var data = map[string]interface{}{
	"name":    "HttpRequest",
	"version": "v1.0",
}

func TestGetRequest(t *testing.T) {
	req := NewRequest()

	test := []interface{}{
		nil,
		data,
		"year=2018",
	}

	var resp *Response
	var err error

	for _, v := range test {
		resp, err = req.Get(localUrl, v)
		if err != nil {
			t.Error(err)
			return
		}
	}

	if resp.StatusCode() != 200 {
		t.Error("GET "+localUrl, "expected code 200", fmt.Sprintf("return code %d", resp.StatusCode()))
	}
}

func TestPostRequest(t *testing.T) {
	req := NewRequest()

	test := []interface{}{
		data,
		bytes.NewReader([]byte{97}),
		[]byte{97},
		nil,
		"github",
		`{"name":"github","year":2018}`,
		100,
		int8(100),
		int16(100),
		int32(100),
		int64(100),
		"title=github&type=1",
	}

	var resp *Response
	var err error

	for _, v := range test {
		resp, err = req.Post(localUrl, v)
		if err != nil {
			t.Error(err)
			return
		}
	}

	if resp.StatusCode() != 200 {
		t.Error("GET "+localUrl, "expected code 200", fmt.Sprintf("return code %d", resp.StatusCode()))
	}

}
