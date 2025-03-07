package migrant

import (
	"context"
	"log/slog"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
)

func NewCertificate(from *query.Query, dest *query.Query, log *slog.Logger) *Certificate {
	return &Certificate{from: from, dest: dest, log: log}
}

type Certificate struct {
	from *query.Query
	dest *query.Query
	log  *slog.Logger
}

func (*Certificate) TableName() string {
	return new(model.Certificate).TableName()
}

func (mig *Certificate) Execute(ctx context.Context) error {
	dats, err := mig.from.Certificate.WithContext(ctx).Find()
	if err != nil {
		return err
	}

	return mig.dest.Certificate.WithContext(ctx).Create(dats...)
}
