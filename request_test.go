package HttpRequest

import (
	"fmt"
	"testing"
)

const localUrl = "http://127.0.0.1:8000"

var data = map[string]interface{}{
	"name":    "HttpRequest",
	"version": "v1.0",
}

func TestGetRequest(t *testing.T) {
	req := NewRequest()

	test := []map[string]interface{}{
		nil,
		data,
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

	test := []map[string]interface{}{
		data,
		nil,
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
