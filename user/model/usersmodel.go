package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
	// UsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUsersModel.
	UsersModel interface {
		usersModel
		CheckUserByEmailOrUsername(ctx context.Context, email string, username string) (bool, error)
		GetByEmail(ctx context.Context, email string) (*Users, error)
		GetByCode(ctx context.Context, code int64) (*Users, error)
		ActivateCode(ctx context.Context, u Users) (*Users, error)
	}

	customUsersModel struct {
		*defaultUsersModel
	}
)

// NewUsersModel returns a model for the database table.
func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UsersModel {
	return &customUsersModel{
		defaultUsersModel: newUsersModel(conn, c, opts...),
	}
}

func (m *defaultUsersModel) findOneByQuery(ctx context.Context, q string, id ...any) (*Users, error) {
	usersIdKey := fmt.Sprintf("%s%v", cacheUsersIdPrefix, id)
	var resp Users
	err := m.QueryRowCtx(ctx, &resp, usersIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s %s limit 1", usersRows, m.table, q)
		return conn.QueryRowCtx(ctx, v, query, id...)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUsersModel) CheckUserByEmailOrUsername(ctx context.Context, email string, username string) (bool, error) {
	var cnt int64
	query := fmt.Sprintf("select count(*) from %s where `email` = ? OR `username` = ? ", m.table)
	err := m.QueryRowNoCacheCtx(ctx, &cnt, query, email, username)
	return cnt > 0, err
}

func (m *customUsersModel) GetByEmail(ctx context.Context, email string) (*Users, error) {
	return m.findOneByQuery(ctx, "where `email` = ?", email)
}

func (m *customUsersModel) GetByCode(ctx context.Context, code int64) (*Users, error) {
	return m.findOneByQuery(ctx, "where `code` = ?", code)
}

func (m *customUsersModel) ActivateCode(ctx context.Context, u Users) (*Users, error) {
	u.Code = sql.NullInt64{
		Int64: 0,
		Valid: false,
	}
	err := m.Update(ctx, &u)
	return &u, err
}
