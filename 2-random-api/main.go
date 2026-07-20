package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

const port = ":8085"

func handleRandomInt(w http.ResponseWriter, r *http.Request) {
	rand := rand.Intn(6) + 1

	fmt.Fprint(w, rand)
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/random", handleRandomInt)

	server := http.Server{
		Addr:    port,
		Handler: router,
	}

	fmt.Printf("Server started %s on port", port)
	server.ListenAndServe()

}
