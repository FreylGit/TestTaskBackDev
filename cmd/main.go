package main

import (
	"context"
	"github.com/FreylGit/TestTaskBackDev/internal/app"
	"log"
)

func main() {
	ctx := context.Background()
	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatal("Failed config app")
	}
	err = a.Start(ctx)
	if err != nil {
		log.Fatal("Failed start server")
	}
}
