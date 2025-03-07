package migrant

import (
	"context"
	"log/slog"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
)

func NewElastic(from *query.Query, dest *query.Query, log *slog.Logger) *Elastic {
	return &Elastic{from: from, dest: dest, log: log}
}

type Elastic struct {
	from *query.Query
	dest *query.Query
	log  *slog.Logger
}

func (*Elastic) TableName() string {
	return new(model.Elastic).TableName()
}

func (mig *Elastic) Execute(ctx context.Context) error {
	dats, err := mig.from.Elastic.WithContext(ctx).Find()
	if err != nil {
		return err
	}

	return mig.dest.Elastic.WithContext(ctx).Create(dats...)
}
