package updatecarrito

type Body struct {
	Movements []Movement `json:"movements"`
}

type Movement struct {
	DulceID  uint64 `json:"dulce_id"`
	Unidades int `json:"unidades"`
}
