package log

import "fmt"
import "os"
import "sync"
import "time"
import "runtime"

const (
	debugLevel int = iota
	infoLevel
	warnLevel
	errorLevel
	criticalLevel
)

// Logger impl.
type loggerImpl struct {
	module string
	level  int
	sync.Mutex
}

var defaultLogger = makeDefaultLogger()

// Creates default logger.
func makeDefaultLogger() Logger {
	var module = os.Args[0]
	var level = debugLevel

	if val := os.Getenv("LOG_LEVEL"); val != "" {
		var levels = map[string]int{
			"debug":    debugLevel,
			"info":     infoLevel,
			"warning":  warnLevel,
			"warn":     warnLevel,
			"error":    errorLevel,
			"critical": criticalLevel,
			"fatal":    criticalLevel,
		}
		level = levels[val]
	}

	return &loggerImpl{module: module, level: level}
}

// Debug log.
func (l *loggerImpl) Debug(msg string, args ...interface{}) {
	l.print(debugLevel, "DEBU", msg, args...)
}

// Info log.
func (l *loggerImpl) Info(msg string, args ...interface{}) {
	l.print(infoLevel, "INFO", msg, args...)
}

// Warning log.
func (l *loggerImpl) Warning(msg string, args ...interface{}) {
	l.print(warnLevel, "WARN", msg, args...)
}

// Error log.
func (l *loggerImpl) Error(msg string, args ...interface{}) {
	l.print(errorLevel, "ERRO", msg, args...)
}

// Fatal log.
func (l *loggerImpl) Fatal(msg string, args ...interface{}) {
	l.print(criticalLevel, "FATA", msg, args...)
}

func color(v int) string {
	return fmt.Sprintf("\033[%dm", 30+v)
}

// console colors
var (
	black   = color(0)
	red     = color(1)
	green   = color(2)
	yellow  = color(3)
	blue    = color(4)
	magenta = color(5)
	cyan    = color(6)
	white   = color(7)
	reset   = "\033[0m"
	colors  = []string{
		debugLevel:    cyan,
		infoLevel:     green,
		warnLevel:     yellow,
		errorLevel:    red,
		criticalLevel: magenta,
	}
	isWindows = runtime.GOOS == "windows"
)

// Print a message.
func (l *loggerImpl) print(level int, lstr string, msg string, args ...interface{}) {
	l.Lock()
	defer l.Unlock()

	if l.level > level {
		return
	}

	var t = time.Now().Format("15:04:05.000000")
	var p string
	if isWindows {
		p = fmt.Sprintf("[%s] %s %s ", l.module, t, lstr)
	} else {
		p = fmt.Sprintf("[%s] %s%s %s%s%s ", l.module, magenta, t, colors[level], lstr, reset)
	}
	var s = fmt.Sprintf(msg, args...)
	fmt.Println(p, s)
}
