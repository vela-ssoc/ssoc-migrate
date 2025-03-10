package migrant

import (
	"context"
	"log/slog"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
)

func NewSubstanceTask(from *query.Query, dest *query.Query, log *slog.Logger) *SubstanceTask {
	return &SubstanceTask{from: from, dest: dest, log: log}
}

type SubstanceTask struct {
	from *query.Query
	dest *query.Query
	log  *slog.Logger
}

func (*SubstanceTask) TableName() string {
	return new(model.SubstanceTask).TableName()
}

func (mig *SubstanceTask) Execute(ctx context.Context) error {
	return mig.dest.SubstanceTask.WithContext(ctx).
		UnderlyingDB().
		Migrator().
		AutoMigrate(model.SubstanceTask{})
}
