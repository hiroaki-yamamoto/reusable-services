package vldfuncs

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/hiroaki-yamamoto/reusable-services/adapter"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
)

// DBUnique checks whether the value is unique on the db.
func DBUnique(logger *zap.Logger, adapter adapter.IAdapter) validator.FuncCtx {
	return func(ctx context.Context, fl validator.FieldLevel) bool {
		fieldName := fl.FieldName()
		query := make(bson.M)
		query[fieldName] = fl.Field().Interface()
		count, err := adapter.Count(ctx, query)
		if err != nil {
			logger.Error(
				"DBUnique has error", zap.Any("query", query), zap.Error(err),
			)
		}
		return count < 1
	}
}
