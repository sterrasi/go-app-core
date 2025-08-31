package core

const UnknownErrorCode ErrorCode = 0

// InternalErrorCode relates to a general internal server error that should be avoided if
// a more specific one can be chosen
const InternalErrorCode ErrorCode = 1
const InternalErrorCodeValue string = "internal"

func BuildInternalError() *ErrorBuilder {
	return NewErrorBuilder(InternalErrorCode, InternalErrorCodeValue)
}
func NewInternalError(format string, args ...any) Error {
	return BuildInternalError().Msgf(format, args...)
}

// SystemConfigurationErrorCode signifies a server error that keeps the server from starting.  It is
// related to an issue that can be fixed in the software's configuration
const SystemConfigurationErrorCode ErrorCode = 2
const SystemConfigurationErrorCodeValue string = "system-configuration"

func BuildSysConfigError() *ErrorBuilder {
	return NewErrorBuilder(SystemConfigurationErrorCode, SystemConfigurationErrorCodeValue)
}
func NewSysConfigError(format string, args ...any) Error {
	return BuildSysConfigError().Msgf(format, args...)
}

// ServiceUnavailableErrorCode signifies that either the server or one of its dependencies is not able to service the
// request.
const ServiceUnavailableErrorCode ErrorCode = 3
const ServiceUnavailableErrorCodeValue string = "service-unavailable"

func BuildSvcUnavailableError() *ErrorBuilder {
	return NewErrorBuilder(ServiceUnavailableErrorCode, ServiceUnavailableErrorCodeValue)
}
func NewSvcUnavailableError(format string, args ...any) Error {
	return BuildSvcUnavailableError().Msgf(format, args...)
}

// IllegalArgumentErrorCode relates to an internal server error that means an internal argument check failed.  This type
// of error usually signifies a bug in the software
const IllegalArgumentErrorCode ErrorCode = 4
const IllegalArgumentErrorCodeValue string = "illegal-argument"

func BuildIllegalArgumentError() *ErrorBuilder {
	return NewErrorBuilder(IllegalArgumentErrorCode, IllegalArgumentErrorCodeValue)
}
func NewIllegalArgumentError(format string, args ...any) Error {
	return BuildIllegalArgumentError().Msgf(format, args...)
}

// ValidationErrorCode signifies a client level error that means the data provided by the client to the server
// is invalid
const ValidationErrorCode ErrorCode = 5
const ValidationErrorCodeValue string = "validation"

func BuildValidationError() *ErrorBuilder {
	return NewErrorBuilder(ValidationErrorCode, ValidationErrorCodeValue)
}
func NewValidationError(format string, args ...any) Error {
	return BuildValidationError().Msgf(format, args...)
}

// IllegalStateErrorCode relates to a client level error that signifies the operation asked of the server
// cannot be performed because it is not in the proper state
const IllegalStateErrorCode ErrorCode = 6
const IllegalStateErrorCodeValue string = "illegal-state"

func BuildIllegalStateError() *ErrorBuilder {
	return NewErrorBuilder(IllegalStateErrorCode, IllegalStateErrorCodeValue)
}
func NewIllegalStateError(format string, args ...any) Error {
	return BuildIllegalStateError().Msgf(format, args...)
}

// NotFoundErrorCode related to an error where a referenced resource does not exist
const NotFoundErrorCode ErrorCode = 7
const NotFoundErrorCodeValue string = "not-found"

func BuildNotFoundError() *ErrorBuilder {
	return NewErrorBuilder(NotFoundErrorCode, NotFoundErrorCodeValue)
}
func NewNotFoundError(format string, args ...any) Error {
	return BuildNotFoundError().Msgf(format, args...)
}

// AlreadyExistsErrorCode related to an operation that tries to create an entity that
// already exists in the system
const AlreadyExistsErrorCode ErrorCode = 8
const AlreadyExistsErrorCodeValue string = "already-exists"

func BuildAlreadyExistsError() *ErrorBuilder {
	return NewErrorBuilder(AlreadyExistsErrorCode, AlreadyExistsErrorCodeValue)
}
func NewAlreadyExistsError(format string, args ...any) Error {
	return BuildAlreadyExistsError().Msgf(format, args...)
}

// IOErrorCode relates to an error while trying to interface with a stream of bytes
const IOErrorCode ErrorCode = 9
const IOErrorCodeValue string = "io"

func BuildIOError() *ErrorBuilder {
	return NewErrorBuilder(IOErrorCode, IOErrorCodeValue)
}
func NewIOError(format string, args ...any) Error {
	return BuildIOError().Msgf(format, args...)
}
