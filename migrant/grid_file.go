import (
	"context"
	"log/slog"

	"github.com/vela-ssoc/vela-common-mb/dal/model"
	"github.com/vela-ssoc/vela-common-mb/dal/query"
)

func NewGridFile(from *query.Query, dest *query.Query, log *slog.Logger) *GridFile {
	return &GridFile{from: from, dest: dest, log: log}
}

type GridFile struct {
	from *query.Query
	dest *query.Query
	log  *slog.Logger
}

func (*GridFile) TableName() string {
	return new(model.GridFile).TableName()
}

func (mig *GridFile) Execute(ctx context.Context) error {
	dats, err := mig.from.GridFile.WithContext(ctx).Find()
	if err != nil {
		return err
	}

	return mig.dest.GridFile.WithContext(ctx).Create(dats...)
}
