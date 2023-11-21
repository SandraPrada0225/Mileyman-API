package getpurchaselistbyuserid

import (
	"Mileyman-API/internal/domain/dto/responses"
	"Mileyman-API/internal/repositories/providers"
)

type Implementation struct {
	VentasProvider providers.VentasProvider
}

func (useCase Implementation) Execute(userID uint64) (responses.GetPurchaseList, error) {
	return useCase.VentasProvider.GetListByUserID(userID)
}
