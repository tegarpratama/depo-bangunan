package controllers

import (
	"depo-bangunan/models"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

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
