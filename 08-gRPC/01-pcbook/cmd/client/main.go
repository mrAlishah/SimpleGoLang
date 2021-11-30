package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"pcbook/pb"
	"pcbook/sample"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	serverAddress := flag.String("address", "", "the server address")
	flag.Parse()
	log.Printf("dial server %s", *serverAddress)

	conn, err := grpc.Dial(*serverAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	laptopClient := pb.NewLaptopServiceClient(conn)
	testCreateLaptop(laptopClient)
	//testSearchLaptop(laptopClient)
	//testUploadImage(laptopClient)
	//testRateLaptop(laptopClient)
}

func testCreateLaptop(laptopClient pb.LaptopServiceClient) {
	createLaptop(laptopClient, sample.NewLaptop())
}

func testSearchLaptop(laptopClient pb.LaptopServiceClient) {
	for i := 0; i < 10; i++ {
		createLaptop(laptopClient, sample.NewLaptop())
	}

	filter := &pb.Filter{
		MaxPriceUsd: 3000,
		MinCpuCores: 4,
		MinCpuGhz:   2.5,
		MinRam:      &pb.Memory{Value: 8, Unit: pb.Memory_GIGABYTE},
	}

	searchLaptop(laptopClient, filter)
}

func testUploadImage(laptopClient pb.LaptopServiceClient) {
	laptop := sample.NewLaptop()
	createLaptop(laptopClient, laptop)
	uploadImage(laptopClient, laptop.GetId(), "tmp/golang.jpg")
}

func testRateLaptop(laptopClient pb.LaptopServiceClient) {

	//Let’s say we want to rate 3 laptops, so we declare a slice to keep the laptop IDs
	n := 3
	laptopIDs := make([]string, n)

	//use a for loop to generate a random laptop, save its ID to the slice, and call createLaptop() function to create it on the server
	for i := 0; i < n; i++ {
		laptop := sample.NewLaptop()
		laptopIDs[i] = laptop.GetId()
		createLaptop(laptopClient, laptop)
	}

	scores := make([]float64, n)
	for {
		fmt.Print("rate laptop (y/n)? ")
		var answer string
		fmt.Scan(&answer)

		//I will use a for loop here and ask if we want to do another round of rating or not
		if strings.ToLower(answer) != "y" {
			break
		}

		//If the answer is no, we break the loop. Else we generate a new set of scores for the laptops and call rateLaptop() function to rate them with the generated scores.
		for i := 0; i < n; i++ {
			scores[i] = sample.RandomLaptopScore()
		}

		err := rateLaptop(laptopClient, laptopIDs, scores)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func createLaptop(laptopClient pb.LaptopServiceClient, laptop *pb.Laptop) {

	//laptop.Id = ""
	req := &pb.CreateLaptopRequest{
		Laptop: laptop,
	}

	// set timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := laptopClient.CreateLaptop(ctx, req)
	//res, err := laptopClient.CreateLaptop(context.Background(), req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.AlreadyExists {
			// not a big deal
			log.Print("laptop already exists")
		} else {
			log.Fatal("cannot create laptop: ", err)
		}
		return
	}

	log.Printf("created laptop with id: %s", res.Id)
}

func searchLaptop(laptopClient pb.LaptopServiceClient, filter *pb.Filter) {
	log.Print("search filter: ", filter)

	//We first create a context with timeout of 5 seconds.
	//We make a SearchLaptopRequest object with the input filter.
	// we call laptopClient.SearchLaptop() to get the stream.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.SearchLaptopRequest{Filter: filter}
	stream, err := laptopClient.SearchLaptop(ctx, req)
	if err != nil {
		log.Fatal("cannot search laptop: ", err)
	}

	//If the stream.Recv() function call returns and end-of-file (EOF) error, this means it’s the end of the stream, so we just return.
	//Otherwise, if error is not nil, we write a fatal log.
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal("cannot receive response: ", err)
		}

		laptop := res.GetLaptop()
		log.Print("- found: ", laptop.GetId())
		log.Print("  + brand: ", laptop.GetBrand())
		log.Print("  + name: ", laptop.GetName())
		log.Print("  + cpu cores: ", laptop.GetCpu().GetNumberCores())
		log.Print("  + cpu min ghz: ", laptop.GetCpu().GetMinGhz())
		log.Print("  + ram: ", laptop.GetRam())
		log.Print("  + price: ", laptop.GetPriceUsd())
	}
}

