package migrant

import (
	"context"
	"log/slog"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
)

func NewGridChunk(from *query.Query, dest *query.Query, log *slog.Logger) *GridChunk {
	return &GridChunk{from: from, dest: dest, log: log}
}

type GridChunk struct {
	from *query.Query
	dest *query.Query
	log  *slog.Logger
}

func (*GridChunk) TableName() string {
	return new(model.GridChunk).TableName()
}

func (mig *GridChunk) Execute(ctx context.Context) error {
	dats, err := mig.from.GridChunk.WithContext(ctx).Find()
	if err != nil {
		return err
	}

	return mig.dest.GridChunk.WithContext(ctx).Create(dats...)
}
