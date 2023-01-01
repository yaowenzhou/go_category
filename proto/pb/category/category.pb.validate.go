// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: category.proto

package pb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
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
	_ = sort.Sort
)

// Validate checks the field values on CreateCategoryReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateCategoryReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCategoryReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCategoryReqMultiError, or nil if none found.
func (m *CreateCategoryReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCategoryReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := CreateCategoryReqValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Level

	// no validation rules for Parent

	// no validation rules for Images

	// no validation rules for Description

	if len(errors) > 0 {
		return CreateCategoryReqMultiError(errors)
	}

	return nil
}

// CreateCategoryReqMultiError is an error wrapping multiple validation errors
// returned by CreateCategoryReq.ValidateAll() if the designated constraints
// aren't met.
type CreateCategoryReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCategoryReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCategoryReqMultiError) AllErrors() []error { return m }

// CreateCategoryReqValidationError is the validation error returned by
// CreateCategoryReq.Validate if the designated constraints aren't met.
type CreateCategoryReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCategoryReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCategoryReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCategoryReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCategoryReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCategoryReqValidationError) ErrorName() string {
	return "CreateCategoryReqValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCategoryReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCategoryReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCategoryReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCategoryReqValidationError{}

// Validate checks the field values on CreateCategoryRsp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateCategoryRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCategoryRsp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCategoryRspMultiError, or nil if none found.
func (m *CreateCategoryRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCategoryRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return CreateCategoryRspMultiError(errors)
	}

	return nil
}

// CreateCategoryRspMultiError is an error wrapping multiple validation errors
// returned by CreateCategoryRsp.ValidateAll() if the designated constraints
// aren't met.
type CreateCategoryRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCategoryRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCategoryRspMultiError) AllErrors() []error { return m }

// CreateCategoryRspValidationError is the validation error returned by
// CreateCategoryRsp.Validate if the designated constraints aren't met.
type CreateCategoryRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCategoryRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCategoryRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCategoryRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCategoryRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCategoryRspValidationError) ErrorName() string {
	return "CreateCategoryRspValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCategoryRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCategoryRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCategoryRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCategoryRspValidationError{}

// Validate checks the field values on UpdateCategoryReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateCategoryReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateCategoryReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateCategoryReqMultiError, or nil if none found.
func (m *UpdateCategoryReq) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateCategoryReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Level

	// no validation rules for Parent

	// no validation rules for Images

	// no validation rules for Description

	if len(errors) > 0 {
		return UpdateCategoryReqMultiError(errors)
	}

	return nil
}

// UpdateCategoryReqMultiError is an error wrapping multiple validation errors
// returned by UpdateCategoryReq.ValidateAll() if the designated constraints
// aren't met.
type UpdateCategoryReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateCategoryReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateCategoryReqMultiError) AllErrors() []error { return m }

// UpdateCategoryReqValidationError is the validation error returned by
// UpdateCategoryReq.Validate if the designated constraints aren't met.
type UpdateCategoryReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateCategoryReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateCategoryReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateCategoryReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateCategoryReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateCategoryReqValidationError) ErrorName() string {
	return "UpdateCategoryReqValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateCategoryReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateCategoryReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateCategoryReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateCategoryReqValidationError{}

// Validate checks the field values on UpdateCategoryRsp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UpdateCategoryRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateCategoryRsp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateCategoryRspMultiError, or nil if none found.
func (m *UpdateCategoryRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateCategoryRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return UpdateCategoryRspMultiError(errors)
	}

	return nil
}

// UpdateCategoryRspMultiError is an error wrapping multiple validation errors
// returned by UpdateCategoryRsp.ValidateAll() if the designated constraints
// aren't met.
type UpdateCategoryRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateCategoryRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateCategoryRspMultiError) AllErrors() []error { return m }

// UpdateCategoryRspValidationError is the validation error returned by
// UpdateCategoryRsp.Validate if the designated constraints aren't met.
type UpdateCategoryRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateCategoryRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateCategoryRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateCategoryRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateCategoryRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateCategoryRspValidationError) ErrorName() string {
	return "UpdateCategoryRspValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateCategoryRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateCategoryRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateCategoryRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateCategoryRspValidationError{}

// Validate checks the field values on DeleteCategoryReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteCategoryReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteCategoryReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteCategoryReqMultiError, or nil if none found.
func (m *DeleteCategoryReq) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteCategoryReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return DeleteCategoryReqMultiError(errors)
	}

	return nil
}

