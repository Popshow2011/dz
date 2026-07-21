package verify

type SendRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type Response struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}
