package main

import (
	"context"
	"fmt"

	"github.com/diyor200/uof/internal"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		internal.Module(),
		fx.Invoke(startServer),
	)

	app.Run()
}

func startServer(lc fx.Lifecycle, router *gin.Engine) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := router.Run(":8080"); err != nil {
					fmt.Printf("Failed to start server: %v\n", err.Error())
				}
			}()
			fmt.Println("server started at :8080")
			return nil
		},
	})
}
