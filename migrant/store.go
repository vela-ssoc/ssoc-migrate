package migrant

import (
	"context"
	"log/slog"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
)

func NewStore(from *query.Query, dest *query.Query, log *slog.Logger) *Store {
	return &Store{from: from, dest: dest, log: log}
}

type Store struct {
	from *query.Query
	dest *query.Query
	log  *slog.Logger
}

func (*Store) TableName() string {
	return new(model.Store).TableName()
}

func (mig *Store) Execute(ctx context.Context) error {
	dats, err := mig.from.Store.WithContext(ctx).Find()
	if err != nil {
		return err
	}

	return mig.dest.Store.WithContext(ctx).Create(dats...)
}
