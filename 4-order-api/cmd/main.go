package main

import (
	"dz/4-order-api/configs"
	"fmt"
)

func main() {
	config := configs.NewConfig()
	fmt.Println(config)
}
