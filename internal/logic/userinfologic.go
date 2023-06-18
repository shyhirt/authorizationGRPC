package logic

import (
	"context"
	"errors"

	"authorizationGRPC/internal/svc"
	"authorizationGRPC/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

var errUserInfo = errors.New("access denied")

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoReq) (*user.UserInfoResp, error) {
	id, ok, err := l.svcCtx.Jwt.Validate(in.AccessToken)
	if err != nil {
		return &user.UserInfoResp{}, err
	}
	if !ok {
		return &user.UserInfoResp{}, errUserInfo
	}
	currentUser, err := l.svcCtx.Model.FindOne(l.ctx, id)
	if err != nil {
		return &user.UserInfoResp{}, err
	}
	return &user.UserInfoResp{
		Id:        currentUser.Id,
		Email:     currentUser.Email,
		Username:  currentUser.Username,
		FirstName: currentUser.Firstname,
		LastName:  currentUser.Lastname,
	}, nil
}
