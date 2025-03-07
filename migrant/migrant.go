package migrant

import "context"

type Migranter interface {
	TableName() string
	Execute(context.Context) error
}
