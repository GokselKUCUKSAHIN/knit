package custom_json

import (
	"github.com/json-iterator/go"
)

var JsonIter = jsoniter.ConfigDefault

func Unmarshal(data []byte, obj interface{}) error {
	return JsonIter.Unmarshal(data, obj)
}

func Marshal(obj interface{}) ([]byte, error) {
	return JsonIter.Marshal(obj)
}
