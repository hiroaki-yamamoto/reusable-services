package errors

import "fmt"

// NoIDFound will be raised when the Value object doens't contain ID.
type NoIDFound struct {
	Value interface{}
}

func (me *NoIDFound) Error() string {
	return fmt.Sprintf("Value %v has no id", me.Value)
}
