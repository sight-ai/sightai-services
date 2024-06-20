package sight_middleware

import (
	"bytes"
	"encoding/json"
	"strings"
)

func compactJson(jsonBytes []byte) []byte {
	buffer := &bytes.Buffer{}
	_ = json.Indent(buffer, jsonBytes, "", "")
	str := buffer.String()
	newStr := strings.Replace(str, "\n", "", -1)
	if newStr == "" {
		newStr = "\"\""
	}

	return []byte(newStr)
}
