package models

import (
	"depo-bangunan/config"
	"time"
)

type Product struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"type:varchar(100)"`
	Price     int32    `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func CreateProduct(product *Product) error {
	return config.DB.Create(product).Error
}

func GetAllProducts(offset int, limit int, search string) ([]Product, int64, error) {
	var products []Product
	var count int64

	if search != "" {
			if err := config.DB.Where("name LIKE ?", "%"+search+"%").
				Model(&Product{}).
				Count(&count).
				Error; err != nil {
					return nil, 0, err
			}
	
			if err := config.DB.Where("name LIKE ?", "%"+search+"%").
				Offset(offset).Limit(limit).Find(&products).Error; err != nil {
					return nil, 0, err
			}
	} else {
			if err := config.DB.Model(&Product{}).Count(&count).Error; err != nil {
					return nil, 0, err
			}
	
			if err := config.DB.Offset(offset).Limit(limit).Find(&products).Error; err != nil {
					return nil, 0, err
			}
	}

	return products, count, nil
}

func GetProductById(id int) (Product, error) {
	var product Product
	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		return product, err
	}

	return product, nil
}

func UpdateProduct(product *Product) error {
	return config.DB.Save(product).Error
}

func DeleteProduct(id int) error {
	return config.DB.Where("id = ?", id).Delete(&Product{}).Error
}