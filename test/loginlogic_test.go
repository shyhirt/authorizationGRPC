package test

import (
	pb "authorizationGRPC/user"
	"context"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"testing"
)

func TestLogin(t *testing.T) {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	assert.NoError(t, err)
	client := pb.NewUserClient(conn)
	login, pass := "test@email.com", "password"
	request := &pb.RegReq{
		Email:     login,
		Username:  faker.Name(),
		FirstName: faker.FirstName(),
		LastName:  faker.LastName(),
		Password:  pass,
	}
	client.Registration(context.Background(), request)
	requestLogin := &pb.LoginReq{
		Login:    login,
		Password: pass,
	}
	responseLogin, err := client.Login(context.Background(), requestLogin)
	assert.NoError(t, err)
	assert.NotEmpty(t, responseLogin.AccessToken)
	assert.NotEmpty(t, responseLogin.Refresh)
}
