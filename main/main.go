package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/vela-ssoc/ssoc-migrate/launch"
)

func main() {
	cares := []os.Signal{syscall.SIGTERM, syscall.SIGHUP, syscall.SIGKILL, syscall.SIGINT}
	ctx, cancel := signal.NotifyContext(context.Background(), cares...)
	defer cancel()

	logOpts := &slog.HandlerOptions{AddSource: true, Level: slog.LevelDebug}
	logHandler := slog.NewJSONHandler(os.Stdout, logOpts)
	log := slog.New(logHandler)

	if err := launch.Run(ctx, "application.jsonc"); err != nil {
		log.Error("运行出错", slog.Any("error", err))
	} else {
		log.Info("执行结束")
	}
}
