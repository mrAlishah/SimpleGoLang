package authclient

import (
	"context"
	"pcbook/pb"
	"time"

	"google.golang.org/grpc"
)

// AuthClient is a client to call authentication RPC, to call authentication service
type AuthClient struct {
	service  pb.AuthServiceClient
	username string
	password string
}

// NewAuthClient returns a new auth client
func NewAuthClient(cc *grpc.ClientConn, username string, password string) *AuthClient {
	//he service is created by calling pb.NewAuthServiceClient() and pass in the connection
	service := pb.NewAuthServiceClient(cc)
	return &AuthClient{service, username, password}
}

// Login login user and returns the access token
func (client *AuthClient) Login() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &pb.LoginRequest{
		Username: client.username,
		Password: client.password,
	}

	res, err := client.service.Login(ctx, req)
	if err != nil {
		return "", err
	}

	return res.GetAccessToken(), nil
}
