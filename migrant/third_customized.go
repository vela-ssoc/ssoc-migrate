package migrant

import (
	"context"
	"log/slog"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
)

func NewThirdCustomized(from *query.Query, dest *query.Query, log *slog.Logger) *ThirdCustomized {
	return &ThirdCustomized{
		from: from,
		dest: dest,
		log:  log,
	}
}

type ThirdCustomized struct {
	from *query.Query
	dest *query.Query
	log  *slog.Logger
}

func (*ThirdCustomized) TableName() string {
	return new(model.ThirdCustomized).TableName()
}

func (mig *ThirdCustomized) Execute(ctx context.Context) error {
	dats, err := mig.from.ThirdCustomized.WithContext(ctx).Find()
	if err != nil {
		return err
	}

	return mig.dest.ThirdCustomized.WithContext(ctx).Create(dats...)
}
