package go_testing_tools

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o example_counterfeiter/mocks/log.go --fake-name Log . CustomLogger
//go:generate mockgen -destination=./example_gomock/mocks/log.go -source=log.go
type CustomLogger interface {
	Log(message string, args ...interface{})
}