//the laptop client, the laptop ID, and the path to the laptop image.
func uploadImage(laptopClient pb.LaptopServiceClient, laptopID string, imagePath string) {
	//First we call os.Open() to open the image file
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatal("cannot open image file: ", err)
	}
	defer file.Close()

	//Then we create a context with timeout of 5 seconds, and call laptopClient.UploadImage() with that context. It will return a stream object, and an error
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//call laptopClient.UploadImage() with that context. It will return a stream object, and an error
	stream, err := laptopClient.UploadImage(ctx)
	if err != nil {
		log.Fatal("cannot upload image: ", err)
	}

	//we create the first request to send some image information to the server
	req := &pb.UploadImageRequest{
		Data: &pb.UploadImageRequest_Info{
			Info: &pb.ImageInfo{
				LaptopId:  laptopID,
				ImageType: filepath.Ext(imagePath),
			},
		},
	}

	//we call stream.Send() to send the first request to the server. If we get an error, write a fatal log
	//he reason we got EOF is because when an error occurs, the server will close the stream, and thus the client cannot send more data to it.
	//To get the real error that contains the gRPC status code, we must call stream.RecvMsg() with a nil parameter.
	//The nil parameter basically means that we don't expect to receive any message, but we just want to get the error that function returns.
	err = stream.Send(req)
	if err != nil {
		log.Fatal("cannot send image info to server: ", err, stream.RecvMsg(nil))
	}

	//we will create a buffer reader to read the content of the image file in chunks.
	//Let’s say each chunk will be 1 KB, or 1024 bytes. We will read the image data chunks sequentially in a for loop:
	reader := bufio.NewReader(file)
	buffer := make([]byte, 1024)

	for {
		//First we call reader.Read() to read the data to the buffer. It will return the number of bytes read and an error.
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("cannot read chunk to buffer: ", err)
		}

		//Otherwise, we create a new request with the chunk data. Make sure that the chunk only contains the first n bytes of the buffer, since the last chunk might contain less than 1024 bytes.
		req := &pb.UploadImageRequest{
			Data: &pb.UploadImageRequest_ChunkData{
				ChunkData: buffer[:n],
			},
		}

		//Then we call stream.Send() to send it to the server
		err = stream.Send(req)
		if err != nil {
			log.Fatal("cannot send chunk to server: ", err, stream.RecvMsg(nil))
		}
	}

	//Finally, after the for loop, We call stream.CloseAndRecv() to receive a response from the server
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("cannot receive response: ", err)
	}

	log.Printf("image uploaded with id: %s, size: %d", res.GetId(), res.GetSize())
}

func rateLaptop(laptopClient pb.LaptopServiceClient, laptopIDs []string, scores []float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	stream, err := laptopClient.RateLaptop(ctx)
	if err != nil {
		return fmt.Errorf("cannot rate laptop: %v", err)
	}

	//we will have to make a channel to wait for the responses from the server.
	// The waitResponse channel will receive an error when it occurs, or a nil if all responses are received successfully.
	waitResponse := make(chan error)
	// go routine to receive responses
	//Note that the requests and responses are sent concurrently, so we have to start a new go routine to receive the responses.
	// In the go routine, we use a for loop, and call stream.Recv() to get a response from the server
	go func() {
		for {
			res, err := stream.Recv()
			//If error is EOF, it means there’s no more responses, so we send nil to the waitResponse channel, and return. Else,
			// if error is not nil, we send the error to the waitResponse channel, and return as well. If no errors occur, we just write a simple log.
			if err == io.EOF {
				log.Print("no more responses")
				waitResponse <- nil
				return
			}
			if err != nil {
				waitResponse <- fmt.Errorf("cannot receive stream response: %v", err)
				return
			}

			log.Print("received response: ", res)
		}
	}()

	// send requests
	for i, laptopID := range laptopIDs {
		req := &pb.RateLaptopRequest{
			LaptopId: laptopID,
			Score:    scores[i],
		}

		err := stream.Send(req)
		//Note that here we call stream.RecvMsg() to get the real error
		if err != nil {
			return fmt.Errorf("cannot send stream request: %v - %v", err, stream.RecvMsg(nil))
		}

		log.Print("sent request: ", req)
	}

	//Now one important thing that we must do after sending all requests, which is, to call stream.CloseSend() to tell the server that we won’t send any more data.
	//And finally read from the waitResponse channel and return the received error.
	err = stream.CloseSend()
	if err != nil {
		return fmt.Errorf("cannot close send: %v", err)
	}

	err = <-waitResponse
	return err
}
