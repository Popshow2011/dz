package product

import (
	"dz/4-order-api/pkg/db"
)

type ProductRepository struct {
	Database *db.Db
}

func NewProductRepository(database *db.Db) *ProductRepository {
	return &ProductRepository{
		Database: database,
	}
}

func (repo *ProductRepository) Create(product *Product) (*Product, error) {
	res := repo.Database.DB.Create(product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (repo *ProductRepository) Find(id uint) (*Product, error) {
	var product Product
	res := repo.Database.DB.First(&product, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &product, nil
}

func (repo *ProductRepository) Update(product *Product) (*Product, error) {
	res := repo.Database.DB.Updates(product)
	if res.Error != nil {
		return nil, res.Error
	}
	return product, nil
}

func (repo *ProductRepository) Delete(id uint) error {
	res := repo.Database.DB.Delete(&Product{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
