package entities

type CarritoDulce struct {
	ID        uint64  `json:"id" gorm:"primaryKey"`
	CarritoID uint64  `json:"carrito_id"`
	DulceID   uint64  `json:"dulce_id"`
	Unidades  int     `json:"unidades"`
	Subtotal  float64 `json:"subtotal"`
}

func NewCarritoDulce(carritoID uint64, dulceID uint64) CarritoDulce{
	return  CarritoDulce{
		CarritoID: carritoID,
		DulceID: dulceID,
	} 
}

func UpdateCarritoDulce(carritoDulce CarritoDulce, unidades int, precio float64) CarritoDulce{
	carritoDulce.Unidades = unidades
	carritoDulce.Subtotal = float64(unidades) * precio
	return carritoDulce
}

func (CarritoDulce) TableName() string {
	return "carritos_dulces"
}
