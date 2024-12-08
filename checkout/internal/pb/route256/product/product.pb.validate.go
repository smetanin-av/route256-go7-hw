// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: route256/product/product.proto

package product

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
)

// Validate checks the field values on GetProductRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *GetProductRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Token

	// no validation rules for Sku

	return nil
}

// GetProductRequestValidationError is the validation error returned by
// GetProductRequest.Validate if the designated constraints aren't met.
type GetProductRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProductRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProductRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProductRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProductRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProductRequestValidationError) ErrorName() string {
	return "GetProductRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProductRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProductRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProductRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProductRequestValidationError{}

// Validate checks the field values on GetProductResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *GetProductResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Price

	return nil
}

// GetProductResponseValidationError is the validation error returned by
// GetProductResponse.Validate if the designated constraints aren't met.
type GetProductResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProductResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProductResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProductResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProductResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProductResponseValidationError) ErrorName() string {
	return "GetProductResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetProductResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProductResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProductResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProductResponseValidationError{}

// Validate checks the field values on ListSkusRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListSkusRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Token

	// no validation rules for StartAfterSku

	// no validation rules for Count

	return nil
}

// ListSkusRequestValidationError is the validation error returned by
// ListSkusRequest.Validate if the designated constraints aren't met.
type ListSkusRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSkusRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSkusRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSkusRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSkusRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSkusRequestValidationError) ErrorName() string { return "ListSkusRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListSkusRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSkusRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSkusRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSkusRequestValidationError{}

// Validate checks the field values on ListSkusResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListSkusResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ListSkusResponseValidationError is the validation error returned by
// ListSkusResponse.Validate if the designated constraints aren't met.
type ListSkusResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListSkusResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListSkusResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListSkusResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListSkusResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListSkusResponseValidationError) ErrorName() string { return "ListSkusResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListSkusResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListSkusResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListSkusResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListSkusResponseValidationError{}
