package migrant

import (
	"context"
	"log/slog"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
)

func NewBroker(from *query.Query, dest *query.Query, log *slog.Logger) *Broker {
	return &Broker{from: from, dest: dest, log: log}
}

type Broker struct {
	from *query.Query
	dest *query.Query
	log  *slog.Logger
}

func (*Broker) TableName() string {
	return new(model.Broker).TableName()
}

func (mig *Broker) Execute(ctx context.Context) error {
	dats, err := mig.from.Broker.WithContext(ctx).Find()
	if err != nil {
		return err
	}
	for _, dat := dats {
		dat.Status = false
	}

	return mig.dest.Broker.WithContext(ctx).Create(dats...)
}
