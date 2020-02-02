package errors

// NotFound indicates the model wasn't found / rotted.
type NotFound struct{}

func (me *NotFound) Error() string {
	return "Model Not Found"
}
