package main

import (
	"context"
	"fmt"

	"github.com/la-haus/master-brokers-charge-leads/cmd/chargeLeadsJob/di"
)

func main() {
	app := di.NewApp()
	if err := app.Start(context.Background()); err != nil {
		fmt.Println(err)
	}
}
