package server

import pr "go.mongodb.org/mongo-driver/bson/primitive"

// State means the state of the authmodel such as "inactive", "active", or
// "banned".
type State string

const (
	// Inactive state means the auth model is not active.
	Inactive State = "inactive"
	// Active state means the auth model is active.
	Active = "active"
	// Banned state means the auth model is Banned.
	Banned = "banned"
)

// Auth denoates the authentication model.
type Auth struct {
	ID        pr.ObjectID
	State     State  `validate:"required"`
	UserName  string `validate:"required,dbunique"`
	Email     string `validate:"required,dbunique,email"`
	PWHash    []byte `validate:"required"`
	PWAlgo    string `validate:"required"`
	OTPSecret string `validate:"base36"`
	Recaptcha string `bson:"-" validate:"required,recaptcha=IPAddr"`
	IPAddr    string `bson:"-" validate:"required,ip"`
}
