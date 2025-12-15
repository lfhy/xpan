package log

type Log interface {
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

type DefaultLogger struct{}

func (l *DefaultLogger) Printf(format string, v ...interface{}) {
}

func (l *DefaultLogger) Println(v ...interface{}) {
}

var Logger Log

func SetLogger(l Log) {
	Logger = l
}

func GetLogger() Log {
	if Logger == nil {
		Logger = &DefaultLogger{}
	}
	return Logger
}

func Printf(format string, v ...interface{}) {
	GetLogger().Printf(format, v...)
}

func Println(v ...interface{}) {
	GetLogger().Println(v...)
}
