package logic

import (
	"authorizationGRPC/internal/svc"
	"authorizationGRPC/user"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"time"
)

var errLogin = errors.New("password is incorrect")

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func validatePassword(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	u, err := l.svcCtx.Model.GetByEmail(l.ctx, in.Login)
	if err != nil {
		return &user.LoginResp{}, err
	}
	if !validatePassword(in.Password, u.Password) {
		return &user.LoginResp{}, errLogin
	}
	token, err := l.svcCtx.Jwt.GetJwtToken(time.Now().Unix(), l.svcCtx.Config.Jwt.AccessExpire, u.Id)
	if err != nil {
		return &user.LoginResp{}, err
	}
	refresh := l.svcCtx.Refresh.Set(u.Id)
	return &user.LoginResp{
		AccessToken: token,
		Refresh:     refresh.Refresh,
	}, nil
}
