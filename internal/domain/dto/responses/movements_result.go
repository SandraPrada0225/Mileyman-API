package responses

type MovementsResult struct {
	Result []MovementResult
}

type MovementResult struct {
	Movement int    `json:"movement"`
	DulceID  uint64 `json:"dulce_id"`
	Result   string `json:"result"`
	Error    string `json:"error,omitempty"`
}

func NewMovementsResult() MovementsResult {
	return MovementsResult{}
}

func (m *MovementsResult) AddResult(movement int, dulceID uint64, result string, err string) {
	m.Result = append(m.Result, MovementResult{
		Movement: movement,
		DulceID:  dulceID,
		Result:   result,
		Error:    err,
	})
}
