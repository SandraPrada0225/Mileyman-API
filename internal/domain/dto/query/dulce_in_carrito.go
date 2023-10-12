package query

type DulceInCarrito struct {
	DetalleDulce
	Unidades int     `json:"unidades"`
	Subtotal float64 `json:"subtotal"`
}

func NewDulceInCarrito(dulce DetalleDulce, unidades int, subtotal float64) DulceInCarrito {
	return DulceInCarrito{
		DetalleDulce: dulce,
		Unidades:     unidades,
		Subtotal:     subtotal,
	}
}
