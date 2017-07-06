package logger

type NewLog struct {
	loggerInit logInit
	Set1 func (parameter...interface{})
	Set2 func (parameter...interface{})
	Set3 func (parameter...interface{})
}
type logInit struct {
	LogLevel

}
type LogLevel int

const (
	_      LogLevel = iota
	level1
	level2
	level3
)
