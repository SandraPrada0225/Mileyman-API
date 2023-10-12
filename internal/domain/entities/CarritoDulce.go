package entities

type CarritoDulce struct {
	ID        uint64
	CarritoID uint64
	DulceID   uint64
	Unidades  int
	Subtotal  float64
}

func (CarritoDulce) TableName() string {
	return "carritos_dulces"
}
