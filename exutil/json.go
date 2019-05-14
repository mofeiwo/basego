package exutil

import "github.com/json-iterator/go"


var (
	jsonObj = jsoniter.ConfigCompatibleWithStandardLibrary
)

func EncodeMarshal(data interface{}) {
	jsonObj.Marshal(&data)
}

func DecodeUnMarshal(input []byte, data interface{})  {
	jsonObj.Unmarshal(input, &data)
}
