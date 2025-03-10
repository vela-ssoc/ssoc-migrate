package migrant

import (
	"context"
	"log/slog"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
)

func NewThird(from *query.Query, dest *query.Query, log *slog.Logger) *Third {
	return &Third{
		from: from,
		dest: dest,
		log:  log,
	}
}

type Third struct {
	from *query.Query
	dest *query.Query
	log  *slog.Logger
}

func (*Third) TableName() string {
	return new(model.Third).TableName()
}

func (mig *Third) Execute(ctx context.Context) error {
	dats, err := mig.from.Third.WithContext(ctx).Find()
	if err != nil {
		return err
	}

	return mig.dest.Third.WithContext(ctx).Create(dats...)
}
