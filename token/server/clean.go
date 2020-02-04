package server

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// CleanRottedToken removes rotted token from the DB.
func (me *Server) CleanRottedToken() {
	ctx, stop := me.TimeoutContext(context.Background())
	defer stop()
	me.adapter.DeleteMany(ctx, bson.M{
		"expires": bson.M{"$lt": me.Now()},
	})
}
