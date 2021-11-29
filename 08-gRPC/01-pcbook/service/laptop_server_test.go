package service_test

import (
	"context"
	"pcbook/pb"
	"pcbook/sample"
	"pcbook/service"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestServerCreateLaptop(t *testing.T) {
	t.Parallel()

	laptopNoID := sample.NewLaptop()
	laptopNoID.Id = ""

	laptopInvalidID := sample.NewLaptop()
	laptopInvalidID.Id = "invalid-uuid"

	laptopDuplicateID := sample.NewLaptop()
	storeDuplicateID := service.NewInMemoryLaptopStore()
	err := storeDuplicateID.Save(laptopDuplicateID)
	require.Nil(t, err)

	testCases := []struct {
		name        string
		laptop      *pb.Laptop
		laptopStore service.LaptopStore
		code        codes.Code
	}{
		{
			name:        "success_with_id",
			laptop:      sample.NewLaptop(),
			laptopStore: service.NewInMemoryLaptopStore(),
			code:        codes.OK,
		},
		{
			name:        "success_no_id",
			laptop:      laptopNoID,
			laptopStore: service.NewInMemoryLaptopStore(),
			code:        codes.OK,
		},
		{
			name:        "failure_invalid_id",
			laptop:      laptopInvalidID,
			laptopStore: service.NewInMemoryLaptopStore(),
			code:        codes.InvalidArgument,
		},
		{
			name:        "failure_duplicate_id",
			laptop:      laptopDuplicateID,
			laptopStore: storeDuplicateID,
			code:        codes.AlreadyExists,
		},
	}

	for i := range testCases {
		//We must save the current test case to a local variable.
		//This is very important to avoid concurrency issues, because we want to create multiple parallel subtests.
		tc := testCases[i]

		//To create a subtest, we call t.Run() and use tc.name for the name of the subtest.
		//We call t.Parallel() to make it run in parallel with other tests.
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req := &pb.CreateLaptopRequest{
				Laptop: tc.laptop,
			}
			server := service.NewLaptopServer(tc.laptopStore, nil, nil)
			res, err := server.CreateLaptop(context.Background(), req)

			//when tc.code is OK
			if tc.code == codes.OK {
				//no error.
				require.NoError(t, err)
				//response should be not nil
				require.NotNil(t, res)
				//The returned ID should be not empty
				require.NotEmpty(t, res.Id)
				//if the input laptop already has ID, then the returned ID should equal to it.
				if len(tc.laptop.Id) > 0 {
					require.Equal(t, tc.laptop.Id, res.Id)
				}
			} else {
				//there should be an error
				require.Error(t, err)
				//the response should be nil
				require.Nil(t, res)
				//To check the status code, we call status.FromError() to get the status object.
				//Check that ok should be true and st.Code() should equal to tc.code
				st, ok := status.FromError(err)
				require.True(t, ok)
				require.Equal(t, tc.code, st.Code())
			}
		})
	}
}
