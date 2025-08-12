package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Food struct {
	ID               primitive.ObjectID `bson:"_id"`
	Food_Name        *string            `bson:"food_name" validate:"required,min=2,max=100"`
	Food_Description *string            `bson:"food_description"`
	Price            *float64           `bson:"price" validate:"required"`
	Food_Image_URL   *string            `bson:"food_image_url" validate:"required"`
	Created_At       time.Time          `bson:"created_at"`
	Updated_At       time.Time          `bson:"updated_at"`
	Food_id          *string            `bson:"food_id"`
	Menu_id          *string            `bson:"menu_id" validate:"required"`
}
