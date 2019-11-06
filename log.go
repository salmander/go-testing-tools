package go_testing_tools

//go:generate mockgen -destination=./mocks/log.go -source=log.go
type CustomLogger interface {
	Log(message string, args ...interface{})
}
