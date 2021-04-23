// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: proto/helloworld/helloworld.proto

package helloworld

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

// Validate checks the field values on HelloRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned. When asked to return all errors, validation continues after
// first violation, and the result is a list of violation errors wrapped in
// HelloRequestMultiError, or nil if none found. Otherwise, only the first
// error is returned, if any.
func (m *HelloRequest) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetName()) > 15 {
		err := HelloRequestValidationError{
			field:  "Name",
			reason: "value length must be at most 15 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return HelloRequestMultiError(errors)
	}
	return nil
}

// HelloRequestMultiError is an error wrapping multiple validation errors
// returned by HelloRequest.Validate(true) if the designated constraints
// aren't met.
type HelloRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HelloRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HelloRequestMultiError) AllErrors() []error { return m }

// HelloRequestValidationError is the validation error returned by
// HelloRequest.Validate if the designated constraints aren't met.
type HelloRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HelloRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HelloRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HelloRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HelloRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HelloRequestValidationError) ErrorName() string { return "HelloRequestValidationError" }

// Error satisfies the builtin error interface
func (e HelloRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHelloRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HelloRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HelloRequestValidationError{}

// Validate checks the field values on HelloReply with the rules defined in the
// proto definition for this message. If any rules are violated, an error is
// returned. When asked to return all errors, validation continues after first
// violation, and the result is a list of violation errors wrapped in
// HelloReplyMultiError, or nil if none found. Otherwise, only the first error
// is returned, if any.
func (m *HelloReply) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Message

	if len(errors) > 0 {
		return HelloReplyMultiError(errors)
	}
	return nil
}

// HelloReplyMultiError is an error wrapping multiple validation errors
// returned by HelloReply.Validate(true) if the designated constraints aren't met.
type HelloReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HelloReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HelloReplyMultiError) AllErrors() []error { return m }

// HelloReplyValidationError is the validation error returned by
// HelloReply.Validate if the designated constraints aren't met.
type HelloReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HelloReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HelloReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HelloReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HelloReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HelloReplyValidationError) ErrorName() string { return "HelloReplyValidationError" }

// Error satisfies the builtin error interface
func (e HelloReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHelloReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HelloReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HelloReplyValidationError{}

// Validate checks the field values on Point with the rules defined in the
// proto definition for this message. If any rules are violated, an error is
// returned. When asked to return all errors, validation continues after first
// violation, and the result is a list of violation errors wrapped in
// PointMultiError, or nil if none found. Otherwise, only the first error is
// returned, if any.
func (m *Point) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Latitude

	// no validation rules for Longitude

	if len(errors) > 0 {
		return PointMultiError(errors)
	}
	return nil
}

// PointMultiError is an error wrapping multiple validation errors returned by
// Point.Validate(true) if the designated constraints aren't met.
type PointMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PointMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PointMultiError) AllErrors() []error { return m }

// PointValidationError is the validation error returned by Point.Validate if
// the designated constraints aren't met.
type PointValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PointValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PointValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PointValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PointValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PointValidationError) ErrorName() string { return "PointValidationError" }

