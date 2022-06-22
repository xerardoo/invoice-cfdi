package validator

type IAttributeValidator interface {
	GetName() string
	GetValue() string
	GetCode() string
	GetErrMessage() string
	GetError() IErrValidation
	GetRule() func(attributeValue string) bool
}

type AttributeValidation struct {
	Name       string
	Value      string
	Code       string
	ErrMessage string
	Rule       func(attributeValue string) bool
}

func NewAttributeValidation(name, value, code, errMessage string, rule func(string) bool) AttributeValidation {
	return AttributeValidation{
		Name:       name,
		Value:      value,
		Rule:       rule,
		Code:       code,
		ErrMessage: errMessage,
	}
}

func (validation AttributeValidation) GetName() string {
	return validation.Name
}

func (validation AttributeValidation) GetValue() string {
	return validation.Value
}

func (validation AttributeValidation) GetRule() func(attributeValue string) bool {
	return validation.Rule
}

func (validation AttributeValidation) GetCode() string {
	return validation.Code
}

func (validation AttributeValidation) GetErrMessage() string {
	return validation.ErrMessage
}

func (validation AttributeValidation) GetError() IErrValidation {
	return ErrAttrValidation{
		Name:       validation.GetName(),
		Value:      validation.GetValue(),
		Code:       validation.GetCode(),
		ErrMessage: validation.GetErrMessage(),
	}
}

func (validation AttributeValidation) Validate() (result bool, err IErrValidation) {
	rule := validation.GetRule()
	value := validation.GetValue()

	if result = rule(value); !result {
		err = validation.GetError()
		return
	}
	return
}
