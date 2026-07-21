package verify

import (
	"crypto/sha256"
	"dz/3-validation-api/configs"
	"dz/3-validation-api/pkg/file"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/jordan-wright/email"
)

type VerifyHandler struct{}

func NewVerifyHandler(router *http.ServeMux) {
	handler := VerifyHandler{}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		config := configs.LoadCofig()
		var payload SendRequest
		json.NewDecoder(r.Body).Decode(&payload)

		v := validator.New()
		err := v.Struct(payload)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(402)
			return
		}

		hasher := sha256.New()
		hasher.Write([]byte(payload.Email))
		hs := hasher.Sum(nil)

		file.WriteFile("db.json", payload.Email, hex.EncodeToString(hs))

		e := &email.Email{
			To:      []string{config.Address},
			From:    config.Address,
			Subject: "Подтверждение регистрации",
			HTML:    fmt.Appendf(nil, "<a href='http://localhost:8081/verify/%s'>Подтвердить email %s</a>", hex.EncodeToString(hs), payload.Email),
		}

		go func(emailAddr string) {
			auth := smtp.PlainAuth("", config.Address, config.Password, "smtp.yandex.ru")

			err := e.Send("smtp.yandex.ru:587", auth)
			if err != nil {
				log.Printf("ОШИБКА отправки email на %s: %v", emailAddr, err)
			} else {
				log.Printf("Письмо успешно отправлено на %s", emailAddr)
			}
		}(payload.Email)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := map[string]string{
			"message": "Verification email has been sent",
		}
		json.NewEncoder(w).Encode(response)

		fmt.Println("SEND", payload, string(hs))
	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var items []file.File
		b, err := os.ReadFile("db.json")
		if err != nil {
			fmt.Println(err)
		}
		err = json.Unmarshal(b, &items)
		if err != nil {
			fmt.Println(err)
		}

		hash := r.PathValue("hash")
		var newJson []file.File
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		for _, item := range items {
			if item.Hash != hash {
				newJson = append(newJson, item)

				response := false
				json.NewEncoder(w).Encode(response)
			}

			response := false
			json.NewEncoder(w).Encode(response)
			fmt.Println(item.Hash)
		}

		data, err := json.MarshalIndent(newJson, "", "  ")
		if err != nil {
			fmt.Println("marshal error: %w", err)
		}

		err = os.WriteFile("db.json", data, 0644)
		if err != nil {
			fmt.Println("write file error: %w", err)
		}
		fmt.Println("Verify")
	}
}