// DeleteCategoryReqMultiError is an error wrapping multiple validation errors
// returned by DeleteCategoryReq.ValidateAll() if the designated constraints
// aren't met.
type DeleteCategoryReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteCategoryReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteCategoryReqMultiError) AllErrors() []error { return m }

// DeleteCategoryReqValidationError is the validation error returned by
// DeleteCategoryReq.Validate if the designated constraints aren't met.
type DeleteCategoryReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteCategoryReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteCategoryReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteCategoryReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteCategoryReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteCategoryReqValidationError) ErrorName() string {
	return "DeleteCategoryReqValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteCategoryReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteCategoryReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteCategoryReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteCategoryReqValidationError{}

// Validate checks the field values on DeleteCategoryRsp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteCategoryRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteCategoryRsp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteCategoryRspMultiError, or nil if none found.
func (m *DeleteCategoryRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteCategoryRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return DeleteCategoryRspMultiError(errors)
	}

	return nil
}

// DeleteCategoryRspMultiError is an error wrapping multiple validation errors
// returned by DeleteCategoryRsp.ValidateAll() if the designated constraints
// aren't met.
type DeleteCategoryRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteCategoryRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteCategoryRspMultiError) AllErrors() []error { return m }

// DeleteCategoryRspValidationError is the validation error returned by
// DeleteCategoryRsp.Validate if the designated constraints aren't met.
type DeleteCategoryRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteCategoryRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteCategoryRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteCategoryRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteCategoryRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteCategoryRspValidationError) ErrorName() string {
	return "DeleteCategoryRspValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteCategoryRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteCategoryRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteCategoryRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteCategoryRspValidationError{}

// Validate checks the field values on CategoryInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CategoryInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CategoryInfo with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CategoryInfoMultiError, or
// nil if none found.
func (m *CategoryInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *CategoryInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Level

	// no validation rules for Parent

	// no validation rules for Images

	// no validation rules for Description

	if len(errors) > 0 {
		return CategoryInfoMultiError(errors)
	}

	return nil
}

// CategoryInfoMultiError is an error wrapping multiple validation errors
// returned by CategoryInfo.ValidateAll() if the designated constraints aren't met.
type CategoryInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CategoryInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CategoryInfoMultiError) AllErrors() []error { return m }

// CategoryInfoValidationError is the validation error returned by
// CategoryInfo.Validate if the designated constraints aren't met.
type CategoryInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CategoryInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CategoryInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CategoryInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CategoryInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CategoryInfoValidationError) ErrorName() string { return "CategoryInfoValidationError" }

// Error satisfies the builtin error interface
func (e CategoryInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCategoryInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CategoryInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CategoryInfoValidationError{}

// Validate checks the field values on FindAllCategoryReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FindAllCategoryReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FindAllCategoryReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FindAllCategoryReqMultiError, or nil if none found.
func (m *FindAllCategoryReq) ValidateAll() error {
	return m.validate(true)
}

func (m *FindAllCategoryReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return FindAllCategoryReqMultiError(errors)
	}

	return nil
}

// FindAllCategoryReqMultiError is an error wrapping multiple validation errors
// returned by FindAllCategoryReq.ValidateAll() if the designated constraints
// aren't met.
type FindAllCategoryReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FindAllCategoryReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FindAllCategoryReqMultiError) AllErrors() []error { return m }

// FindAllCategoryReqValidationError is the validation error returned by
// FindAllCategoryReq.Validate if the designated constraints aren't met.
type FindAllCategoryReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FindAllCategoryReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FindAllCategoryReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FindAllCategoryReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FindAllCategoryReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FindAllCategoryReqValidationError) ErrorName() string {
	return "FindAllCategoryReqValidationError"
}

