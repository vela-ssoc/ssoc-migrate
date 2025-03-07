package migrant

import (
	"context"
	"log/slog"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
)

func NewNotifier(from *query.Query, dest *query.Query, log *slog.Logger) *Notifier {
	return &Notifier{from: from, dest: dest, log: log}
}

type Notifier struct {
	from *query.Query
	dest *query.Query
	log  *slog.Logger
}

func (mig *Notifier) TableName() string {
	return new(model.Notifier).TableName()
}

func (mig *Notifier) Execute(ctx context.Context) error {
	ntfs, err := mig.from.Notifier.WithContext(ctx).Find()
	if err != nil {
		return err
	}

	return mig.dest.Notifier.WithContext(ctx).Create(ntfs...)
}
