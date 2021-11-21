package serializer_test

import (
	"pcbook/pb"
	"pcbook/sample"
	"pcbook/serializer"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
)

//https://dev.to/techschoolguru/go-generate-serialize-protobuf-message-4m7a
func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)

	require.True(t, proto.Equal(laptop1, laptop2))
}
