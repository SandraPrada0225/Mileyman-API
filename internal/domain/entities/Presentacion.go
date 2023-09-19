package entities

type Presentacion struct {
	ID     uint64 `json:"id"`
	Nombre string `json:"nombre"`
}

func (Presentacion) TableName() string {
	return "presentaciones"
}
