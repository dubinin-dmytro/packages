package ABTower

type ResponseInterface interface {
	isSuccess() bool
	errorMessage() string
}
