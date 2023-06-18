package test

import (
	"authorizationGRPC/internal/logic"
	pb "authorizationGRPC/user"
	"context"
	"errors"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
)

var rpcErr = errors.New("rpc error: code = Unknown desc = ")

func TestRegistration(t *testing.T) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	assert.NoError(t, err)
	client := pb.NewUserClient(conn)
	request := &pb.RegReq{
		Email:     faker.Email(),
		Username:  faker.Username(),
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Password:  faker.Password(),
	}
	response, err := client.Registration(context.Background(), request)
	assert.NoError(t, err)
	assert.Greater(t, response.Id, int64(0))
}

func TestRegistrationWrongEmail(t *testing.T) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	assert.NoError(t, err)
	client := pb.NewUserClient(conn)
	request := &pb.RegReq{
		Email:     faker.Name(),
		Username:  faker.Username(),
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Password:  faker.Password(),
	}
	_, err = client.Registration(context.Background(), request)
	assert.Equal(t, err.Error(), rpcErr.Error()+""+logic.ErrRegistrationValidateEmail.Error())
}

func TestRegistrationDuplicate(t *testing.T) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	assert.NoError(t, err)
	client := pb.NewUserClient(conn)
	request := &pb.RegReq{
		Email:     "uVCpMjU@loZGxIy.net",
		Username:  "oQhlUrk",
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Password:  faker.Password(),
	}
	_, err = client.Registration(context.Background(), request)
	assert.Equal(t, err.Error(), rpcErr.Error()+""+logic.ErrRegistration.Error())
}
