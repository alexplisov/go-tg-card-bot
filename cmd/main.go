package main

import (
	"fmt"

	"github.com/alexplisov/go-tg-card-bot/internal"
)

func main() {
	if err := internal.Run(); err != nil {
		fmt.Printf("error running app: %v\n", err)
	}
}
