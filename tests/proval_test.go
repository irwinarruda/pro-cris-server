package tests

import (
	"fmt"
	"testing"

	"github.com/irwinarruda/pro-cris-server/libs/proval"
)

func HasMessage(errs []error, message string) bool {
	for _, err := range errs {
		if err.Error() == message {
			return true
		}
	}
	return false
}

func TestValString(t *testing.T) {
	v := proval.New()
	message := "any message"
	errs := v.String(message).Validate(4)
	ok := HasMessage(errs, message)
	if !ok {
		t.Fatal("It should return error if primitive value is passed.")
	}
	errs = v.String(message).Validate(struct{}{})
	ok = HasMessage(errs, message)
	if !ok {
		t.Fatal("It should return error if struct value is passed.")
	}
}

func TestValStringMin(t *testing.T) {
	v := proval.New()
	message := "any message"
	errs := v.String("").Min(4, message).Validate("any")
	ok := HasMessage(errs, message)
	if !ok {
		t.Fatal("It should return an error if message length is less than the param.")
	}
	errs = v.String("").Min(3, message).Validate("any")
	ok = HasMessage(errs, message)
	if ok {
		t.Fatal("It should not return an error if message length is greater than the param.")
	}
	errs = v.String("").Min(-3, message).Validate("")
	ok = HasMessage(errs, message)
	if ok {
		t.Fatal("It should not return an error if min value is negative.")
	}
}

func TestValStringMax(t *testing.T) {
	v := proval.New()
	message := "any message"
	errs := v.String("").Max(4, message).Validate("any")
	ok := HasMessage(errs, message)
	if ok {
		t.Fatal("It should not return an error if message length is less than the param.")
	}
	errs = v.String("").Max(2, message).Validate("any")
	ok = HasMessage(errs, message)
	if !ok {
		t.Fatal("It should return an error if message length is greater than the param.")
	}
	errs = v.String("").Max(-3, message).Validate("")
	ok = HasMessage(errs, message)
	if !ok {
		t.Fatal("It should return an error if min value is negative.")
	}
}

func TestValStringEmail(t *testing.T) {
	v := proval.New()
	message := "any message"
	validator := v.String("").Email(message)
	errs := validator.Validate("any")
	ok := HasMessage(errs, message)
	if !ok {
		t.Fatal("It should return an error if message is not a valid email.")
	}
	errs = []error{}
	errs = append(errs, validator.Validate("any@any.any")...)
	errs = append(errs, validator.Validate("any.any@any.any")...)
	fmt.Println(errs)
	ok = HasMessage(errs, message)
	if ok {
		t.Fatal("It should not return an error if message is a valid email.")
	}
}
