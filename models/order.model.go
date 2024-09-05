package models

import (
	"depo-bangunan/config"
	"time"
)

type Order struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	OrderNumber string `json:"order_number"`
	UserID    *uint	  	`json:"user_id"`
	ProductID *uint	  	`json:"product_id"`
	Qty		    int32     `json:"qty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User   		User  		`json:"user" gorm:"foreignkey:UserID"`
	Product   Product  	`json:"product" gorm:"foreignkey:ProductID"`
}


type CreateOrderReq struct {
	ProductID *uint 	`json:"product_id"`
	Qty			int32 `json:"qty" binding:"required"`
}

type OrderNumberRes struct {
	OrderNumber string `json:"order_number"`
}

type UpdatedOrder struct {
	OrderNumber string `json:"order_number"`
	ProductID string `json:"product_id"`
	Qty string `json:"qty"`
}

type SwaggerCreateOrderRes struct {
	Status string  `json:"status" example:"ok"`
    Data   OrderNumberRes `json:"data"`
}

type SwaggerUpdateOrderRes struct {
	Status string  `json:"status" example:"ok"`
    Data   UpdatedOrder `json:"data"`
}
type SwaggerOrderRes struct {
	Status string  `json:"status" example:"ok"`
    CurrentPate int `json:"current_page"`
    TotalPage int `json:"total_page"`
    TotalData int `json:"total_data"`
	Data   OrderNumberRes `json:"data"`	
}

type SwaggerDetailOrderRes struct {
	Status string  `json:"status" example:"ok"`
    Data   Order `json:"data"`
}

type SwaggerDeleteOrderRes struct {
	Status string  `json:"status" example:"ok"`
}

func GetAllOrders(offset int, limit int, search string) ([]Order, int64, error) {
	var orders []Order
	var count int64

	if search != "" {
		if err := config.DB.Where("order_number LIKE ?", "%"+search+"%").Model(&Order{}).Count(&count).Error; err != nil {
			return nil, 0, err
		}

		if err := config.DB.Where("order_number LIKE ?", "%"+search+"%").Preload("User").Preload("Product").Offset(offset).Limit(limit).Find(&orders).Error; err != nil {
			return nil, 0, err
		}
	} else {
		if err := config.DB.Model(&Order{}).Count(&count).Error; err != nil {
			return nil, 0, err
		}

		if err := config.DB.Preload("User").Preload("Product").Offset(offset).Limit(limit).Find(&orders).Error; err != nil {
			return nil, 0, err
		}
	}

	return orders, count, nil
}


// 	return orders, count, nil
// }

func GetOrderByID(id int) (Order, error) {
	var order Order
	if err := config.DB.Preload("User").Preload("Product").Where("id = ?", id).First(&order).Error; err != nil {
			return order, err
	}
	return order, nil
}

func CreateOrder(order *Order) error {
	return config.DB.Create(order).Error
}

func UpdateOrder(order *Order) error {
	return config.DB.Save(order).Error
}

func DeleteOrder(id int) error {
	return config.DB.Where("id = ?", id).Delete(&Order{}).Error
}

// func SearchOrders(query string, offset int, limit int) ([]Order, int64, error) {
// 	var orders []Order
// 	var count int64

// 	if err := config.DB.Where("order_number LIKE ? OR status LIKE ?", "%"+query+"%", "%"+query+"%").
// 			Model(&Order{}).Count(&count).Error; err != nil {
// 			return nil, 0, err
// 	}

// 	if err := config.DB.Where("order_number LIKE ? OR status LIKE ?", "%"+query+"%", "%"+query+"%").
// 			Preload("Customer").Offset(offset).Limit(limit).Find(&orders).Error; err != nil {
// 			return nil, 0, err
// 	}

// 	return orders, count, nil
// }