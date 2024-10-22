package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Table struct {
	// Define the fields for the Table struct
	ID                primitive.ObjectID `bson:"_id"`
	Name              string
	Table_id          string    `json:"table_id" bson:"table_id"`
	Created_at        time.Time `json:"created_at" bson:"created_at"`
	Updated_at        time.Time `json:"updated_at" bson:"updated_at"`
	Table_number      *int
	Number_of_guests  *int
	Table_description *string

	// Add other fields as necessary
}
