package app

import (
	"context"
	"fmt"
	"os"

	"github.com/diyor200/uof/internal/gateway/rest"
	"github.com/diyor200/uof/internal/repository"
	"github.com/diyor200/uof/internal/usecase/users"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func Run() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}

	defer conn.Close(context.Background())
	err = conn.Ping(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to database!")

	// repos
	repo := repository.NewRepos(conn)

	// use cases
	usecases := users.New(repo)

	// configure handlers
	handler := rest.NewHandler(usecases)
	router := handler.NewRouter(gin.Default())

	err = router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
