gen:
    protoc proto/*.proto --go_out=pb
	
clean:
	del /f /s /q pb

run:
	go run main.go

pkg:
go get -u github.com/gogo/protobuf/proto
go get -u  github.com/gogo/protobuf/jsonpb
go get -u github.com/gogo/protobuf/protoc-gen-gogo
go get -u github.com/gogo/protobuf/gogoproto