// Error satisfies the builtin error interface
func (e PointValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPoint.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PointValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PointValidationError{}

// Validate checks the field values on Rectangle with the rules defined in the
// proto definition for this message. If any rules are violated, an error is
// returned. When asked to return all errors, validation continues after first
// violation, and the result is a list of violation errors wrapped in
// RectangleMultiError, or nil if none found. Otherwise, only the first error
// is returned, if any.
func (m *Rectangle) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if v, ok := interface{}(m.GetLo()).(interface{ Validate(bool) error }); ok {
		if err := v.Validate(all); err != nil {
			err = RectangleValidationError{
				field:  "Lo",
				reason: "embedded message failed validation",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
	}

	if v, ok := interface{}(m.GetHi()).(interface{ Validate(bool) error }); ok {
		if err := v.Validate(all); err != nil {
			err = RectangleValidationError{
				field:  "Hi",
				reason: "embedded message failed validation",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return RectangleMultiError(errors)
	}
	return nil
}

// RectangleMultiError is an error wrapping multiple validation errors returned
// by Rectangle.Validate(true) if the designated constraints aren't met.
type RectangleMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RectangleMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RectangleMultiError) AllErrors() []error { return m }

// RectangleValidationError is the validation error returned by
// Rectangle.Validate if the designated constraints aren't met.
type RectangleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RectangleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RectangleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RectangleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RectangleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RectangleValidationError) ErrorName() string { return "RectangleValidationError" }

// Error satisfies the builtin error interface
func (e RectangleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRectangle.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RectangleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RectangleValidationError{}

// Validate checks the field values on Feature with the rules defined in the
// proto definition for this message. If any rules are violated, an error is
// returned. When asked to return all errors, validation continues after first
// violation, and the result is a list of violation errors wrapped in
// FeatureMultiError, or nil if none found. Otherwise, only the first error is
// returned, if any.
func (m *Feature) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if v, ok := interface{}(m.GetLocation()).(interface{ Validate(bool) error }); ok {
		if err := v.Validate(all); err != nil {
			err = FeatureValidationError{
				field:  "Location",
				reason: "embedded message failed validation",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
	}

	if len(errors) > 0 {
		return FeatureMultiError(errors)
	}
	return nil
}

// FeatureMultiError is an error wrapping multiple validation errors returned
// by Feature.Validate(true) if the designated constraints aren't met.
type FeatureMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FeatureMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FeatureMultiError) AllErrors() []error { return m }

// FeatureValidationError is the validation error returned by Feature.Validate
// if the designated constraints aren't met.
type FeatureValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FeatureValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FeatureValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FeatureValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FeatureValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FeatureValidationError) ErrorName() string { return "FeatureValidationError" }

// Error satisfies the builtin error interface
func (e FeatureValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeature.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FeatureValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FeatureValidationError{}

// Validate checks the field values on RouteNote with the rules defined in the
// proto definition for this message. If any rules are violated, an error is
// returned. When asked to return all errors, validation continues after first
// violation, and the result is a list of violation errors wrapped in
// RouteNoteMultiError, or nil if none found. Otherwise, only the first error
// is returned, if any.
func (m *RouteNote) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if v, ok := interface{}(m.GetLocation()).(interface{ Validate(bool) error }); ok {
		if err := v.Validate(all); err != nil {
			err = RouteNoteValidationError{
				field:  "Location",
				reason: "embedded message failed validation",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
	}

	// no validation rules for Message

	if len(errors) > 0 {
		return RouteNoteMultiError(errors)
	}
	return nil
}

// RouteNoteMultiError is an error wrapping multiple validation errors returned
// by RouteNote.Validate(true) if the designated constraints aren't met.
type RouteNoteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RouteNoteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RouteNoteMultiError) AllErrors() []error { return m }

// RouteNoteValidationError is the validation error returned by
// RouteNote.Validate if the designated constraints aren't met.
type RouteNoteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouteNoteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouteNoteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouteNoteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouteNoteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouteNoteValidationError) ErrorName() string { return "RouteNoteValidationError" }

// Error satisfies the builtin error interface
func (e RouteNoteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRouteNote.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouteNoteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouteNoteValidationError{}

// Validate checks the field values on RouteSummary with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned. When asked to return all errors, validation continues after
// first violation, and the result is a list of violation errors wrapped in
// RouteSummaryMultiError, or nil if none found. Otherwise, only the first
// error is returned, if any.
func (m *RouteSummary) Validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PointCount

	// no validation rules for FeatureCount

	// no validation rules for Distance

	// no validation rules for ElapsedTime

	if len(errors) > 0 {
		return RouteSummaryMultiError(errors)
	}
	return nil
}

// RouteSummaryMultiError is an error wrapping multiple validation errors
// returned by RouteSummary.Validate(true) if the designated constraints
// aren't met.
type RouteSummaryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RouteSummaryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RouteSummaryMultiError) AllErrors() []error { return m }

// RouteSummaryValidationError is the validation error returned by
// RouteSummary.Validate if the designated constraints aren't met.
type RouteSummaryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RouteSummaryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RouteSummaryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RouteSummaryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RouteSummaryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RouteSummaryValidationError) ErrorName() string { return "RouteSummaryValidationError" }

// Error satisfies the builtin error interface
func (e RouteSummaryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRouteSummary.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RouteSummaryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RouteSummaryValidationError{}