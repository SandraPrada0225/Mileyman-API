package entities

type Categoria struct {
	ID     uint64 `json:"id"`
	Nombre string `json:"nombre"`
}

func (Categoria) TableName() string {
	return "Categorias"
}
