package log

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"k8s.io/klog/v2"
)

// type Logger struct {
// 	level  LogLevel
// 	logger *log.Logger
// }

//	func NewLogger(level LogLevel) *Logger {
//		return &Logger{
//			level:  level,
//			logger: log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds|log.Lshortfile),
//		}
//	}
func getProjectRoot() string {
	_, b, _, _ := runtime.Caller(0)
	// Trim the file path to the root of the project directory
	return filepath.Join(filepath.Dir(b), "../..")
}
func Print(level LogLevel, format string, args ...interface{}) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	// Retrieve the file and line number of the caller
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		file = "unknown"
		line = 0
	}
	// Convert the file path to a relative path
	relFile, err := filepath.Rel(getProjectRoot(), file)
	if err != nil {
		relFile = file // fallback to the absolute path if there's an error
	}
	// Format the log message with the current time, file, and line number
	logMsg := fmt.Sprintf("[%s] %s:%d %s", currentTime, relFile, line, fmt.Sprintf(format, args...))

	// logMsg := fmt.Sprintf(format, args...)
	switch level {
	case Warn:
		klog.Warning(logMsg)
	case Error:
		klog.Error(logMsg)
	case Info:
		fallthrough
	default:
		klog.Info(logMsg)
	}
}
