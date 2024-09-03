package controllers

import (
	"github.com/ChanchalS7/restaurant_management/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItemPack struct {
	Table_id    *string
	Order_items []models.OrderItem
}

func GetOrderItems() gin.HandlerFunc {
	return func(*gin.Context) {

	}
}
func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(*gin.Context) {

	}
}

func ItemsByOrder(id string) (OrderItems []primitive.M, err error) {
}

func GetOrderItem() gin.HandlerFunc {
	return func(*gin.Context) {

	}
}
func CreateOrderItem() gin.HandlerFunc {
	return func(*gin.Context) {

	}
}
func UpdateOrderItem() gin.HandlerFunc {
	return func(*gin.Context) {

	}
}
