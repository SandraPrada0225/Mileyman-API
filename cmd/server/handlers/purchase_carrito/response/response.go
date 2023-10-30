package response

type Response struct {
	NuevoCarritoID uint64 `json:"nuevo_carrito_id"`
}

func NewResponse(CarritoID uint64) Response {
	return Response{
		NuevoCarritoID: CarritoID,
	}
}
