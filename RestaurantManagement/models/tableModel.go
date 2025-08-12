package models	

import (
	"go.mongodb.org/mongo-driver/bson/primitive"	
	"time"
)

type Table struct {
	ID          primitive.ObjectID `bson:"_id"`
	Table_Number *string            `bson:"table_number" validate:"required"`
	Seats       *int               `bson:"seats" validate:"required"`
	Status      *string            `bson:"status" validate:"required,oneof=available occupied"`
	Created_At  time.Time          `bson:"created_at"`	
	Updated_At  time.Time          `bson:"updated_at"`
	Table_ID    *string            `bson:"table_id" validate:"required"`
	Reservation_ID *string         `bson:"reservation_id"`
}