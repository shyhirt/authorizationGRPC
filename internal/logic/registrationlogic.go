package logic

import (
	"authorizationGRPC/internal/svc"
	"authorizationGRPC/user"
	"authorizationGRPC/user/model"
	"authorizationGRPC/utils/emailSender"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

type RegistrationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegistrationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegistrationLogic {
	return &RegistrationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

const randMax = 999999
const randMin = 100000
const cost = 14

var regex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

var ErrRegistration = errors.New("user already exist")
var ErrRegistrationValidateEmail = errors.New("invalid email")

func genPassword(password string) (hash string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
func validateEmail(email string) bool {
	indexAt, indexPoint := strings.Index(email, "@"), strings.Index(email, ".")
	if strings.Count(email, ".") <= 1 && (indexAt > indexPoint || indexAt < 0) {
		return false
	}
	return regex.MatchString(email)
}

func (l *RegistrationLogic) Registration(in *user.RegReq) (*user.RegResp, error) {
	if !validateEmail(in.Email) {
		return &user.RegResp{}, ErrRegistrationValidateEmail
	}
	exist, err := l.svcCtx.Model.CheckUserByEmailOrUsername(l.ctx, in.Email, in.Username)
	if err != nil {
		return &user.RegResp{}, err
	}
	if exist {
		return &user.RegResp{}, ErrRegistration
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := r.Int63n(randMax-randMin) + randMin
	in.Password, err = genPassword(in.Password)
	if err != nil {
		return &user.RegResp{}, err
	}
	res, err := l.svcCtx.Model.Insert(l.ctx, &model.Users{
		Code: sql.NullInt64{
			Int64: code,
			Valid: true,
		},
		Username:  in.Username,
		Email:     in.Email,
		Password:  in.Password,
		Firstname: in.FirstName,
		Lastname:  in.LastName,
	})
	if err != nil {
		return &user.RegResp{}, err
	}
	l.svcCtx.ThreadEmailSender.SetMessage(
		emailSender.Email{
			Message: fmt.Sprintf("%d", code),
			To:      []string{in.Email},
		})
	id, err := res.LastInsertId()
	return &user.RegResp{
		Id: id,
	}, err
}
