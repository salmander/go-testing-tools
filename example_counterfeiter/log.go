package example_counterfeiter

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o mocks/log.go --fake-name Log . CustomLogger
type CustomLogger interface {
	Log(message string, args ...interface{})
}
