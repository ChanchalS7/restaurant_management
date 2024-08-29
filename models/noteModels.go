package models

import (
	"go.mongodb.org/mongo-driver/bson/primtive"
	"time"
)

type Note struct {
	ID         primtive.ObjectID `bson:"_id"`
	Text       *string           `json:"text"`
	Title      *string           `json:"title"`
	Created_at time.Time         `json:"created_at"`
	Updated_at time.Time         `json:"updated_at"`
	Note_id    string            `json:"note_id"`
}
