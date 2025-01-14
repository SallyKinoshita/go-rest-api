package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/SallyKinoshita/go-rest-api/backend/internal/config"
	"github.com/SallyKinoshita/go-rest-api/backend/internal/di"
	"github.com/SallyKinoshita/go-rest-api/backend/internal/interfaces/router"
	"github.com/SallyKinoshita/go-rest-api/backend/pkg/clog"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(ctx); err != nil {
		// clog.Initがコケたらclogが使えないので、fmt.Printlnでエラーを出力
		fmt.Println("server run error: ", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	// 環境変数
	cfg, err := config.New(ctx)
	if err != nil {
		return err
	}

	// カスタムログ
	if err := clog.New(cfg.App.Name, clog.LogTypeApp); err != nil {
		return err
	}

	// DB接続
	db, err := config.NewDB(&cfg.DB)
	if err != nil {
		return err
	}

	defer func() {
		if err := db.Close(); err != nil {
			clog.Error(ctx, err)
		}
	}()

	// DI
	di, err := di.New(ctx, cfg, db)
	if err != nil {
		return err
	}

	// Router
	r := router.New(db, cfg, di.Controller, di.Middleware).SetUp(ctx)

	// サーバー起動
	r.Logger.Fatal(r.Start(fmt.Sprintf(":%s", cfg.App.Port)))

	return nil
}
