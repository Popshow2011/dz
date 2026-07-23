package product

import (
	"dz/4-order-api/pkg/req"
	"dz/4-order-api/pkg/res"
	"fmt"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type ProductHandlerDeps struct {
	ProductRepository *ProductRepository
}

type ProductHandler struct {
	ProductRepository *ProductRepository
}

func NewProductHandler(router *http.ServeMux, deps ProductHandlerDeps) {
	handler := ProductHandler{
		ProductRepository: deps.ProductRepository,
	}

	router.HandleFunc("POST /product", handler.Create())
	router.HandleFunc("GET /product/{id}", handler.Find())
	router.HandleFunc("PATCH /product/{id}", handler.Update())
	router.HandleFunc("DELETE /product/{id}", handler.Delete())
}

func (handler *ProductHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[ProductCreatedRequest](&w, r)
		if err != nil {
			return
		}
		p := NewProduct(body.Name, body.Description, body.Images)

		createdProduct, err := handler.ProductRepository.Create(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		res.Json(w, createdProduct, http.StatusCreated)
		fmt.Println("Create")
	}
}

func (handler *ProductHandler) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		product, err := handler.ProductRepository.Find(uint(id))

		res.Json(w, product, http.StatusOK)
		fmt.Println("Read")
	}
}

func (handler *ProductHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[ProductUpdateRequest](&w, r)
		if err != nil {
			return
		}
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		product, err := handler.ProductRepository.Update(&Product{
			Model:       gorm.Model{ID: uint(id)},
			Name:        body.Name,
			Description: body.Description,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		res.Json(w, product, 201)

		fmt.Println("Update")
	}
}

func (handler *ProductHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		err = handler.ProductRepository.Delete(uint(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		res.Json(w, "Delete", http.StatusOK)
		fmt.Println("DELETE")
	}
}
