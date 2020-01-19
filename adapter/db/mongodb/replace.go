package mongodb

import (
	"context"
	"errors"
)

// Replace the matched document by query, with doc.
func (me *Mongo) Replace(
	ctx context.Context,
	query interface{},
	doc interface{},
) (interface{}, error) {
	return nil, errors.New("Not Implemented Yet")
}
