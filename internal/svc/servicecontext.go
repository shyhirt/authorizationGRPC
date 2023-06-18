package svc

import (
	"authorizationGRPC/internal/config"
	"authorizationGRPC/user/model"
	"authorizationGRPC/utils/jwtUtil"
	"authorizationGRPC/utils/refresh"
	"authorizationGRPC/utils/thread"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config            config.Config
	Model             model.UsersModel
	Jwt               *jwtUtil.JwtUtil
	Refresh           *refresh.Refresh
	ThreadEmailSender *thread.EmailSenderThread
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		Model:             model.NewUsersModel(sqlx.NewMysql(c.DataSource), c.Cache),
		Jwt:               jwtUtil.NewJwtUtil(c.Jwt.AccessSecret, c.Jwt.AccessPublic),
		Refresh:           refresh.NewRefresh(c.Jwt.RefreshExpire),
		ThreadEmailSender: thread.NewThreadEmailSender(c),
	}
}
