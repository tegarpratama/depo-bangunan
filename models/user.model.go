package models

import (
	"depo-bangunan/config"
	"time"
)

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Name      string `json:"name" gorm:"type:varchar(100)"`
	Email  		string `json:"email" gorm:"type:varchar(100)"`
	Role  		string `json:"role" gorm:"type:varchar(20)"`
	Password  string `json:"-" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Register struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
}

type Login struct {
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required"`
}

type CreateCustomerReq struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
}

type UpdateCustomerReq struct {
	Name    string `json:"name"`
	Email   string `json:"email" binding:"email"`
}

type UserLoggedIn struct {
	ID int  `json:"id"`
	Email string `json:"email"`
	Role string `json:"role"`
	Token string `json:"token"`
}

type SwaggerRegisterRes struct {
    Status string  `json:"status" example:"ok"`
    Data   User `json:"data"`
}

type SwaggerLoginRes struct {
    Status string  `json:"status" example:"ok"`
    Data   UserLoggedIn `json:"data"`
}

type SwaggerErrorRes struct {
    Status string  	`json:"status" example:"error"`
    Message   string `json:"message"`
}

type SwaggerUserRes struct {
	Status string  `json:"status" example:"ok"`
    CurrentPate int `json:"current_page"`
    TotalPage int `json:"total_page"`
    TotalData int `json:"total_data"`
	Data   User `json:"data"`	
}

type SwaggerDeleteUserRes struct {
	Status string  `json:"status" example:"ok"`
}

func GetCustomerByEmail(email string) User {
	var customer User
	if err := config.DB.Where("email = ?", email).First(&customer).Error; err != nil {
			return customer
	}

	return customer
}

func CreateCustomer(customer *User) error {
	return config.DB.Create(customer).Error
}

func GetAllCustomers(offset int, limit int, search string) ([]User, int64, error) {
	var customers []User
	var count int64

	if search != "" {
			if err := config.DB.Where("name LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%").
				Where("role = ?", "customer").
				Model(&User{}).
				Count(&count).
				Error; err != nil {
					return nil, 0, err
			}
	
			if err := config.DB.Where("name LIKE ? OR email LIKE ?", "%"+search+"%", "%"+search+"%").
				Where("role = ?", "customer").
				Offset(offset).Limit(limit).Find(&customers).Error; err != nil {
					return nil, 0, err
			}
	} else {
			if err := config.DB.Where("role = ?", "customer").Model(&User{}).Count(&count).Error; err != nil {
					return nil, 0, err
			}
	
			if err := config.DB.Where("role = ?", "customer").Offset(offset).Limit(limit).Find(&customers).Error; err != nil {
					return nil, 0, err
			}
	}

	return customers, count, nil
}

func GetCustomerByID(id int) (User, error) {
	var customer User
	if err := config.DB.Where("id = ?", id).First(&customer).Error; err != nil {
		return customer, err
	}

	return customer, nil
}

func UpdateCustomer(customer *User) error {
	return config.DB.Save(customer).Error
}

func DeleteCustomer(id int) error {
	return config.DB.Where("id = ?", id).Delete(&User{}).Error
}