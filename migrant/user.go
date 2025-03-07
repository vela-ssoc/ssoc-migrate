package migrant

import (
	"context"
	"log/slog"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
)

func NewUser(from *query.Query, dest *query.Query, log *slog.Logger) *User {
	return &User{from: from, dest: dest, log: log}
}

type User struct {
	from *query.Query
	dest *query.Query
	log  *slog.Logger
}

func (mig *User) TableName() string {
	return new(model.User).TableName()
}

func (mig *User) Execute(ctx context.Context) error {
	dats, err := mig.from.User.WithContext(ctx).Find()
	if err != nil {
		return err
	}

	var inserts []*model.User
	for _, u := range dats {
		if u.Username == "root" || u.DeletedAt.Valid {
			mig.log.Info("跳过用户", slog.Any("user", u))
			continue
		}
		inserts = append(inserts, u)
	}

	return mig.dest.User.WithContext(ctx).Create(inserts...)
}
