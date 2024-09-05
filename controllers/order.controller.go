package controllers

import (
	"depo-bangunan/models"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Create Order
// @Description Create Order
// @Tags Orders
// @Accept json
// @Produce json
// @Param product_id   path     string  true  "product_id"
// @Param qty   path     string  true  "qty"
// @Success 200 {object} models.SwaggerCreateOrderRes
// @Failure 400 {object} models.SwaggerErrorRes
// @Router /orders [post]
func CreateOrder(c *gin.Context) {
	var input models.CreateOrderReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	_, err := models.GetProductById(int(*input.ProductID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "error": "Product not found"})
		return
	}

	currentTime := time.Now()
	timestamp := currentTime.Format("20060102150405")
	currentUser, _ := c.Get("currentUser")
	userLoggedIn, _ := currentUser.(*models.UserLoggedIn)
	customerIDStr := strconv.Itoa(userLoggedIn.ID)

	userId := uint(userLoggedIn.ID) 
	userIdUint := &userId     

	order := models.Order{
		OrderNumber: timestamp + customerIDStr,
		UserID: userIdUint,
		ProductID: input.ProductID,
		Qty: input.Qty,
	}

	if err := models.CreateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data": map[string]string{
			"order_number": order.OrderNumber,
		},
	})
}

// @Summary Get all orders
// @Description Get all orders
// @Tags Orders
// @Accept json
// @Produce json
//  @Param  page  query string  false  "move page"  
//  @Param  limit  query string  false  "limit data"  
//  @Param  search  query string  false  "search data" 
// @Success 200 {object} models.SwaggerOrderRes
// @Failure 400 {object} models.SwaggerErrorRes
// @Router /orders [get]
func GetOrders(c *gin.Context) {
	search := c.Query("search")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "5")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	orders, count, err := models.GetAllOrders(offset, intLimit, search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		return
	}

	totalPage := int(math.Ceil(float64(count) / float64(intLimit)))

	c.JSON(http.StatusOK, gin.H{
		"status":       "ok",
		"current_page": intPage,
		"total_page":   totalPage,
		"total_data":   count,
		"data":         orders,
	})
}

// @Summary Get Detail Order
// @Description  Get Detail Order
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} models.SwaggerDetailOrderRes
// @Failure 400 {object} models.SwaggerErrorRes
// @Router /orders/{id}/detail [get]
func DetailOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Invalid ID"})
		return
	}

	order, err := models.GetOrderByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "error": "Order not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data":   order,
	})
}

// @Summary Update Order
// @Description  Update Order
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} models.SwaggerUpdateOrderRes
// @Failure 400 {object} models.SwaggerErrorRes
// @Router /orders/{id}/update [put]
func UpdateOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Invalid ID"})
		return
	}

	var input models.CreateOrderReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	order, err := models.GetOrderByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "error": "Order not found"})
		return
	}

	_, err = models.GetProductById(int(*input.ProductID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "error", "error": "Product not found"})
		return
	}

	if *input.ProductID != 0 {
	    order.ProductID = input.ProductID
	}

	if input.Qty != 0 {
		order.Qty = input.Qty
	}


	if err := models.UpdateOrder(&order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data": map[string]any{
			"order_number": order.OrderNumber,
			"product_id": input.ProductID,
			"qty": input.Qty,
		},
	})
}

// @Summary Delete Order
// @Description  Delete Order
// @Tags Orders
// @Accept json
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} models.SwaggerDeleteOrderRes
// @Failure 400 {object} models.SwaggerErrorRes
// @Router /orders/{id}/delete [delete]
func DeleteOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": "Invalid ID"})
		return
	}

	if err := models.DeleteOrder(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
