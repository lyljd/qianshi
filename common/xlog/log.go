package xlog

func Debug(content string) {
	generateStandardLog(LevelDebug, content)
}

func Info(content string) {
	generateStandardLog(LevelInfo, content)
}

func Warn(content string) {
	generateStandardLog(LevelWarn, content)
}

func Error(content string) {
	generateStandardLog(LevelError, content)
}

func Fatal(content string) {
	generateStandardLog(LevelFatal, content)
}
