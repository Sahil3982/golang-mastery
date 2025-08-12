package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderItem struct {
	ID          primitive.ObjectID `bson:"_id"`
	OrderID     primitive.ObjectID `bson:"order_id" validate:"required"`
	FoodID      primitive.ObjectID `bson:"food_id" validate:"required"`
	Quantity    *int               `bson:"quantity" validate:"required,min=1"`
	Price       *float64           `bson:"price" validate:"required"`
	CreatedAt   primitive.DateTime `bson:"created_at"`
	UpdatedAt   primitive.DateTime `bson:"updated_at"`
}