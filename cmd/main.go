package main

import (
	"fmt"

	"github.com/Denio1337/go-wallet-service/internal/config"
)

func main() {
	fmt.Println(config.Get(config.EnvAppAddress))
}
