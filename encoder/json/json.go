package json

import (
	"github.com/json-iterator/go"
	"github.com/micro/go-config/encoder"
)

type jsonEncoder struct{}

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

// UseNumber fix unmarshal Number(8234567890123456789) to interface(8.234567890123457e+18)
func UseNumber() {
	json = jsoniter.Config{
		UseNumber:              true,
		EscapeHTML:             true,
		SortMapKeys:            true,
		ValidateJsonRawMessage: true,
	}.Froze()
}

func (j jsonEncoder) Encode(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (j jsonEncoder) Decode(d []byte, v interface{}) error {
	return json.Unmarshal(d, v)
}

func (j jsonEncoder) String() string {
	return "json"
}

func NewEncoder() encoder.Encoder {
	return jsonEncoder{}
}
