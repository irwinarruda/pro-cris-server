package proval

import (
	"errors"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+\/=?^_{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)

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
		}
		if key == "max" {
			if len(parsedVal) > *s.max {
				errs = append(errs, errors.New(message))
			}
		}
		if key == "email" {
			if !emailRegex.MatchString(parsedVal) {
				errs = append(errs, errors.New(message))
			}
		}
	}
	return errs
}
