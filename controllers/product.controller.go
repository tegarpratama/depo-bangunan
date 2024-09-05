package controllers

import (
	"depo-bangunan/models"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Product
// @Description Create Product
// @Tags Products
// @Accept json
// @Produce json
// @Param name   body     string  true  "name"
// @Param price   body     string  true  "price"
// @Success 200 {object} models.SwaggerProductRes
// @Failure 400 {object} models.SwaggerErrorRes
// @Router /products [post]
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

// @Summary Get Products
// @Description Get Products
// @Tags Products
// @Accept json
// @Produce json
//  @Param  page  query string  false  "move page"  
//  @Param  limit  query string  false  "limit data"  
//  @Param  search  query string  false  "search data"  
// @Success 200 {object} models.SwaggerProductsRes
// @Failure 400 {object} models.SwaggerErrorRes
// @Router /products [get]
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

// @Summary Update Products
// @Description Update Products
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param name   body     string  true  "name"
// @Param price   body     string  true  "price"
// @Success 200 {object} models.SwaggerCreateProduct
// @Failure 400 {object} models.SwaggerErrorRes
// @Router /products/{id}/update [put]
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

// @Summary Delete Products
// @Description Delete Products
// @Tags Products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} models.SwaggerDeleteProductRes
// @Failure 400 {object} models.SwaggerErrorRes
// @Router /products/{id}/delete [delete]
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
