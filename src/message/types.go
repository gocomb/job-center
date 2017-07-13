package message

func init() {

}

type Interface interface {
	Info(...interface{}) func(parameter ...interface{})
}

type JobMessage struct {
	FatalMessage map[string]error
	ErrMessage   map[string]string
	WarnMessage  map[string]string
	InfoMessage  map[string]string
	DebugMassage map[string]string
}

//消息文本，输出消息信息，返回文本数组
