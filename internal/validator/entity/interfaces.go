package entity

type IFormaPagoRepository interface {
	GetAll() ([]FormaPago, error)
	GetByClave(clave string) (FormaPago, error)
}

type IUsoCFDIRepository interface {
	GetAll() ([]UsoCFDI, error)
	GetByClave(clave string) (UsoCFDI, error)
}
