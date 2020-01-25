package errors

// InvalidType indicates the type was not valid.
type InvalidType struct{}

// Error implements go lang's error
func (me *InvalidType) Error() string {
	return "Invalid Type"
}
