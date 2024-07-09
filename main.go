package main

import (
	"api/app"
	"context"
	"fmt"
)

func main() {
	app := app.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}
