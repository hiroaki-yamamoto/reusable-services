package errors

import "fmt"

// NoTemplateFound structs an error that there's no template in the service.
type NoTemplateFound struct {
	TmpName string
}

// Error implements error interface.
func (me *NoTemplateFound) Error() string {
	return fmt.Sprintf("Template \"%s\" not found", me.TmpName)
}
