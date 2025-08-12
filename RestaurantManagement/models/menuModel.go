package models


import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Menu struct {
	ID          primitive.ObjectID `bson:"_id"`
	Menu_Name   *string            `bson:"menu_name" validate:"required,min=2,max=100"`
	Menu_Description *string      `bson:"menu_description"`
	Created_At  time.Time          `bson:"created_at"`
	Updated_At  time.Time          `bson:"updated_at"`
	Menu_ID     *string            `bson:"menu_id" validate:"required"`
	Restaurant_ID *string         `bson:"restaurant_id" validate:"required"`
}