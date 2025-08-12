package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                 primitive.ObjectID  `bson:"_id"`
	Username           *string             `bson:"username" validate:"required,min=2,max=100"`
	Email              *string             `bson:"email" validate:"required,email"`
	Password           *string             `bson:"password" validate:"required,min=6"`
	Role               *string             `bson:"role" validate:"required"`
	Created_At         primitive.DateTime  `bson:"created_at"`
	Updated_At         primitive.DateTime  `bson:"updated_at"`
	Is_Active          *bool               `bson:"is_active" validate:"required"`
	Is_Deleted         *bool               `bson:"is_deleted" validate:"required"`
	Phone              *string             `bson:"phone" validate:"required"`
	Address            *string             `bson:"address" validate:"required"`
	Profile_Image_URL  *string             `bson:"profile_image_url"`
	Reset_Token        *string             `bson:"reset_token"`
	Reset_Token_Expiry *primitive.DateTime `bson:"reset_token_expiry"`
}
