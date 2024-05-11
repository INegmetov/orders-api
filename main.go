package main

import (
	"context"
	"fmt"
	"github.com/inegmetov/orders-api/application"
	"os"
	"os/signal"
)

func main() {
	app := application.New(application.LoadConfig())
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start App:", err)
	}
}
