package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct {
	ID             primitive.ObjectID `bson:"_id"`
	Invoice_Number *string            `bson:"invoice_number" validate:"required"`
	Customer_Name  *string            `bson:"customer_name" validate:"required"`
	Customer_Email *string            `bson:"customer_email" validate:"required,email"`
	Customer_Phone *string            `bson:"customer_phone" validate:"required"`
	Total_Amount   *float64           `bson:"total_amount" validate:"required"`
	Payment_Status *string            `bson:"payment_status" validate:"required"`
	Created_At     time.Time          `bson:"created_at"`
	Updated_At     time.Time          `bson:"updated_at"`
	Order_ID       primitive.ObjectID `bson:"order_id" validate:"required"`
}
