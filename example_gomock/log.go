package example_gomock

//go:generate mockgen -destination=./mocks/log.go -source=log.go
type CustomLogger interface {
	Log(message string, args ...interface{})
}