// Error satisfies the builtin error interface
func (e FindAllCategoryReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFindAllCategoryReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FindAllCategoryReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FindAllCategoryReqValidationError{}

// Validate checks the field values on FindAllCategoryRsp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FindAllCategoryRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FindAllCategoryRsp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FindAllCategoryRspMultiError, or nil if none found.
func (m *FindAllCategoryRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *FindAllCategoryRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetDatas() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, FindAllCategoryRspValidationError{
						field:  fmt.Sprintf("Datas[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, FindAllCategoryRspValidationError{
						field:  fmt.Sprintf("Datas[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return FindAllCategoryRspValidationError{
					field:  fmt.Sprintf("Datas[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Total

	if len(errors) > 0 {
		return FindAllCategoryRspMultiError(errors)
	}

	return nil
}

// FindAllCategoryRspMultiError is an error wrapping multiple validation errors
// returned by FindAllCategoryRsp.ValidateAll() if the designated constraints
// aren't met.
type FindAllCategoryRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FindAllCategoryRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FindAllCategoryRspMultiError) AllErrors() []error { return m }

// FindAllCategoryRspValidationError is the validation error returned by
// FindAllCategoryRsp.Validate if the designated constraints aren't met.
type FindAllCategoryRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FindAllCategoryRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FindAllCategoryRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FindAllCategoryRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FindAllCategoryRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FindAllCategoryRspValidationError) ErrorName() string {
	return "FindAllCategoryRspValidationError"
}

// Error satisfies the builtin error interface
func (e FindAllCategoryRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFindAllCategoryRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FindAllCategoryRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FindAllCategoryRspValidationError{}

// Validate checks the field values on FindCategoryByNameReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FindCategoryByNameReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FindCategoryByNameReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FindCategoryByNameReqMultiError, or nil if none found.
func (m *FindCategoryByNameReq) ValidateAll() error {
	return m.validate(true)
}

func (m *FindCategoryByNameReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return FindCategoryByNameReqMultiError(errors)
	}

	return nil
}

// FindCategoryByNameReqMultiError is an error wrapping multiple validation
// errors returned by FindCategoryByNameReq.ValidateAll() if the designated
// constraints aren't met.
type FindCategoryByNameReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FindCategoryByNameReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FindCategoryByNameReqMultiError) AllErrors() []error { return m }

// FindCategoryByNameReqValidationError is the validation error returned by
// FindCategoryByNameReq.Validate if the designated constraints aren't met.
type FindCategoryByNameReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FindCategoryByNameReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FindCategoryByNameReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FindCategoryByNameReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FindCategoryByNameReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FindCategoryByNameReqValidationError) ErrorName() string {
	return "FindCategoryByNameReqValidationError"
}

// Error satisfies the builtin error interface
func (e FindCategoryByNameReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFindCategoryByNameReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FindCategoryByNameReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FindCategoryByNameReqValidationError{}

// Validate checks the field values on FindCategoryByNameRsp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *FindCategoryByNameRsp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FindCategoryByNameRsp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// FindCategoryByNameRspMultiError, or nil if none found.
func (m *FindCategoryByNameRsp) ValidateAll() error {
	return m.validate(true)
}

func (m *FindCategoryByNameRsp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetDatas()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, FindCategoryByNameRspValidationError{
					field:  "Datas",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, FindCategoryByNameRspValidationError{
					field:  "Datas",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDatas()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FindCategoryByNameRspValidationError{
				field:  "Datas",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return FindCategoryByNameRspMultiError(errors)
	}

	return nil
}

// FindCategoryByNameRspMultiError is an error wrapping multiple validation
// errors returned by FindCategoryByNameRsp.ValidateAll() if the designated
// constraints aren't met.
type FindCategoryByNameRspMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FindCategoryByNameRspMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FindCategoryByNameRspMultiError) AllErrors() []error { return m }

// FindCategoryByNameRspValidationError is the validation error returned by
// FindCategoryByNameRsp.Validate if the designated constraints aren't met.
type FindCategoryByNameRspValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FindCategoryByNameRspValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FindCategoryByNameRspValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FindCategoryByNameRspValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FindCategoryByNameRspValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FindCategoryByNameRspValidationError) ErrorName() string {
	return "FindCategoryByNameRspValidationError"
}

// Error satisfies the builtin error interface
func (e FindCategoryByNameRspValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFindCategoryByNameRsp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FindCategoryByNameRspValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FindCategoryByNameRspValidationError{}
