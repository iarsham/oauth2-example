package main

import (
	"fmt"
	"github.com/iarsham/oauth2-example/configs"
	"github.com/iarsham/oauth2-example/internal/database"
	"github.com/iarsham/oauth2-example/internal/routers"
	"github.com/iarsham/oauth2-example/pkg/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"net/http"
	"time"
)

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		panic(err)
	}

	logs, err := logger.NewZapLog(cfg.App.Debug)
	if err != nil {
		panic(err)
	}
	defer logs.Sync()

	db, err := database.OpenDB()
	if err != nil {
		logs.Fatal(err.Error())
	}
	defer db.Close()
	logs.Info("Connected to database successfully")

	srv := http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.App.Port),
		Handler:      routers.New(db, logs, cfg),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	logs.Info("Starting server", zap.String("host", cfg.App.Host), zap.Int("port", cfg.App.Port))
	if err := srv.ListenAndServe(); err != nil {
		logs.Fatal(err.Error())
	}
}

func init() {
	if err := godotenv.Load("./.env"); err != nil {
		panic(err)
	}
}
