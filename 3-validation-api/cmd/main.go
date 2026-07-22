package main

import (
	"dz/3-validation-api/internal/verify"
	"fmt"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	verify.NewVerifyHandler(router)
	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err.Error())
	}

}
