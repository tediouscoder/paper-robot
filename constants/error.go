package constants

import "errors"

// Errors that paper robot used.
var (
	ErrRequiredFiledMissing = errors.New("required filed missing")
	ErrInvalidPaperURL      = errors.New("invalid paper url")
	ErrNotImplemented       = errors.New("not implemented")
	ErrPaperNotExist        = errors.New("paper not exist")
)
