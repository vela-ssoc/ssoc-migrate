package migrant

import (
	"context"
	"log/slog"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
)

func NewSubstance(from *query.Query, dest *query.Query, log *slog.Logger) *Substance {
	return &Substance{from: from, dest: dest, log: log}
}

type Substance struct {
	from *query.Query
	dest *query.Query
	log  *slog.Logger
}

func (*Substance) TableName() string {
	return new(model.Substance).TableName()
}

func (mig *Substance) Execute(ctx context.Context) error {
	dats, err := mig.from.Substance.WithContext(ctx).Find()
	if err != nil {
		return err
	}

	return mig.dest.Substance.WithContext(ctx).Create(dats...)
}
