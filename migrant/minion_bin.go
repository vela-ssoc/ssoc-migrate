package migrant

import (
	"context"
	"log/slog"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
)

func NewMinionBin(from *query.Query, dest *query.Query, log *slog.Logger) *MinionBin {
	return &MinionBin{from: from, dest: dest, log: log}
}

type MinionBin struct {
	from *query.Query
	dest *query.Query
	log  *slog.Logger
}

func (*MinionBin) TableName() string {
	return new(model.MinionBin).TableName()
}

func (mig *MinionBin) Execute(ctx context.Context) error {
	dats, err := mig.from.MinionBin.WithContext(ctx).Find()
	if err != nil {
		return err
	}

	return mig.dest.MinionBin.WithContext(ctx).Create(dats...)
}
