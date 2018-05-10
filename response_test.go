package HttpRequest

import (
	"testing"
)

func TestResponse(t *testing.T) {
	req := NewRequest()

	resp, err := req.Get(localUrl, nil)
	if err != nil {
		t.Error(err)
		return
	}

	if resp.Response() == nil {
		t.Error("For Response()", "return nil")
	}

}

func TestStatusCode(t *testing.T) {
	req := NewRequest()

	resp, err := req.Get(localUrl, nil)
	if err != nil {
		t.Error(err)
		return
	}

	if resp.StatusCode() == 0 {
		t.Error("For StatusCode()", "return code 0")
	}

}

func TestBody(t *testing.T) {
	req := NewRequest()

	resp, err := req.Get(localUrl, nil)
	if err != nil {
		t.Error(err)
		return
	}

	b, err := resp.Body()

	if err != nil {
		t.Error(err)
		return
	}

	if b == nil {
		t.Error("For Body()", "return nil")
	}

}
