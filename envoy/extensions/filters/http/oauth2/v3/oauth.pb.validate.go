//go:build !disable_pgv
// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/http/oauth2/v3/oauth.proto

package oauth2v3

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

// Validate checks the field values on OAuth2Credentials with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *OAuth2Credentials) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OAuth2Credentials with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OAuth2CredentialsMultiError, or nil if none found.
func (m *OAuth2Credentials) ValidateAll() error {
	return m.validate(true)
}

func (m *OAuth2Credentials) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetClientId()) < 1 {
		err := OAuth2CredentialsValidationError{
			field:  "ClientId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetTokenSecret() == nil {
		err := OAuth2CredentialsValidationError{
			field:  "TokenSecret",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetTokenSecret()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OAuth2CredentialsValidationError{
					field:  "TokenSecret",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OAuth2CredentialsValidationError{
					field:  "TokenSecret",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTokenSecret()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OAuth2CredentialsValidationError{
				field:  "TokenSecret",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCookieNames()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OAuth2CredentialsValidationError{
					field:  "CookieNames",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OAuth2CredentialsValidationError{
					field:  "CookieNames",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCookieNames()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OAuth2CredentialsValidationError{
				field:  "CookieNames",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	oneofTokenFormationPresent := false
	switch v := m.TokenFormation.(type) {
	case *OAuth2Credentials_HmacSecret:
		if v == nil {
			err := OAuth2CredentialsValidationError{
				field:  "TokenFormation",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofTokenFormationPresent = true

		if m.GetHmacSecret() == nil {
			err := OAuth2CredentialsValidationError{
				field:  "HmacSecret",
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetHmacSecret()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, OAuth2CredentialsValidationError{
						field:  "HmacSecret",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OAuth2CredentialsValidationError{
						field:  "HmacSecret",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetHmacSecret()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OAuth2CredentialsValidationError{
					field:  "HmacSecret",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofTokenFormationPresent {
		err := OAuth2CredentialsValidationError{
			field:  "TokenFormation",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return OAuth2CredentialsMultiError(errors)
	}

	return nil
}

// OAuth2CredentialsMultiError is an error wrapping multiple validation errors
// returned by OAuth2Credentials.ValidateAll() if the designated constraints
// aren't met.
type OAuth2CredentialsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OAuth2CredentialsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OAuth2CredentialsMultiError) AllErrors() []error { return m }

// OAuth2CredentialsValidationError is the validation error returned by
// OAuth2Credentials.Validate if the designated constraints aren't met.
type OAuth2CredentialsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OAuth2CredentialsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OAuth2CredentialsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OAuth2CredentialsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OAuth2CredentialsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OAuth2CredentialsValidationError) ErrorName() string {
	return "OAuth2CredentialsValidationError"
}

// Error satisfies the builtin error interface
func (e OAuth2CredentialsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOAuth2Credentials.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OAuth2CredentialsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OAuth2CredentialsValidationError{}

// Validate checks the field values on OAuth2Config with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OAuth2Config) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OAuth2Config with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OAuth2ConfigMultiError, or
// nil if none found.
func (m *OAuth2Config) ValidateAll() error {
	return m.validate(true)
}

func (m *OAuth2Config) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTokenEndpoint()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OAuth2ConfigValidationError{
					field:  "TokenEndpoint",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OAuth2ConfigValidationError{
					field:  "TokenEndpoint",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTokenEndpoint()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OAuth2ConfigValidationError{
				field:  "TokenEndpoint",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetAuthorizationEndpoint()) < 1 {
		err := OAuth2ConfigValidationError{
			field:  "AuthorizationEndpoint",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetCredentials() == nil {
		err := OAuth2ConfigValidationError{
			field:  "Credentials",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetCredentials()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OAuth2ConfigValidationError{
					field:  "Credentials",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OAuth2ConfigValidationError{
					field:  "Credentials",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCredentials()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OAuth2ConfigValidationError{
				field:  "Credentials",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetRedirectUri()) < 1 {
		err := OAuth2ConfigValidationError{
			field:  "RedirectUri",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetRedirectPathMatcher() == nil {
		err := OAuth2ConfigValidationError{
			field:  "RedirectPathMatcher",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetRedirectPathMatcher()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OAuth2ConfigValidationError{
					field:  "RedirectPathMatcher",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OAuth2ConfigValidationError{
					field:  "RedirectPathMatcher",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRedirectPathMatcher()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OAuth2ConfigValidationError{
				field:  "RedirectPathMatcher",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetSignoutPath() == nil {
		err := OAuth2ConfigValidationError{
			field:  "SignoutPath",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetSignoutPath()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OAuth2ConfigValidationError{
					field:  "SignoutPath",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OAuth2ConfigValidationError{
					field:  "SignoutPath",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSignoutPath()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OAuth2ConfigValidationError{
				field:  "SignoutPath",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ForwardBearerToken

	for idx, item := range m.GetPassThroughMatcher() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, OAuth2ConfigValidationError{
						field:  fmt.Sprintf("PassThroughMatcher[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OAuth2ConfigValidationError{
						field:  fmt.Sprintf("PassThroughMatcher[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OAuth2ConfigValidationError{
					field:  fmt.Sprintf("PassThroughMatcher[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if _, ok := OAuth2Config_AuthType_name[int32(m.GetAuthType())]; !ok {
		err := OAuth2ConfigValidationError{
			field:  "AuthType",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetUseRefreshToken()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OAuth2ConfigValidationError{
					field:  "UseRefreshToken",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OAuth2ConfigValidationError{
					field:  "UseRefreshToken",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUseRefreshToken()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OAuth2ConfigValidationError{
				field:  "UseRefreshToken",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDefaultExpiresIn()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OAuth2ConfigValidationError{
					field:  "DefaultExpiresIn",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OAuth2ConfigValidationError{
					field:  "DefaultExpiresIn",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDefaultExpiresIn()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OAuth2ConfigValidationError{
				field:  "DefaultExpiresIn",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetDenyRedirectMatcher() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, OAuth2ConfigValidationError{
						field:  fmt.Sprintf("DenyRedirectMatcher[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OAuth2ConfigValidationError{
						field:  fmt.Sprintf("DenyRedirectMatcher[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OAuth2ConfigValidationError{
					field:  fmt.Sprintf("DenyRedirectMatcher[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return OAuth2ConfigMultiError(errors)
	}

	return nil
}

// OAuth2ConfigMultiError is an error wrapping multiple validation errors
// returned by OAuth2Config.ValidateAll() if the designated constraints aren't met.
type OAuth2ConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OAuth2ConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OAuth2ConfigMultiError) AllErrors() []error { return m }

// OAuth2ConfigValidationError is the validation error returned by
// OAuth2Config.Validate if the designated constraints aren't met.
type OAuth2ConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OAuth2ConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OAuth2ConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OAuth2ConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OAuth2ConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OAuth2ConfigValidationError) ErrorName() string { return "OAuth2ConfigValidationError" }

// Error satisfies the builtin error interface
func (e OAuth2ConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOAuth2Config.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OAuth2ConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OAuth2ConfigValidationError{}

// Validate checks the field values on OAuth2 with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OAuth2) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OAuth2 with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in OAuth2MultiError, or nil if none found.
func (m *OAuth2) ValidateAll() error {
	return m.validate(true)
}

func (m *OAuth2) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, OAuth2ValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, OAuth2ValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OAuth2ValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return OAuth2MultiError(errors)
	}

	return nil
}

// OAuth2MultiError is an error wrapping multiple validation errors returned by
// OAuth2.ValidateAll() if the designated constraints aren't met.
type OAuth2MultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OAuth2MultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OAuth2MultiError) AllErrors() []error { return m }

// OAuth2ValidationError is the validation error returned by OAuth2.Validate if
// the designated constraints aren't met.
type OAuth2ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OAuth2ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OAuth2ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OAuth2ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OAuth2ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OAuth2ValidationError) ErrorName() string { return "OAuth2ValidationError" }

// Error satisfies the builtin error interface
func (e OAuth2ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOAuth2.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OAuth2ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OAuth2ValidationError{}

// Validate checks the field values on OAuth2Credentials_CookieNames with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OAuth2Credentials_CookieNames) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OAuth2Credentials_CookieNames with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// OAuth2Credentials_CookieNamesMultiError, or nil if none found.
func (m *OAuth2Credentials_CookieNames) ValidateAll() error {
	return m.validate(true)
}

func (m *OAuth2Credentials_CookieNames) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetBearerToken() != "" {

		if !_OAuth2Credentials_CookieNames_BearerToken_Pattern.MatchString(m.GetBearerToken()) {
			err := OAuth2Credentials_CookieNamesValidationError{
				field:  "BearerToken",
				reason: "value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetOauthHmac() != "" {

		if !_OAuth2Credentials_CookieNames_OauthHmac_Pattern.MatchString(m.GetOauthHmac()) {
			err := OAuth2Credentials_CookieNamesValidationError{
				field:  "OauthHmac",
				reason: "value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetOauthExpires() != "" {

		if !_OAuth2Credentials_CookieNames_OauthExpires_Pattern.MatchString(m.GetOauthExpires()) {
			err := OAuth2Credentials_CookieNamesValidationError{
				field:  "OauthExpires",
				reason: "value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetIdToken() != "" {

		if !_OAuth2Credentials_CookieNames_IdToken_Pattern.MatchString(m.GetIdToken()) {
			err := OAuth2Credentials_CookieNamesValidationError{
				field:  "IdToken",
				reason: "value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetRefreshToken() != "" {

		if !_OAuth2Credentials_CookieNames_RefreshToken_Pattern.MatchString(m.GetRefreshToken()) {
			err := OAuth2Credentials_CookieNamesValidationError{
				field:  "RefreshToken",
				reason: "value does not match regex pattern \"^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return OAuth2Credentials_CookieNamesMultiError(errors)
	}

	return nil
}

// OAuth2Credentials_CookieNamesMultiError is an error wrapping multiple
// validation errors returned by OAuth2Credentials_CookieNames.ValidateAll()
// if the designated constraints aren't met.
type OAuth2Credentials_CookieNamesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OAuth2Credentials_CookieNamesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OAuth2Credentials_CookieNamesMultiError) AllErrors() []error { return m }

// OAuth2Credentials_CookieNamesValidationError is the validation error
// returned by OAuth2Credentials_CookieNames.Validate if the designated
// constraints aren't met.
type OAuth2Credentials_CookieNamesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OAuth2Credentials_CookieNamesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OAuth2Credentials_CookieNamesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OAuth2Credentials_CookieNamesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OAuth2Credentials_CookieNamesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OAuth2Credentials_CookieNamesValidationError) ErrorName() string {
	return "OAuth2Credentials_CookieNamesValidationError"
}

// Error satisfies the builtin error interface
func (e OAuth2Credentials_CookieNamesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOAuth2Credentials_CookieNames.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OAuth2Credentials_CookieNamesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OAuth2Credentials_CookieNamesValidationError{}

var _OAuth2Credentials_CookieNames_BearerToken_Pattern = regexp.MustCompile("^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$")

var _OAuth2Credentials_CookieNames_OauthHmac_Pattern = regexp.MustCompile("^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$")

var _OAuth2Credentials_CookieNames_OauthExpires_Pattern = regexp.MustCompile("^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$")

var _OAuth2Credentials_CookieNames_IdToken_Pattern = regexp.MustCompile("^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$")

var _OAuth2Credentials_CookieNames_RefreshToken_Pattern = regexp.MustCompile("^:?[0-9a-zA-Z!#$%&'*+-.^_|~`]+$")
