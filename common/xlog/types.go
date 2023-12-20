package xlog

type Level uint

type Config struct {
	Cache      int
	PrintLevel Level // Only output logs greater than or equal to this level
	Report     ReportConfig
}

type ReportConfig struct {
	Enabled     bool
	ReportLevel Level // Only report logs greater than or equal to this level
	Send        SendConfig
	Receive     ReceiveConfig
}

type SendConfig struct {
	Host     string
	Port     int
	Sender   string
	Username string
	Password string
}

type ReceiveConfig struct {
	To          []string // emails
	Subject     string
	ContentType string // text or html
	ContentTmpl string /*
		%s must exist in the template, and it will be replaced by log information at %s.

		example:

		The following is log information:\n%s\nPlease Notice!

		The following is log information:

		service_name:user-rpc

		content:mysql server is not available

		level:fatal

		...

		Please Notice!*/
}

const (
	LevelDebug Level = 0
	LevelInfo  Level = 1
	LevelWarn  Level = 2
	LevelError Level = 3
	LevelFatal Level = 4
)
