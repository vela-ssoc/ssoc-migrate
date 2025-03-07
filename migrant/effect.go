package migrant

import (
	"context"
	"log/slog"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
)

func NewEffect(from *query.Query, dest *query.Query, log *slog.Logger) *Effect {
	return &Effect{from: from, dest: dest, log: log}
}

type Effect struct {
	from *query.Query
	dest *query.Query
	log  *slog.Logger
}

func (*Effect) TableName() string {
	return new(model.Effect).TableName()
}

func (mig *Effect) Execute(ctx context.Context) error {
	dats, err := mig.from.Effect.WithContext(ctx).Find()
	if err != nil {
		return err
	}

	return mig.dest.Effect.WithContext(ctx).Create(dats...)
}
