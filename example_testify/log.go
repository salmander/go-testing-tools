package example_testify

type CustomLogger interface {
	Log(message string, args ...interface{})
}
