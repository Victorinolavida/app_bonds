package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"boundsApp.victorinolavida/internal/data"
	_ "github.com/lib/pq"
)

type rateLimit struct {
	enable bool
	limit  int
}
type config struct {
	port      int
	dsn       string
	secret    string
	rateLimit rateLimit
}

type application struct {
	config config
	logger *slog.Logger
	models data.Models
}

func main() {

	var config config

	flag.IntVar(&config.port, "port", 4000, "API server port")
	flag.StringVar(&config.dsn, "dsn", "postgres://postgres:postgres@localhost/bonds?sslmode=disable", "Postgres connection string")
	flag.StringVar(&config.secret, "secret", "super_duper_secret", "JTW secret")
	flag.BoolVar(&config.rateLimit.enable, "limiter-enable", true, "Enable rate limit")
	flag.IntVar(&config.rateLimit.limit, "limiter-request-min", 1000, "Rate limit per minute")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := ConnectDB(config)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()
	logger.Info("database connection pool established")
	models := data.NewModels(db)

	app := &application{
		config: config,
		models: models,
		logger: logger,
	}

	server := &http.Server{
		Addr:     fmt.Sprintf(":%d", app.config.port),
		Handler:  app.routes(),
		ErrorLog: slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
	}

	app.logger.Info("Starting server", "addr", app.config.port)

	err = server.ListenAndServe()

	if err != nil {
		panic(1)
	}
}

func ConnectDB(conf config) (*sql.DB, error) {
	db, err := sql.Open("postgres", conf.dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
