package interfaces

import "io"

// ITemplate represents an interface of go template system.
type ITemplate interface {
	Execute(wr io.Writer, data interface{}) error
}
