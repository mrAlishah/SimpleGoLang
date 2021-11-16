gen:
	protoc proto/*.proto --go_out=pb


clean:
	rm pb/*.go

run:
	go run main.go

pro:
	 protoc --proto_path=proto/*.proto --go_out=pb --go_opt=paths=source_relative 

other:
    protoc -I. proto/*.proto --go_opt=paths=source_relative --go_out=plugins=grpc:pb
    protoc --go_out=plugins=grpc:pb proto/*.proto

	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    proto/*.proto

	protoc proto/processor_message.proto --go_out=. --go_opt=paths=source_relative
	protoc proto/processor_message.proto --go_out=pb --go_opt=paths=source_relative
	protoc proto/processor_message.proto --go-grpc_out=. --go-grpc_opt=paths=source_relative