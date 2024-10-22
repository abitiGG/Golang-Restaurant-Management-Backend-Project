package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	Name       string             `json:"Name" binding:"required"`
	Category   string             `json:"Category" binding:"required"`
	Price      float64            `json:"Price" binding:"required"`
	FoodImage  string             `json:"FoodImage" binding:"required"`
	Menu_id    string             `json:"Menu_id" binding:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Food_id    string             `json:"food_id"`
}
