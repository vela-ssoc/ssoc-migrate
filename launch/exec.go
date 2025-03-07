package launch

import (
	"context"
	"log/slog"
	"os"

	"github.com/vela-ssoc/ssoc-migrate/migrant"

	"github.com/vela-ssoc/vela-common-mb/dal/query"
	"github.com/vela-ssoc/vela-common-mb/jsonc"
	"github.com/vela-ssoc/vela-common-mb/sqldb"
)

type config struct {
	From datasourceConfig `json:"from"`
	Dest datasourceConfig `json:"dest"`
}

type datasourceConfig struct {
	DSN string `json:"dsn"`
}

func Run(ctx context.Context, cfg string) error {
	data, err := os.ReadFile(cfg)
	if err != nil {
		return err
	}
	conf := new(config)
	if err = jsonc.Unmarshal(data, conf); err != nil {
		return err
	}

	return Exec(ctx, conf.From.DSN, conf.Dest.DSN)
}

func Exec(ctx context.Context, fromDSN, destDSN string) error {
	logOpts := &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug}
	logHandler := slog.NewJSONHandler(os.Stdout, logOpts)
	log := slog.New(logHandler)
	fromDB, err := sqldb.Open(fromDSN, log)
	if err != nil {
		return err
	}
	destDB, err := sqldb.Open(destDSN, log)
	if err != nil {
		return err
	}

	from := query.Use(fromDB)
	dest := query.Use(destDB)

	migrants := []migrant.Migranter{
		migrant.NewUser(from, dest, log),
		migrant.NewStore(from, dest, log),
		migrant.NewEmc(from, dest, log),
		migrant.NewNotifier(from, dest, log),
		migrant.NewSubstance(from, dest, log),
		migrant.NewEffect(from, dest, log),
		migrant.NewCertificate(from, dest, log),
		migrant.NewBroker(from, dest, log),
	}
	log.Info("开始执行数据迁移")
	for _, mig := range migrants {
		name := mig.TableName()
		attrs := []any{slog.String("name", name)}
		log.Info("正在迁移数据", attrs...)
		if exx := mig.Execute(ctx); exx != nil {
			attrs = append(attrs, slog.Any("error", exx))
			log.Error("迁移出错", attrs...)
		} else {
			log.Info("迁移成功", attrs...)
		}
	}
	log.Info("数据迁移执行结束")

	return nil
}
