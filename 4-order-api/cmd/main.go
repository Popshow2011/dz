package main

import (
	"dz/4-order-api/configs"
	"fmt"
)

func main() {
	config := configs.LoadConfig()
	fmt.Println(config)
}
