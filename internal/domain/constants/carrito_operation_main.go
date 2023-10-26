package constants

type CarritoOperationResult string

const (
	Created CarritoOperationResult = "Created"
	Deleted CarritoOperationResult = "Deleted"
	Updated CarritoOperationResult = "Updated"
	Error   CarritoOperationResult = "Error"
)

func (operation CarritoOperationResult) String() string {
	return string(operation)
}
