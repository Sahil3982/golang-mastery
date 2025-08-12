package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"


	"time"
)	

type Note struct {
	ID          primitive.ObjectID `bson:"_id"`
	Note_Text   *string            `bson:"note_text" validate:"required,min=1,max=500"`
	Created_At  time.Time          `bson:"created_at"`
	Updated_At  time.Time          `bson:"updated_at"`
	Created_By  primitive.ObjectID `bson:"created_by" validate:"required"`
	Updated_By  primitive.ObjectID `bson:"updated_by" validate:"required"`
}