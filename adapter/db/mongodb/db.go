package mongodb

import "go.mongodb.org/mongo-driver/mongo"

// Mongo indicates the database adapter between the app and Mongodb.
type Mongo struct {
	col *mongo.Collection
}

// New initializes Mongo Structure.
func New(col *mongo.Collection) *Mongo {
	return &Mongo{col: col}
}
