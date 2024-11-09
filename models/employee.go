package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Employee struct {
    ID          primitive.ObjectID `bson:"_id,omitempty"`
    Name        string             `bson:"name"`
    NIK         string             `bson:"nik"`
    Education   string             `bson:"education"`
    JoinDate    time.Time          `bson:"join_date"`
    JobStatus   string             `bson:"job_status"`
}
