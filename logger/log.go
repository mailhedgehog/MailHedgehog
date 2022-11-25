package logger

import "fmt"

// LogLevel represents an log level name
type LogLevel = string

var (
	LogHandler          func(message string, context ...interface{})
	LogEmergencyHandler func(message string, context ...interface{})
	LogAlertHandler     func(message string, context ...interface{})
	LogCriticalHandler  func(message string, context ...interface{})
	LogErrorHandler     func(message string, context ...interface{})
	LogWarningHandler   func(message string, context ...interface{})
	LogNoticeHandler    func(message string, context ...interface{})
	LogInfoHandler      func(message string, context ...interface{})
	LogDebugHandler     func(message string, context ...interface{})
)

const (
	EMERGENCY LogLevel = "emergency"
	ALERT              = "alert"
	CRITICAL           = "critical"
	ERROR              = "error"
	WARNING            = "warning"
	NOTICE             = "notice"
	INFO               = "info"
	DEBUG              = "debug"
)

var (
	EmergencyColor = Red
	AlertColor     = Red
	CriticalColor  = Red
	ErrorColor     = Red
	WarningColor   = Yellow
	InfoColor      = Teal
	NoticeColor    = Purple
	DebugColor     = Green
)

var (
	Red    = Color("\033[1;31m%s\033[0m")
	Green  = Color("\033[1;32m%s\033[0m")
	Yellow = Color("\033[1;33m%s\033[0m")
	Purple = Color("\033[1;34m%s\033[0m")
	Teal   = Color("\033[1;36m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func Log(level LogLevel, message string, context ...interface{}) {
	switch level {
	case EMERGENCY:
		Emergency(message, context...)
	case ALERT:
		Alert(message, context...)
	case CRITICAL:
		Critical(message, context...)
	case ERROR:
		Error(message, context...)
	case WARNING:
		Warning(message, context...)
	case NOTICE:
		Notice(message, context...)
	case INFO:
		Info(message, context...)
	case DEBUG:
		Debug(message, context...)
	default:
		panic(fmt.Sprintf("Wrong log level: %s", level))
	}
}

func Emergency(message string, context ...interface{}) {
	if LogEmergencyHandler != nil {
		LogEmergencyHandler(message, context...)
		return
	}

	if LogHandler != nil {
		LogHandler(message, context...)
		return
	}

	if len(context) > 0 {
		fmt.Println(EmergencyColor(message), context)
		return
	}
	fmt.Println(EmergencyColor(message))
}

func Alert(message string, context ...interface{}) {
	if LogAlertHandler != nil {
		LogAlertHandler(message, context...)
		return
	}

	if LogHandler != nil {
		LogHandler(message, context...)
		return
	}

	if len(context) > 0 {
		fmt.Println(AlertColor(message), context)
		return
	}
	fmt.Println(AlertColor(message))
}

func Critical(message string, context ...interface{}) {
	if LogCriticalHandler != nil {
		LogCriticalHandler(message, context...)
		return
	}

	if LogHandler != nil {
		LogHandler(message, context...)
		return
	}

	if len(context) > 0 {
		fmt.Println(CriticalColor(message), context)
		return
	}
	fmt.Println(CriticalColor(message))
}

func Error(message string, context ...interface{}) {
	if LogErrorHandler != nil {
		LogErrorHandler(message, context...)
		return
	}

	if LogHandler != nil {
		LogHandler(message, context...)
		return
	}

	if len(context) > 0 {
		fmt.Println(ErrorColor(message), context)
		return
	}
	fmt.Println(ErrorColor(message))
}

func Warning(message string, context ...interface{}) {
	if LogWarningHandler != nil {
		LogWarningHandler(message, context)
		return
	}

	if LogHandler != nil {
		LogHandler(message, context...)
		return
	}

	if len(context) > 0 {
		fmt.Println(WarningColor(message), context)
		return
	}
	fmt.Println(WarningColor(message))
}

func Notice(message string, context ...interface{}) {
	if LogNoticeHandler != nil {
		LogNoticeHandler(message, context...)
		return
	}

	if LogHandler != nil {
		LogHandler(message, context...)
		return
	}

	if len(context) > 0 {
		fmt.Println(NoticeColor(message), context)
		return
	}
	fmt.Println(NoticeColor(message))
}

func Info(message string, context ...interface{}) {
	if LogInfoHandler != nil {
		LogInfoHandler(message, context...)
		return
	}

	if LogHandler != nil {
		LogHandler(message, context...)
		return
	}

	if len(context) > 0 {
		fmt.Println(InfoColor(message), context)
		return
	}
	fmt.Println(InfoColor(message))
}

func Debug(message string, context ...interface{}) {
	if LogDebugHandler != nil {
		LogDebugHandler(message, context...)
		return
	}

	if LogHandler != nil {
		LogHandler(message, context...)
		return
	}

	if len(context) > 0 {
		fmt.Println(DebugColor(message), context)
		return
	}
	fmt.Println(DebugColor(message))
}
