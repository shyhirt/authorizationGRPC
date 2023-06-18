package logic

import (
	"context"
	"errors"

	"authorizationGRPC/internal/svc"
	"authorizationGRPC/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckVerificationCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckVerificationCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckVerificationCodeLogic {
	return &CheckVerificationCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

var errCheckVerificationCode = errors.New("incorrect code")

func (l *CheckVerificationCodeLogic) CheckVerificationCode(in *user.VerificationCodeReq) (*user.VerificationCodeResp, error) {
	u, err := l.svcCtx.Model.GetByCode(l.ctx, in.Code)
	if err != nil {
		return &user.VerificationCodeResp{}, err
	}
	if u.Id == 0 {
		return &user.VerificationCodeResp{}, errCheckVerificationCode
	}
	return &user.VerificationCodeResp{
		Result: true,
	}, err
}
