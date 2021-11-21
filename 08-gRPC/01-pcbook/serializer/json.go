package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ProtobufToJSON converts protocol buffer message to JSON string
func ProtobufToJSON(message proto.Message) (string, error) {
	// UseEnumNumbers:Write enums as integers or strings. false = enum as string
	// EmitUnpopulated:Write fields with default value or not. true= use default value
	// Indent:What's the indentation we want to use.
	// UseProtoNames:Do we want to use the original field name as in the proto file. true = set as proto field , false= set as clamecase

	//https://godocs.io/google.golang.org/protobuf/encoding/protojson#MarshalOptions
	//https://seb-nyberg.medium.com/customizing-protobuf-json-serialization-in-golang-6c58b5890356
	marshaler := protojson.MarshalOptions{
		UseEnumNumbers:  false,
		Indent:          "  ",
		UseProtoNames:   true,
		EmitUnpopulated: true,
	}

	jsonByte, err := marshaler.Marshal(message)
	return string(jsonByte), err
	//return protojson.Format(message),nil

}

// JSONToProtobufMessage converts JSON string to protocol buffer message
func JSONToProtobufMessage(data string, message proto.Message) error {
	return protojson.Unmarshal([]byte(data), message)
}
