// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

func IsUnspecified(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UNSPECIFIED.String() && e.Code == 500
}

func ErrorUnspecified(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_UNSPECIFIED.String(), fmt.Sprintf(format, args...))
}

func IsNoResult(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NO_RESULT.String() && e.Code == 404
}

func ErrorNoResult(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_NO_RESULT.String(), fmt.Sprintf(format, args...))
}