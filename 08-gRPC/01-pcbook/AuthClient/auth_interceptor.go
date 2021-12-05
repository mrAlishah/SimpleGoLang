//The idea is: we will intercept all gRPC requests and attach an access token to them (if necessary) before invoking the server.
package authclient

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// AuthInterceptor is a client interceptor for authentication
type AuthInterceptor struct {
	//It contains an auth client object that will be used to login user
	authClient *AuthClient
	//map to tell us which method needs authentication
	authMethods map[string]bool
	//acquired access token
	accessToken string
}

//basically what we will do is to launch a separate go routine to periodically call login API to get a new access token before the current token expired

// NewAuthInterceptor returns a new auth interceptor
//efresh token duration parameter. It will tell us how often we should call the login API to get a new token
func NewAuthInterceptor(
	authClient *AuthClient,
	authMethods map[string]bool,
	refreshDuration time.Duration,
) (*AuthInterceptor, error) {

	//first we will create a new interceptor object
	interceptor := &AuthInterceptor{
		authClient:  authClient,
		authMethods: authMethods,
	}

	//Then  to schedule refreshing access token and pass in the refresh duration
	err := interceptor.scheduleRefreshToken(refreshDuration)
	if err != nil {
		return nil, err
	}

	return interceptor, nil
}

// Unary returns a client interceptor to authenticate unary RPC
func (interceptor *AuthInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(
		ctx context.Context,
		method string,
		req, reply interface{},
		cc *grpc.ClientConn,
		invoker grpc.UnaryInvoker,
		opts ...grpc.CallOption,
	) error {
		log.Printf("--> unary interceptor: %s", method)

		//If it does, we must attach the access token to the context before invoking the actual RPC
		if interceptor.authMethods[method] {
			return invoker(interceptor.attachToken(ctx), method, req, reply, cc, opts...)
		}

		//nside this interceptor function, let’s write a simple log with the calling method name. Then check if this method needs authentication or not.
		//If the method doesn’t require authentication, then nothing to be done, we simply invoke the RPC with the original context.
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

// Stream returns a client interceptor to authenticate stream RPC
func (interceptor *AuthInterceptor) Stream() grpc.StreamClientInterceptor {
	return func(
		ctx context.Context,
		desc *grpc.StreamDesc,
		cc *grpc.ClientConn,
		method string,
		streamer grpc.Streamer,
		opts ...grpc.CallOption,
	) (grpc.ClientStream, error) {
		log.Printf("--> stream interceptor: %s", method)

		if interceptor.authMethods[method] {
			return streamer(interceptor.attachToken(ctx), desc, cc, method, opts...)
		}

		return streamer(ctx, desc, cc, method, opts...)
	}
}

//So I will define a new attachToken() function to attach the token to the input context and return the result.
func (interceptor *AuthInterceptor) attachToken(ctx context.Context) context.Context {
	//In this function, we just use metadata.AppendToOutgoingContext(), pass in the input context together with an authorization key and the access token value.
	//Make sure that the authorization key string matches with the one we used on the server side.
	return metadata.AppendToOutgoingContext(ctx, "authorization", interceptor.accessToken)
}

func (interceptor *AuthInterceptor) scheduleRefreshToken(refreshDuration time.Duration) error {
	err := interceptor.refreshToken()
	if err != nil {
		return err
	}

	//Then after that, we launch a new go routine. Here I use a wait variable to store how much time we need to wait before refreshing the token
	go func() {
		wait := refreshDuration
		for {
			time.Sleep(wait)
			err := interceptor.refreshToken()

			//If an error occurs, we should only wait a short period of time, let’s say 1 second, before retrying it.
			//If there’s no error, then we definitely should wait for refreshDuration
			if err != nil {
				wait = time.Second
			} else {
				wait = refreshDuration
			}
		}
	}()

	return nil
}

//First we will need a function to just refresh token with no scheduling. In this function, we just use the auth client to login user
func (interceptor *AuthInterceptor) refreshToken() error {
	accessToken, err := interceptor.authClient.Login()
	if err != nil {
		return err
	}

	//we simply store it in the interceptor.accessToken field
	interceptor.accessToken = accessToken
	log.Printf("token refreshed: %v", accessToken)

	return nil
}
