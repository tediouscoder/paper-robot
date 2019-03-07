package constants

import "errors"

// Errors that paper robot used.
var (
	ErrRequiredFiledMissing = errors.New("required filed missing")
	ErrInvalidPaperURL      = errors.New("invalid paper url")
)
