package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/osamikoyo/qrk/internal/app"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	app.Init().Run(ctx)
}