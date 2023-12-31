// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"authorizationGRPC/internal/logic"
	"authorizationGRPC/internal/svc"
	"authorizationGRPC/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) Registration(ctx context.Context, in *user.RegReq) (*user.RegResp, error) {
	l := logic.NewRegistrationLogic(ctx, s.svcCtx)
	return l.Registration(in)
}

func (s *UserServer) Login(ctx context.Context, in *user.LoginReq) (*user.LoginResp, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *UserServer) UserInfo(ctx context.Context, in *user.UserInfoReq) (*user.UserInfoResp, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}

func (s *UserServer) CheckVerificationCode(ctx context.Context, in *user.VerificationCodeReq) (*user.VerificationCodeResp, error) {
	l := logic.NewCheckVerificationCodeLogic(ctx, s.svcCtx)
	return l.CheckVerificationCode(in)
}
