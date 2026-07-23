package product

type ProductCreatedRequest struct {
	Name        string   `json:"name" valdate:"required"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
}

type ProductUpdateRequest struct {
	Name        string `json:"name" valdate:"required"`
	Description string `json:"description"`
}
