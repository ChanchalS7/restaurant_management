package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/ChanchalS7/restaurant_management/database"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/genproto/googleapis/bigtable/admin/table/v1"
)

var tableCollection *mongo.Collection = database.OpenCollection(database.Client, "table")

func GetTables() gin.HandlerFunc {
	return func(*gin.Context) {

	}
}
func GetTable() gin.HandlerFunc {
	return func(*gin.Context) {

	}
}
func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		var table models.Table

		if err := c.BindJSON(&table); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(table)

		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}
		table.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		table.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		table.ID = primitive.NewObjectID()
		table.Table_id = table.ID.Hex()

		result, inserErr := tableCollection.InsertOne(ctx, table)
		if inserErr != nil {
			msg := fmt.Sprintf("Table item was not created")
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return

		}
		defer cancel()
		c.JSON(http.StatusOK, result)
	}
}
func UpdateTable() gin.HandlerFunc {
	return func(*gin.Context) {

	}
}
