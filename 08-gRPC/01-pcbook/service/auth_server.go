package service

import (
	"context"
	"pcbook/pb"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// AuthServer is the server for authentication
type AuthServer struct {
	userStore  UserStore
	jwtManager *JWTManager
	pb.UnimplementedAuthServiceServer
}

// NewAuthServer returns a new auth server
func NewAuthServer(userStore UserStore, jwtManager *JWTManager) *AuthServer {
	unimplemented := pb.UnimplementedAuthServiceServer{}
	return &AuthServer{userStore, jwtManager, unimplemented}
}

// Login is a unary RPC to login user
func (server *AuthServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	// find the user by username
	user, err := server.userStore.Find(req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot find user: %v", err)
	}

	if user == nil || !user.IsCorrectPassword(req.GetPassword()) {
		return nil, status.Errorf(codes.NotFound, "incorrect username/password")
	}

	token, err := server.jwtManager.Generate(user)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "cannot generate access token")
	}

	//we create a new login response object with the generated access token, and return it to the client.
	res := &pb.LoginResponse{AccessToken: token}
	return res, nil
}
