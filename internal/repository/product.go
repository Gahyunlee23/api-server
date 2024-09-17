package repository

import (
	"api-server/internal/config"
	"api-server/internal/models"
)

func CreateProduct(product *models.Product) error {
	return config.DB.Create(product).Error
}

func GetProductById(id uint) (*models.Product, error) {
	var product models.Product
	err := config.DB.First(&product, id).Error
	return &product, err
}

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := config.DB.Find(&products).Error
	return products, err
}

func UpdateProduct(product *models.Product) error {
	return config.DB.Save(product).Error
}

func DeleteProduct(id uint) error {
	return config.DB.Delete(&models.Product{}, id).Error
}
