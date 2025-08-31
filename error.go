package core

import "strings"

type ErrorCode uint8

// Error interface describes an application error
type Error interface {
	error
	Code() ErrorCode
	CodeValue() string
	Cause() error
	Context() string
	SetContext(string)
	MetadataValue(string) string
	Metadata() map[string]string
}

// ErrorImpl implements an Error
type ErrorImpl struct {
	code      ErrorCode
	codeValue string
	context   string
	cause     error
	metadata  map[string]string
	message   string
}

var _ Error = (*ErrorImpl)(nil)

// Error returns a descriptor string that encapsulates all the error's metadata. This
// satisfies the 'error' interface
func (e *ErrorImpl) Error() string {

	var sb strings.Builder

	// add the error code
	sb.WriteString("[")
	sb.WriteString(e.codeValue)
	sb.WriteString("] ")

	// add the optional context
	if e.context != "" {
		sb.WriteString(e.context)
		sb.WriteString(": ")
	}

	// add the message
	sb.WriteString(e.message)

	// add the optional cause
	if e.cause != nil {
		sb.WriteString("; Cause=")
		sb.WriteString(e.cause.Error())
	}
	return sb.String()
}

// Code returns the associated error code
func (e *ErrorImpl) Code() ErrorCode {
	return e.code
}

// CodeValue returns the associated error code
func (e *ErrorImpl) CodeValue() string {
	return e.codeValue
}

// Cause returns the optional underlying error
func (e *ErrorImpl) Cause() error {
	return e.cause
}

// Context returns this error's optional context
func (e *ErrorImpl) Context() string {
	return e.context
}

// SetContext identifies the context of this error
func (e *ErrorImpl) SetContext(ctx string) {
	e.context = ctx
}

// Metadata returns this error's optional metadata. It can be nil
func (e *ErrorImpl) Metadata() map[string]string {
	return e.metadata
}

// MetadataValue returns a value from this error's metadata.
func (e *ErrorImpl) MetadataValue(key string) string {
	if e.metadata != nil {
		return e.metadata[key]
	}
	return ""
}
