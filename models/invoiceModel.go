package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson/primtive"
)

type Invoice struct {
ID					primitive.ObjectID		`bson:"_id"`
Invoice ID			string					`json:"invoice_id"`
Order_id			string 					`json:"order_id"`
Pyament_method   	*string					`json:"payment_method" validate:"eq=CARD|eq=CASH|eq="`
Payment_status		*string					`json::"payment_due_date:`
Payment_due_date	*string 				`json:"created_at"`
Created_at			time.Time 				`json:"updated_at"`
Updated_at			time.Time 


}