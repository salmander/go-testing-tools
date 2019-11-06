package go_testing_tools

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -o counterfeiter_mocks/log.go --fake-name Log . CustomLogger
//go:generate mockgen -package=gomock_mocks -destination=./gomock_mocks/log.go -source=log.go
type CustomLogger interface {
	Log(message string, args ...interface{})
}
