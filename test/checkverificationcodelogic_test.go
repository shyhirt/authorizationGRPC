package test

import (
	pb "authorizationGRPC/user"
	"context"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
)

func TestVerification(t *testing.T) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	assert.NoError(t, err)
	client := pb.NewUserClient(conn)
	request := &pb.VerificationCodeReq{
		Code: 282688,
	}
	response, err := client.CheckVerificationCode(context.Background(), request)
	assert.NoError(t, err)
	assert.Equal(t, true, response.Result)
}
