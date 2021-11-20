package serializer

import (
	"strings"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ProtobufToJSON converts protocol buffer message to JSON string
func ProtobufToJSON(message proto.Message) (string, error) {
	// Write enums as integers or strings.
	// Write fields with default value or not.
	// What's the indentation we want to use.
	// Do we want to use the original field name as in the proto file.

	//https://godocs.io/google.golang.org/protobuf/encoding/protojson#MarshalOptions
	//https://seb-nyberg.medium.com/customizing-protobuf-json-serialization-in-golang-6c58b5890356
	marshaler := protojson.MarshalOptions{
		UseEnumNumbers: false,
		Indent:         "  ",
		UseProtoNames:  true,
	}

	jsonByte, err := marshaler.Marshal(message)
	return string(jsonByte), err
	//return protojson.Format(message),nil

	//------------------------------------------------------
	//github.com/golang/protobuf/jsonpb
	// marshaler := jsonpb.Marshaler{
	// 	EnumsAsInts:  false,
	// 	EmitDefaults: true,
	// 	Indent:       "  ",
	// 	OrigName:     true,
	// }

	// return marshaler.MarshalToString(message)

	//------------------------------------------------------
	//"google.golang.org/protobuf/encoding/protojson"
	// marshaler, err := protojson.MarshalOptions{
	// 	EnumsAsInts:  false,
	// 	EmitDefaults: true,
	// 	Indent:       "  ",
	// 	OrigName:     true,
	// }

	// return string(marshaler), err
}

// JSONToProtobufMessage converts JSON string to protocol buffer message
func JSONToProtobufMessage(data string, message proto.Message) error {
	return new(Unmarshaler).Unmarshal(strings.NewReader(data), message)
	//return jsonpb.UnmarshalString(data, message)
}
