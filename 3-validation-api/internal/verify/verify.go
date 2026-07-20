package verify

import (
	"fmt"
	"net/http"
)

type VerifyHandler struct{}

func NewVerifyHandler(router *http.ServeMux) {
	handler := VerifyHandler{}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("/verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("SEND")
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Verify")
	}
}
