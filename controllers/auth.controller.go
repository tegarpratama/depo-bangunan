package controllers

import (
	"depo-bangunan/config"
	"depo-bangunan/helpers"
	"depo-bangunan/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var input models.Register
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	customerExist := models.GetCustomerByEmail(input.Email) 
	if customerExist.ID > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Email already registered"})
		return
	}
	
	if input.Password != input.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Passwords not match"})
		return
	}

	hashedPassword, err := helpers.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	user := models.User{
		Name:    input.Name,
		Email:   input.Email,
		Password: hashedPassword,
		Role: "customer",
	}

	if err := models.CreateCustomer(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data": &models.User{
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
			Role: user.Role,
		},
	})
}

func Login(c *gin.Context) {
	var input models.Login
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	userExist := models.GetCustomerByEmail(input.Email) 
	if userExist.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Wrong email or password"})
		return
	}

	var userLoggedIn models.UserLoggedIn
	userLoggedIn.ID = int(userExist.ID)
	userLoggedIn.Email = userExist.Email
	userLoggedIn.Role = userExist.Role

	if err := helpers.VerifyPassword(userExist.Password, input.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid email or Password"})
		return
	}

	access_token, err := helpers.CreateToken(config.ENV.AccessTokenExpiresIn, userLoggedIn, config.ENV.AccessTokenPrivateKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": err.Error()})
		return
	}

	userLoggedIn.Token = access_token

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data": userLoggedIn,
	})
}