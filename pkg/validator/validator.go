package validator

type IValidator interface {
	Validate() (result bool, err IErrValidation)
}

type IErrValidation interface {
	GetCode() string
	GetMessage() string
	Error() string
}