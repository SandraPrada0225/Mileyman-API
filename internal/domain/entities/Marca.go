package entities

type Marca struct {
	ID     uint64 `json:"id"`
	Nombre string `json:"nombre"`
}

func (Marca) TableName() string {
	return "Marcas"
}
