package controllers

import (
	"assignment2/data"
	"assignment2/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(ctx *gin.Context) {
	ctx.JSON(http.StatusAccepted, "Ini home")
}

func CreateOrder(ctx *gin.Context) {
	var newOrder model.Order

	// Bind data dari permintaan ke variabel newOrder
	if err := ctx.ShouldBindJSON(&newOrder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "something error",
			"error":   err.Error(),
		})
		return
	}

	// Periksa apakah itemCode ada
	if len(newOrder.Items) == 0 || newOrder.Items[0].ItemCode == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "itemCode is required",
			"error":   "itemCode is empty or missing",
		})
		return
	}

	// Ambil item_code dari permintaan client
	itemCode := newOrder.Items[0].ItemCode

	var count int64
	data.DB.Model(&model.Item{}).Where("item_code = ?", itemCode).Count(&count)

	if count > 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "duplicate item_code",
			"error":   "item_code sudah ada di database",
		})
		return
	}
	// Simpan order ke dalam database
	if err := data.DB.Create(&newOrder).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "failed insert data",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "success insert data",
		"data":    newOrder,
	})

}

func GetOrders(ctx *gin.Context) {
	var getOrders []model.Order

	if err := data.DB.Preload("Items").Find(&getOrders).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"msg":     "failed get all data",
			"error":   err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "success get all data",
		"data":    getOrders,
	})
}

func GetOrderById(ctx *gin.Context) {
	var orderItem model.Order
	orderId := ctx.Param("id")

	// ambil data dari database dengan id dari request param client
	if err := data.DB.Preload("Items").First(&orderItem, "id = ?", orderId).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg":     "failed get data",
			"error":   "data with id " + orderId + " not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "success get data",
		"data":    orderItem,
	})
}

func UpdateOrder(ctx *gin.Context) {
	orderID := ctx.Param("id")

	// Ambil data order yang akan diperbarui
	var order model.Order
	if err := data.DB.Preload("Items").First(&order, "id = ?", orderID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	// Bind data dari permintaan ke variabel order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Perbarui item-item terkait dalam order
	for i := range order.Items {
		itemID := order.Items[i].ID // ID atau primary key item

		// Perbarui item di database berdasarkan ID
		if err := data.DB.Model(&model.Item{}).Where("id = ?", itemID).Updates(&order.Items[i]).Error; err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update item", "details": err.Error()})
			return
		}
	}

	// Simpan pembaruan pada data order
	if err := data.DB.Save(&order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, order)
}
