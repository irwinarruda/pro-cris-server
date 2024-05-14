package proval

type Val struct{}

func (v *Val) String(message string) *StringValidator {
	validator := StringValidator{
		validations: map[string]string{},
	}
	validator.validations["default"] = message
	return &validator
}

func (v *Val) Object(message string) {
	// TODO:
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
