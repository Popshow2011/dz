package verify

import (
	"crypto/sha256"
	"dz/3-validation-api/pkg/file"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type VerifyHandler struct{}

func NewVerifyHandler(router *http.ServeMux) {
	handler := VerifyHandler{}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload SendRequest
		json.NewDecoder(r.Body).Decode(&payload)

		v := validator.New()
		err := v.Struct(payload)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		hasher := sha256.New()
		hasher.Write([]byte(payload.Email))
		hs := hasher.Sum(nil)

		file.WriteFile("db.json", payload.Email, string(hs))

		fmt.Println("SEND", payload, string(hs))
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
