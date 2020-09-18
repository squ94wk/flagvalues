package values

import (
	"fmt"
)

type InvalidValue struct {
	Type   string
	Value  string
	Reason string
	Help   string
	Cause  error
}

type ErrOption func(*InvalidValue)

func NewInvalidValue(opts ...ErrOption) *InvalidValue {
	err := new(InvalidValue)
	for _, opt := range opts {
		opt(err)
	}
	return err
}

func (i InvalidValue) Unwrap() error {
	return i.Cause
}

func (i InvalidValue) Error() string {
	msg := "invalid value"
	if i.Value != "" {
		msg += fmt.Sprintf(" '%s'", i.Value)
	}
	if i.Type != "" {
		msg += fmt.Sprintf(" for type '%s'", i.Type)
	}
	if i.Reason != "" {
		msg += ": " + i.Reason
	}
	if i.Cause != nil {
		msg += ": " + i.Cause.Error()
	}
	if i.Help != "" {
		msg += ": " + i.Help
	}
	return msg
}

func WithType(typ string) ErrOption {
	return func(i *InvalidValue) {
		i.Type = typ
	}
}

func WithValue(val string) ErrOption {
	return func(i *InvalidValue) {
		i.Value = val
	}
}

func WithCause(err error) ErrOption {
	return func(i *InvalidValue) {
		i.Cause = err
	}
}

func WithReason(msg string) ErrOption {
	return func(i *InvalidValue) {
		i.Reason = msg
	}
}

func WithHelp(msg string) ErrOption {
	return func(i *InvalidValue) {
		i.Help = msg
	}
}
