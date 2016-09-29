package yaml

import (
	"bytes"
	"encoding/json"

	goyaml "gopkg.in/yaml.v2"
)

type JSON struct {
	EscapeHTML  bool
	JSONDecoder bool
}

func (j *JSON) Marshal(v interface{}) ([]byte, error) {
	keyJSON := &bytes.Buffer{}
	encoder := json.NewEncoder(keyJSON)
	encoder.SetEscapeHTML(j.EscapeHTML)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}
	return keyJSON.Bytes()[:keyJSON.Len()-1], nil
}

func (j *JSON) Unmarshal(data []byte, v interface{}) error {
	if j.JSONDecoder {
		return json.Unmarshal(data, v)
	}
	return goyaml.Unmarshal(data, v)
}
