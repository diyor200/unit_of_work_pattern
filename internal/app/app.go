package app

import (
	"context"
	"log"
	"os"

	"github.com/diyor200/uof/internal/gateway/rest/v1"
	"github.com/diyor200/uof/internal/uow"
	"github.com/diyor200/uof/internal/usecase/users"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/fx"
)

func NewFXApp() *fx.App {
	return fx.New(
		fx.Provide(
			NewPGXPool,
			uow.NewUOWManager,
			users.New,
			v1.NewServer,
		),
		fx.Invoke(StartServer),
	)
}

// NewPGXPool provides pgxpool
func NewPGXPool(lc fx.Lifecycle) (*pgxpool.Pool, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	pareConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		log.Println("Error parsing config:", err)
		return nil, err
	}
	pool, err := pgxpool.ConnectConfig(context.Background(), pareConfig)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			log.Println("Starting PGX pool")
			log.Println("pinging database ...")
			err := pool.Ping(context.Background())
			if err != nil {
				log.Println("Error pinging database:", err)
				return err
			}

			log.Println("Connected to database!")
			return nil
		},
		OnStop: func(context.Context) error {
			log.Println("Stopping PGX pool")
			pool.Close()
			return nil
		},
	})

	return pool, nil
}

func StartServer(lc fx.Lifecycle, server *v1.Server) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				log.Println("Starting server")
				err := server.Router.Run(":8080")
				if err != nil {
					log.Println("Error starting server:", err)
					return
				}
			}()

			log.Println("Server started on :8080")
			return nil
		},
		OnStop: func(context.Context) error {
			log.Println("Stopping server")
			return nil
		},
	})
}
