package controllers

import (
	"depo-bangunan/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateProduct(c *gin.Context) {
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
			return
	}

	product := models.Product{
		Name:    input.Name,
		Price:   input.Price,
	}

	if err := models.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data": product,
	})
}

func GetProducts(c *gin.Context) {
	search := c.Query("search")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "5")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	products, count, err := models.GetAllProducts(offset, intLimit, search)
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
		"data": products,
	})
}


func UpdateProduct(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Invalid ID"})
			return
    }

    var input models.Product
    if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
			return
    }

    product, err := models.GetProductById(id)
    if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "error", "error": "Product not found"})
			return
    }

    if input.Name != "" {
			product.Name = input.Name
    }

    if input.Price > 0 {
			product.Price = input.Price
    }

    if err := models.UpdateProduct(&product); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
			return
    }

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data": product,
	})
}

func DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Invalid ID"})
		return
	}

	_, err = models.GetProductById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "error": "Product not found"})
		return
	}

	if err := models.DeleteProduct(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
