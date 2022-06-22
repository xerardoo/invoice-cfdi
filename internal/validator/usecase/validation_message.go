package usecase

type CFDIValidationMessage struct {
	Cfdi struct {
		Version    string `json:"version"`
		Validation struct {
			Code   string `json:"code"`
			Result struct {
				IsValid bool   `json:"isValid"`
				Message string `json:"message"`
			} `json:"result"`
		} `json:"validation"`
	} `json:"cfdi"`
}

func NewValidationMessage() *CFDIValidationMessage {
	return &CFDIValidationMessage{}
}

func (c *CFDIValidationMessage) SetVersion(version string) {
	c.Cfdi.Version = version
}

func (c *CFDIValidationMessage) SetCode(code string) {
	c.Cfdi.Validation.Code = code
}

func (c *CFDIValidationMessage) SetIsValid(isvalid bool) {
	c.Cfdi.Validation.Result.IsValid = isvalid
}

func (c *CFDIValidationMessage) SetMessage(message string) {
	c.Cfdi.Validation.Result.Message = message

}
