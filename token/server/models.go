package server

import (
	"time"

	"github.com/hiroaki-yamamoto/reusable-services/random"
	"github.com/hiroaki-yamamoto/reusable-services/token/rpc"
	pr "go.mongodb.org/mongo-driver/bson/primitive"
)

// Model represents a model to store a token.
type Model struct {
	ID pr.ObjectID `bson:"_id"`
	*rpc.Token
	Email   string `validator:"email"`
	Expires time.Time
}

// Generate token generats a new token for Token.Token
func (me *Model) GenerateToken(size int, txtMap ...string) (err error) {
	me.Token.Token, err = random.GenTxt(size, txtMap...)
	return
}
