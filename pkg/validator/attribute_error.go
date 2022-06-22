package validator

import "fmt"

type ErrAttrValidation struct {
	Name       string
	Value      string
	Code       string
	ErrMessage string
	IsValid    bool
}

func (e ErrAttrValidation) GetCode() string {
	return e.Code
}

func (e ErrAttrValidation) GetMessage() string {
	return e.ErrMessage
}

func (e ErrAttrValidation) Error() string {
	return fmt.Sprintf("%s - %s", e.GetCode(), e.GetMessage())
}
