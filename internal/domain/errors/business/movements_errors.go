package business

type MovementsResult struct {
	Result []MovementResult
}

type MovementResult struct {
	Movement int
	DulceID  uint64
	Result   string
	Error    error
}

func NewMovementsError() MovementResult {
	return MovementResult{}
}

func (m MovementsResult) AddResult(movement int, dulceID uint64, err string) {
	m.Result = append(m.Result, MovementResult{
		Movement: movement,
		DulceID:  dulceID,
		Result:   err,
	})
}
