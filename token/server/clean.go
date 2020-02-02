package server

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// CleanRottedToken removes rotted token from the DB.
func (me *Server) CleanRottedToken() {
	ctx, stop := me.TimeoutContext(context.Background())
	defer stop()
	me.adapter.DeleteMany(ctx, bson.M{
		"expires": bson.M{
			"$lt": time.Now().UTC(),
		},
	})
}
