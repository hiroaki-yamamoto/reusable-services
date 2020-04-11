package vldfuncs

import (
	"context"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/hiroaki-yamamoto/reusable-services/adapter"
	"go.mongodb.org/mongo-driver/bson"
)

// DBUnique checks whether the value is unique on the db.
func DBUnique(adapter adapter.IAdapter) validator.FuncCtx {
	return func(ctx context.Context, fl validator.FieldLevel) bool {
		fieldName := fl.FieldName()
		query := make(bson.M)
		query[fieldName] = fl.Field().Interface()
		count, err := adapter.Count(ctx, query)
		if err != nil {
			log.Println("[Error]", err)
		}
		return count < 1
	}
}
