// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: checkout/v1/checkout.proto

package checkout_v1

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

// Validate checks the field values on AddToCartRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AddToCartRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetUser() <= 0 {
		return AddToCartRequestValidationError{
			field:  "User",
			reason: "value must be greater than 0",
		}
	}

	if m.GetSku() <= 0 {
		return AddToCartRequestValidationError{
			field:  "Sku",
			reason: "value must be greater than 0",
		}
	}

	if val := m.GetCount(); val <= 0 || val >= 65535 {
		return AddToCartRequestValidationError{
			field:  "Count",
			reason: "value must be inside range (0, 65535)",
		}
	}

	return nil
}

// AddToCartRequestValidationError is the validation error returned by
// AddToCartRequest.Validate if the designated constraints aren't met.
type AddToCartRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddToCartRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddToCartRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddToCartRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddToCartRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddToCartRequestValidationError) ErrorName() string { return "AddToCartRequestValidationError" }

// Error satisfies the builtin error interface
func (e AddToCartRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddToCartRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddToCartRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddToCartRequestValidationError{}

// Validate checks the field values on AddToCartResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *AddToCartResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// AddToCartResponseValidationError is the validation error returned by
// AddToCartResponse.Validate if the designated constraints aren't met.
type AddToCartResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddToCartResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddToCartResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddToCartResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddToCartResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddToCartResponseValidationError) ErrorName() string {
	return "AddToCartResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AddToCartResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddToCartResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddToCartResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddToCartResponseValidationError{}

// Validate checks the field values on DeleteFromCartRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteFromCartRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetUser() <= 0 {
		return DeleteFromCartRequestValidationError{
			field:  "User",
			reason: "value must be greater than 0",
		}
	}

	if m.GetSku() <= 0 {
		return DeleteFromCartRequestValidationError{
			field:  "Sku",
			reason: "value must be greater than 0",
		}
	}

	if val := m.GetCount(); val <= 0 || val >= 65535 {
		return DeleteFromCartRequestValidationError{
			field:  "Count",
			reason: "value must be inside range (0, 65535)",
		}
	}

	return nil
}

// DeleteFromCartRequestValidationError is the validation error returned by
// DeleteFromCartRequest.Validate if the designated constraints aren't met.
type DeleteFromCartRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteFromCartRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteFromCartRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteFromCartRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteFromCartRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteFromCartRequestValidationError) ErrorName() string {
	return "DeleteFromCartRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteFromCartRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteFromCartRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteFromCartRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteFromCartRequestValidationError{}

// Validate checks the field values on DeleteFromCartResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DeleteFromCartResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// DeleteFromCartResponseValidationError is the validation error returned by
// DeleteFromCartResponse.Validate if the designated constraints aren't met.
type DeleteFromCartResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteFromCartResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteFromCartResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteFromCartResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteFromCartResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteFromCartResponseValidationError) ErrorName() string {
	return "DeleteFromCartResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteFromCartResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteFromCartResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteFromCartResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteFromCartResponseValidationError{}

// Validate checks the field values on ListCartRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListCartRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetUser() <= 0 {
		return ListCartRequestValidationError{
			field:  "User",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// ListCartRequestValidationError is the validation error returned by
// ListCartRequest.Validate if the designated constraints aren't met.
type ListCartRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCartRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCartRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCartRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCartRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCartRequestValidationError) ErrorName() string { return "ListCartRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListCartRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCartRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCartRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCartRequestValidationError{}

// Validate checks the field values on ListCartResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListCartResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListCartResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for TotalPrice

	return nil
}

// ListCartResponseValidationError is the validation error returned by
// ListCartResponse.Validate if the designated constraints aren't met.
type ListCartResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCartResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCartResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCartResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCartResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCartResponseValidationError) ErrorName() string { return "ListCartResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListCartResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCartResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCartResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCartResponseValidationError{}

// Validate checks the field values on ListCartResponseItem with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListCartResponseItem) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Sku

	// no validation rules for Count

	// no validation rules for Name

	// no validation rules for Price

	return nil
}

// ListCartResponseItemValidationError is the validation error returned by
// ListCartResponseItem.Validate if the designated constraints aren't met.
type ListCartResponseItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCartResponseItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCartResponseItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCartResponseItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCartResponseItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCartResponseItemValidationError) ErrorName() string {
	return "ListCartResponseItemValidationError"
}

// Error satisfies the builtin error interface
func (e ListCartResponseItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCartResponseItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCartResponseItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCartResponseItemValidationError{}

// Validate checks the field values on PurchaseRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *PurchaseRequest) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetUser() <= 0 {
		return PurchaseRequestValidationError{
			field:  "User",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// PurchaseRequestValidationError is the validation error returned by
// PurchaseRequest.Validate if the designated constraints aren't met.
type PurchaseRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PurchaseRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PurchaseRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PurchaseRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PurchaseRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PurchaseRequestValidationError) ErrorName() string { return "PurchaseRequestValidationError" }

// Error satisfies the builtin error interface
func (e PurchaseRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPurchaseRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PurchaseRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PurchaseRequestValidationError{}

// Validate checks the field values on PurchaseResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *PurchaseResponse) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetOrderId() <= 0 {
		return PurchaseResponseValidationError{
			field:  "OrderId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// PurchaseResponseValidationError is the validation error returned by
// PurchaseResponse.Validate if the designated constraints aren't met.
type PurchaseResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PurchaseResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PurchaseResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PurchaseResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PurchaseResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PurchaseResponseValidationError) ErrorName() string { return "PurchaseResponseValidationError" }

// Error satisfies the builtin error interface
func (e PurchaseResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPurchaseResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PurchaseResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PurchaseResponseValidationError{}
