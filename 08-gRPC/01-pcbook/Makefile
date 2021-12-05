gen:
    protoc proto/*.proto --go_out=pb --go-grpc_out=pb 
	
#  --go-grpc_opt=require_unimplemented_servers=false

clean:
	del /f /s /q pb

server:
    go run cmd/server/main.go -port 8080

client:
    go run cmd/client/main.go -address 0.0.0.0:8080

run:
	go run main.go

pro:
	 protoc --proto_path=proto/*.proto --go_out=pb --go_opt=paths=source_relative 

test:
	 go test -cover -race ./...

gen1:
    protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb

clean1:
    rm pb/*.go 

.PHONY: gen clean server client test   

fmt:
   go fmt ./...
# -------------------------------------------------------------------
other:
    protoc -I. proto/*.proto --go_opt=paths=source_relative --go_out=plugins=grpc:pb
    protoc --go_out=plugins=grpc:pb proto/*.proto

    protoc -I. proto/*.proto --go_out=plugins=grpc:pb

	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/*.proto

	protoc proto/processor_message.proto --go_out=. --go_opt=paths=source_relative
	protoc proto/processor_message.proto --go_out=pb --go_opt=paths=source_relative
	protoc proto/processor_message.proto --go-grpc_out=. --go-grpc_opt=paths=source_relative