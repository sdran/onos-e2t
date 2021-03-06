// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/e2ap/v1beta1/e2ap_containers.proto

package e2apcontainers

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

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _e_2_ap_containers_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ProtocolIeContainer001 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProtocolIeContainer001) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetValue() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProtocolIeContainer001ValidationError{
					field:  fmt.Sprintf("Value[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ProtocolIeContainer001ValidationError is the validation error returned by
// ProtocolIeContainer001.Validate if the designated constraints aren't met.
type ProtocolIeContainer001ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProtocolIeContainer001ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProtocolIeContainer001ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProtocolIeContainer001ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProtocolIeContainer001ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProtocolIeContainer001ValidationError) ErrorName() string {
	return "ProtocolIeContainer001ValidationError"
}

// Error satisfies the builtin error interface
func (e ProtocolIeContainer001ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProtocolIeContainer001.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProtocolIeContainer001ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProtocolIeContainer001ValidationError{}

// Validate checks the field values on ProtocolIeSingleContainer001 with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProtocolIeSingleContainer001) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProtocolIeSingleContainer001ValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ProtocolIeSingleContainer001ValidationError is the validation error returned
// by ProtocolIeSingleContainer001.Validate if the designated constraints
// aren't met.
type ProtocolIeSingleContainer001ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProtocolIeSingleContainer001ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProtocolIeSingleContainer001ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProtocolIeSingleContainer001ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProtocolIeSingleContainer001ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProtocolIeSingleContainer001ValidationError) ErrorName() string {
	return "ProtocolIeSingleContainer001ValidationError"
}

// Error satisfies the builtin error interface
func (e ProtocolIeSingleContainer001ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProtocolIeSingleContainer001.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProtocolIeSingleContainer001ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProtocolIeSingleContainer001ValidationError{}

// Validate checks the field values on ProtocolIeField001 with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProtocolIeField001) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProtocolIeField001ValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Criticality

	return nil
}

// ProtocolIeField001ValidationError is the validation error returned by
// ProtocolIeField001.Validate if the designated constraints aren't met.
type ProtocolIeField001ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProtocolIeField001ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProtocolIeField001ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProtocolIeField001ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProtocolIeField001ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProtocolIeField001ValidationError) ErrorName() string {
	return "ProtocolIeField001ValidationError"
}

// Error satisfies the builtin error interface
func (e ProtocolIeField001ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProtocolIeField001.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProtocolIeField001ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProtocolIeField001ValidationError{}

// Validate checks the field values on ProtocolIeContainerPair with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProtocolIeContainerPair) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetValue() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProtocolIeContainerPairValidationError{
					field:  fmt.Sprintf("Value[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ProtocolIeContainerPairValidationError is the validation error returned by
// ProtocolIeContainerPair.Validate if the designated constraints aren't met.
type ProtocolIeContainerPairValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProtocolIeContainerPairValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProtocolIeContainerPairValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProtocolIeContainerPairValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProtocolIeContainerPairValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProtocolIeContainerPairValidationError) ErrorName() string {
	return "ProtocolIeContainerPairValidationError"
}

// Error satisfies the builtin error interface
func (e ProtocolIeContainerPairValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProtocolIeContainerPair.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProtocolIeContainerPairValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProtocolIeContainerPairValidationError{}

// Validate checks the field values on ProtocolIeFieldPair with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProtocolIeFieldPair) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ProtocolIeFieldPairValidationError is the validation error returned by
// ProtocolIeFieldPair.Validate if the designated constraints aren't met.
type ProtocolIeFieldPairValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProtocolIeFieldPairValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProtocolIeFieldPairValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProtocolIeFieldPairValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProtocolIeFieldPairValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProtocolIeFieldPairValidationError) ErrorName() string {
	return "ProtocolIeFieldPairValidationError"
}

// Error satisfies the builtin error interface
func (e ProtocolIeFieldPairValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProtocolIeFieldPair.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProtocolIeFieldPairValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProtocolIeFieldPairValidationError{}

// Validate checks the field values on ProtocolIeContainerList with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProtocolIeContainerList) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetValue() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProtocolIeContainerListValidationError{
					field:  fmt.Sprintf("Value[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ProtocolIeContainerListValidationError is the validation error returned by
// ProtocolIeContainerList.Validate if the designated constraints aren't met.
type ProtocolIeContainerListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProtocolIeContainerListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProtocolIeContainerListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProtocolIeContainerListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProtocolIeContainerListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProtocolIeContainerListValidationError) ErrorName() string {
	return "ProtocolIeContainerListValidationError"
}

// Error satisfies the builtin error interface
func (e ProtocolIeContainerListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProtocolIeContainerList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProtocolIeContainerListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProtocolIeContainerListValidationError{}

// Validate checks the field values on ProtocolIeContainerPairList with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProtocolIeContainerPairList) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetValue() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ProtocolIeContainerPairListValidationError{
					field:  fmt.Sprintf("Value[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ProtocolIeContainerPairListValidationError is the validation error returned
// by ProtocolIeContainerPairList.Validate if the designated constraints
// aren't met.
type ProtocolIeContainerPairListValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProtocolIeContainerPairListValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProtocolIeContainerPairListValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProtocolIeContainerPairListValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProtocolIeContainerPairListValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProtocolIeContainerPairListValidationError) ErrorName() string {
	return "ProtocolIeContainerPairListValidationError"
}

// Error satisfies the builtin error interface
func (e ProtocolIeContainerPairListValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProtocolIeContainerPairList.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProtocolIeContainerPairListValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProtocolIeContainerPairListValidationError{}
