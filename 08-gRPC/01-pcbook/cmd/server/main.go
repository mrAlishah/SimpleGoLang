package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"pcbook/pb"
	"pcbook/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := flag.Int("port", 0, "the server port")
	flag.Parse()
	log.Printf("start server on port %d", *port)

	laptopStore := service.NewInMemoryLaptopStore()
	imageStore := service.NewDiskImageStore("img")
	ratingStore := service.NewInMemoryRatingStore()
	laptopServer := service.NewLaptopServer(laptopStore, imageStore, ratingStore)

	grpcServer := grpc.NewServer()
	pb.RegisterLaptopServiceServer(grpcServer, laptopServer)

	//1- google search: golang grpc reflection => Enabled Enable Server Reflection
	//https://github.com/grpc/grpc-go/blob/master/Documentation/server-reflection-tutorial.md
	reflection.Register(grpcServer)
	//2- Install Evans client
	//https://dev.to/techschoolguru/grpc-reflection-and-evans-cli-3oia  - Installation
	//3- use it CLI
	//https://dev.to/techschoolguru/grpc-reflection-and-evans-cli-3oia

	address := fmt.Sprintf("0.0.0.0:%d", *port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
