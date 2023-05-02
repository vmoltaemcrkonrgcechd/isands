package main

import (
	"isands/config"
	"isands/internal/app"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	if err := app.Run(cfg); err != nil {
		panic(err)
	}
}
