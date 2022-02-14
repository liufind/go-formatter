package formatter

import (
	"bytes"
	"encoding/json"
)

func setJSON(in interface{}) (out string, err error) {
	var data []byte

	if data, err = json.Marshal(in); err != nil {
		return "", err
	}

	return string(data), nil
}

func setIndent(in string) (string, error) {
	var buffer bytes.Buffer

	if err := json.Indent(&buffer, []byte(in), "", "\t"); err != nil {
		return "", err
	}

	return buffer.String(), nil
}
