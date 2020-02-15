package server

import (
	"time"

	"github.com/hiroaki-yamamoto/reusable-services/random"
	"github.com/hiroaki-yamamoto/reusable-services/token/rpc"
	pr "go.mongodb.org/mongo-driver/bson/primitive"
)

// Model represents a model to store a token.
type Model struct {
	*rpc.Token
	ID      pr.ObjectID `bson:"_id"`
	Email   string      `validator:"email"`
	Expires time.Time
}

// GenerateToken generats a new token for Token.Token
func (me *Model) GenerateToken(size int, txtMap ...string) (err error) {
	me.Token.Token, err = random.GenTxt(size, txtMap...)
	return
}
