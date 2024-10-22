package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// OrderItem represents an item in an order
type OrderItem struct {
	ID         primitive.ObjectID `bson:"_id"`
	Item_ID    string             `json:"item_id" bson:"item_id"`
	Quantity   int                `json:"quantity" bson:"quantity"`
	Price      float64            `json:"price" bson:"price"`
	Unit_Price float64            `json:"unit_price" bson:"unit_price"`
	Food_id    string             `json:"food_id" bson:"food_id"`
	Updated_at time.Time          `json:"updated_at" bson:"updated_at"`
	Created_at time.Time          `json:"created_at" bson:"created_at"`
}
