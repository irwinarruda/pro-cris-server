package proval

import (
	"errors"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^([A-Z0-9_'+\-\.]+)([A-Z0-9_+\-]+)@([A-Z0-9][A-Z0-9\-]*\.)+[A-Z]{2,}$`)

type StringValidator struct {
	validations map[string]string
	min         *int
	max         *int
}

func (s *StringValidator) Min(min int, message string) *StringValidator {
	s.min = &min
	s.validations["min"] = message
	return s
}

func (s *StringValidator) Max(max int, message string) *StringValidator {
	s.max = &max
	s.validations["max"] = message
	return s
}

func (s *StringValidator) Email(message string) *StringValidator {
	s.validations["email"] = message
	return s
}

func (s *StringValidator) Validate(val any) []error {
	errs := []error{}
	parsedVal, ok := val.(string)
	if !ok {
		errs = append(errs, errors.New(s.validations["default"]))
	}
	for key, message := range s.validations {
		if key == "min" {
			if len(parsedVal) < *s.min {
				errs = append(errs, errors.New(message))
			}
		} else if key == "max" {
			if len(parsedVal) > *s.min {
				errs = append(errs, errors.New(message))
			}

		} else if key == "email" {
			if !emailRegex.MatchString(parsedVal) {
				errs = append(errs, errors.New(message))
			}
		}
	}
	return errs
}

type Val struct{}

func (v *Val) String(message string) *StringValidator {
	validator := StringValidator{
		validations: map[string]string{},
	}
	validator.validations["default"] = message
	return &validator
}

func (v *Val) ToStringSlice(errs []error) []string {
	strs := []string{}
	for _, item := range errs {
		strs = append(strs, item.Error())
	}
	return strs
}

func New() Val {
	return Val{}
}
