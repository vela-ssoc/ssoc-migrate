package migrant

import (
	"context"
	"log/slog"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
)

func NewEmc(from *query.Query, dest *query.Query, log *slog.Logger) *Emc {
	return &Emc{from: from, dest: dest, log: log}
}

type Emc struct {
	from *query.Query
	dest *query.Query
	log  *slog.Logger
}

func (*Emc) TableName() string {
	return new(model.Emc).TableName()
}

func (mig *Emc) Execute(ctx context.Context) error {
	dats, err := mig.from.Emc.WithContext(ctx).Find()
	if err != nil {
		return err
	}

	return mig.dest.Emc.WithContext(ctx).Create(dats...)
}
