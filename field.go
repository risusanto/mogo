package mogo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DateFields struct {
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type IDField struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}

func (f *IDField) SetID(id interface{}) {
	f.ID = id.(primitive.ObjectID)
}

// GetID method return model's id
func (f *IDField) GetID() interface{} {
	return f.ID
}

func (f *IDField) PrepareID(id interface{}) (interface{}, error) {
	if idStr, ok := id.(string); ok {
		return primitive.ObjectIDFromHex(idStr)
	}

	// Otherwise id must be ObjectId
	return id, nil
}

// Creating hook used here to set `created_at` field
// value on inserting new model into database.
func (f *DateFields) Creating() error {
	f.CreatedAt = time.Now().UTC()
	return nil
}

// Saving hook used here to set `updated_at` field value
// on create/update model.
func (f *DateFields) Saving() error {
	f.UpdatedAt = time.Now().UTC()
	return nil
}
