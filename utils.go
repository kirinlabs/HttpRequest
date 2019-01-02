package HttpRequest

import (
	"bytes"
	"encoding/json"
)

// Export converts v into a nice json stringified output
func Export(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	var buf bytes.Buffer
	err = json.Indent(&buf, b, "", "\t")
	if err != nil {
		return ""
	}
	return buf.String()
}

// Json converts v into a nice json stringified output
func Json(v interface{}) string {
	return Export(v)
}
