package service

import (
	"bytes"
	"context"
	"errors"
	"io"
	"log"
	"pcbook/pb"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//1 MB = 2^20 bytes = 1 << 20 bytes
const maxImageSize = 1 << 20

// LaptopServer is the server that provides laptop services
type LaptopServer struct {
	laptopStore LaptopStore
	imageStore  ImageStore
	pb.UnimplementedLaptopServiceServer
}

// NewLaptopServer returns a new LaptopServer
func NewLaptopServer(laptopStore LaptopStore, imageStore ImageStore) *LaptopServer {
	unimplemented := pb.UnimplementedLaptopServiceServer{}
	return &LaptopServer{
		laptopStore,
		imageStore,
		unimplemented,
	}
}

// CreateLaptop is a unary RPC to create a new laptop
func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	req *pb.CreateLaptopRequest,
) (*pb.CreateLaptopResponse, error) {

	//GetLaptop from request
	laptop := req.GetLaptop()
	log.Printf("receive a create-laptop request with id: %s", laptop.Id)

	//Check Laptop is real
	if len(laptop.Id) > 0 {
		// check if it's a valid UUID
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID: %v", err)
		}
	} else {
		// Create New laptop Id
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop ID: %v", err)
		}
		laptop.Id = id.String()
	}

	//Just for test cancel and deadline request some heavy processing
	//time.Sleep(6 * time.Second)
	// if ctx.Err() == context.Canceled {
	// 	log.Print("request is canceled")
	// 	return nil, status.Error(codes.Canceled, "request is canceled")
	// }
	// if ctx.Err() == context.DeadlineExceeded {
	// 	log.Print("deadline is exceeded")
	// 	return nil, status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	// }
	errCtx := contextError(ctx)
	if errCtx != nil {
		return nil, errCtx
	}

	// save the laptop to laptopStore
	err := server.laptopStore.Save(laptop)
	if err != nil {
		code := codes.Internal
		if errors.Is(err, ErrAlreadyExists) {
			code = codes.AlreadyExists
		}

		return nil, status.Errorf(code, "cannot save laptop to the laptopStore: %v", err)
	}

	log.Printf("saved laptop with id: %s", laptop.Id)

	res := &pb.CreateLaptopResponse{
		Id: laptop.Id,
	}
	return res, nil
}

// SearchLaptop is a server-streaming RPC to search for laptops
func (server *LaptopServer) SearchLaptop(
	req *pb.SearchLaptopRequest,
	stream pb.LaptopService_SearchLaptopServer,
) error {
	filter := req.GetFilter()
	log.Printf("receive a search-laptop request with filter: %v", filter)

	err := server.laptopStore.Search(
		stream.Context(),
		filter,
		func(laptop *pb.Laptop) error {
			res := &pb.SearchLaptopResponse{Laptop: laptop}
			err := stream.Send(res)
			if err != nil {
				return err
			}

			log.Printf("sent laptop with id: %s", laptop.GetId())
			return nil
		},
	)

	if err != nil {
		return status.Errorf(codes.Internal, "unexpected error: %v", err)
	}

	return nil
}

// UploadImage is a client-streaming RPC to upload a laptop image
func (server *LaptopServer) UploadImage(
	stream pb.LaptopService_UploadImageServer) error {
	//First we call stream.Recv() to receive the first request, which contains the metadata information of the image.
	// If there’s an error, we write a log and return the status code Unknown to the client.
	req, err := stream.Recv()
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot receive image info"))
	}

	//can get the laptop ID and the image type from the request
	laptopID := req.GetInfo().GetLaptopId()
	imageType := req.GetInfo().GetImageType()
	log.Printf("receive an upload-image request for laptop %s with image type %s", laptopID, imageType)

	//Before saving the laptop image, we have to make sure that the laptop ID really exists
	laptop, err := server.laptopStore.Find(laptopID)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot find laptop: %v", err))
	}
	if laptop == nil {
		//status code InvalidArgument, or you might use code NotFound
		return logError(status.Errorf(codes.InvalidArgument, "laptop id %s doesn't exist", laptopID))
	}

	imageData := bytes.Buffer{}
	imageSize := 0

	for {
		err := contextError(stream.Context())
		if err != nil {
			return err
		}

		log.Print("waiting to receive more data")

		//we can start receiving the image chunks data. So let’s create a new byte buffer to store them, and also a variable to keep track of the total image size.
		req, err := stream.Recv()

		//we first check if the error is EOF or not. If it is, this means that no more data will be sent, and we can safely break the loop.
		// Else, if the error is still not nil, we return it with Unknown status code to the client.
		if err == io.EOF {
			log.Print("no more data")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}

		//Otherwise, if there’s no error, we can get the chunk data from the request.
		// We get its size using the len() function, and add this size to the total image size.
		chunk := req.GetChunkData()
		size := len(chunk)

		log.Printf("received a chunk with size: %d", size)

		//Now if the image size is greater than max image size, we can return an error with InvalidArgument status code
		imageSize += size
		if imageSize > maxImageSize {
			return logError(status.Errorf(codes.InvalidArgument, "image is too large: %d > %d", imageSize, maxImageSize))
		}

		// write slowly
		// time.Sleep(time.Second)

		//we can append the chunk to the image data with the Write() function
		_, err = imageData.Write(chunk)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
		}
	}

	//After the for loop, we have collected all data of the image in the buffer. So we can call imageStore.Save() to save the image data to the store and get back the image ID
	imageID, err := server.imageStore.Save(laptopID, imageType, imageData)
	if err != nil {
		return logError(status.Errorf(codes.Internal, "cannot save image to the store: %v", err))
	}

	res := &pb.UploadImageResponse{
		Id:   imageID,
		Size: uint32(imageSize),
	}

	//If the image is saved successfully, we create a response object with the image ID and image size. Then we call stream.SendAndClose() to send the response to client.
	err = stream.SendAndClose(res)
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot send response: %v", err))
	}

	log.Printf("saved image with id: %s, size: %d", imageID, imageSize)
	return nil
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}
