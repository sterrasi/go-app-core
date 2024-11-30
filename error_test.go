package core

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type BuilderFunc = func() *ErrorBuilder

var (
	internalErrorTriple = &Triple[BuilderFunc, ErrorCode, string]{
		First:  BuildInternalError,
		Second: InternalErrorCode,
		Third:  InternalErrorCodeValue}
	sysConfigErrorTriple = &Triple[BuilderFunc, ErrorCode, string]{
		First:  BuildSysConfigError,
		Second: SystemConfigurationErrorCode,
		Third:  SystemConfigurationErrorCodeValue}
	serviceUnavailableErrorTriple = &Triple[BuilderFunc, ErrorCode, string]{
		First:  BuildSvcUnavailableError,
		Second: ServiceUnavailableErrorCode,
		Third:  ServiceUnavailableErrorCodeValue}
	illegalArgumentErrorTriple = &Triple[BuilderFunc, ErrorCode, string]{
		First:  BuildIllegalArgumentError,
		Second: IllegalArgumentErrorCode,
		Third:  IllegalArgumentErrorCodeValue}
	validationErrorTriple = &Triple[BuilderFunc, ErrorCode, string]{
		First:  BuildValidationError,
		Second: ValidationErrorCode,
		Third:  ValidationErrorCodeValue}
	illegalStateErrorTriple = &Triple[BuilderFunc, ErrorCode, string]{
		First:  BuildIllegalStateError,
		Second: IllegalStateErrorCode,
		Third:  IllegalStateErrorCodeValue}
	notFoundErrorTriple = &Triple[BuilderFunc, ErrorCode, string]{
		First:  BuildNotFoundError,
		Second: NotFoundErrorCode,
		Third:  NotFoundErrorCodeValue}
	alredyExistsErrorTriple = &Triple[BuilderFunc, ErrorCode, string]{
		First:  BuildAlreadyExistsError,
		Second: AlreadyExistsErrorCode,
		Third:  AlreadyExistsErrorCodeValue}
	ioErrorTriple = &Triple[BuilderFunc, ErrorCode, string]{
		First:  BuildIOError,
		Second: IOErrorCode,
		Third:  IOErrorCodeValue}

	errorTypes = [...]*Triple[BuilderFunc, ErrorCode, string]{
		internalErrorTriple, sysConfigErrorTriple, serviceUnavailableErrorTriple, illegalArgumentErrorTriple,
		validationErrorTriple, illegalStateErrorTriple, notFoundErrorTriple, alredyExistsErrorTriple, ioErrorTriple}
)

// TestAppErrorMessageFormat iterates the builder functions and exercises each to ensure that they produce well-formed Errors
func TestAppErrorMessageFormat(t *testing.T) {
	for _, builderFunc := range errorTypes {
		testErrorBuilder(t, builderFunc.First, builderFunc.Second, builderFunc.Third)
	}
}

// testErrorBuilder tests that the given BuilderFunc and corresponding ErrorCode produces a hydrated, well-formed Error
func testErrorBuilder(t *testing.T, builderFunc BuilderFunc, code ErrorCode, codeValue string) {

	err := builderFunc().Context("some context").
		Cause(errors.New("the cause")).
		Msgf("there was a %s", "problem")
	assert.NotNil(t, err)
	assert.Equal(t, code, err.Code())
	assert.Equal(t, codeValue, err.CodeValue())
	assert.Equal(t, "the cause", err.Cause().Error())
	errMsg := fmt.Sprintf("[%s] some context: there was a problem; Cause=the cause", codeValue)
	assert.Equal(t, errMsg, err.Error())
}
