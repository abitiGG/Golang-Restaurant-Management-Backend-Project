package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	// ... existing fields ...
	Table_id   *primitive.ObjectID `bson:"table_id,omitempty"`
	ID         primitive.ObjectID  `bson:"_id,omitempty"`
	Created_at time.Time           `bson:"created_at"`
	Updated_at time.Time           `bson:"updated_at"`
	Order_id   string              `bson:"order_id"`
	Order_date time.Time           `bson:"order_date"`
	// ... existing fields ...
}
