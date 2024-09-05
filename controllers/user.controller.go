package controllers

import (
	"depo-bangunan/helpers"
	"depo-bangunan/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Customer
// @Description Create new customer
// @Tags Customers
// @Accept json
// @Produce json
// @Param name   body     string  true  "name"
// @Param email   body     string  true  "email"
// @Success 200 {object} models.SwaggerRegisterRes
// @Failure 400 {object} models.SwaggerErrorRes
// @Router /customers [post]
func CreateCustomers(c *gin.Context) {
	var input models.CreateCustomerReq

	if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
			return
	}

	customerExist := models.GetCustomerByEmail(input.Email) 
	if customerExist.ID > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Email already registered"})
		return
	}

	hashedPassword, err := helpers.HashPassword("password")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	customer := models.User{
		Name:    input.Name,
		Email:   input.Email,
		Password: hashedPassword,
		Role: "customer",
	}

	if err := models.CreateCustomer(&customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
	"status": "ok",
	"data": &models.User{
		ID: customer.ID,
		Name: customer.Name,
		Email: customer.Email,
		Role: customer.Role,
	},
})
}

// @Summary Get All Customers
// @Description  Get All Customers
// @Tags Customers
// @Accept json
// @Produce json
//  @Param  page  query string  false  "move page"  
//  @Param  limit  query string  false  "limit data"  
//  @Param  search  query string  false  "search data"  
// @Success 200 {object} models.SwaggerUserRes
// @Failure 400 {object} models.SwaggerErrorRes
// @Router /customers [get]
func GetCustomers(c *gin.Context) {
	search := c.Query("search")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "5")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	customers, count, err := models.GetAllCustomers(offset, intLimit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		return
	}

	totalPage := int(math.Ceil(float64(count) / float64(intLimit)))

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"current_page": intPage,
		"total_page": totalPage,
		"total_data": count,
		"data": customers,
	})
}

// @Summary Get Detail Customer
// @Description  Get Detail Customer
// @Tags Customers
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.SwaggerRegisterRes
// @Failure 400 {object} models.SwaggerErrorRes
// @Router /customers/{id}/detail [get]
func DetailCustomers(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Invalid ID"})
			return
    }

    customer, err := models.GetCustomerByID(id)
    if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "error": "Customer not found"})
			return
    }

    c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data": customer,
	})
}

// @Summary Update Customer
// @Description  Update Customer
// @Tags Customers
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param name   body     string  true  "name"
// @Param email   body     string  true  "email"
// @Success 200 {object} models.SwaggerRegisterRes
// @Failure 400 {object} models.SwaggerErrorRes
// @Router /customers/{id}/update [put]
func UpdateCustomer(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Invalid ID"})
			return
    }

    var input models.UpdateCustomerReq
    if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
			return
    }

    customer, err := models.GetCustomerByID(id)
    if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "error": "Customer not found"})
			return
    }

    if input.Name != "" {
			customer.Name = input.Name
    }

    if input.Email != "" {
			customer.Email = input.Email
    }

    if err := models.UpdateCustomer(&customer); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
			return
    }

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data": customer,
	})
}

// @Summary Delete Custtomer
// @Description  Delete Custtomer
// @Tags Customers
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} models.SwaggerDeleteUserRes
// @Failure 400 {object} models.SwaggerErrorRes
// @Router /customers/{id}/delete [delete]
func DeleteCustomer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Invalid ID"})
			return
	}

	currentUser, _ := c.Get("currentUser")
	userLoggedIn, _ := currentUser.(*models.UserLoggedIn)
	
	if userLoggedIn.Role == "customer" {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "error": "Only role admin can delete customer"})
		return
	}

	_, err = models.GetCustomerByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "error": "Customer not found"})
		return
	}

	if err := models.DeleteCustomer(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
