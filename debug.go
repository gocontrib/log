package log

import (
	"os"
	"strings"
	"sync"
)

var (
	debugModules = &envlist{name: "DEBUG"}
)

// DebugLogger with extra utilities.
type DebugLogger interface {
	Logger
	Enabled() bool
	Err(op string, err error) error
}

// IfDebug returns non empty logger if debug configuration property contains given module.
func IfDebug(module string) DebugLogger {
	module = strings.TrimSpace(module)

	if len(module) == 0 {
		return &emptyLogger{}
	}

	return &debugLogger{
		module: strings.ToLower(module),
		prefix: module + ">>",
	}
}

type debugLogger struct {
	module  string
	prefix  string
	enabled int
}

func (l *debugLogger) Disabled() bool {
	if l.enabled == 0 {
		l.enabled = -1
		if debugModules.has(l.module) {
			l.enabled = 1
		}
	}
	return l.enabled == -1
}

func (l *debugLogger) Enabled() bool {
	return !l.Disabled()
}

func (l *debugLogger) Debug(msg string, args ...interface{}) {
	if l.Disabled() {
		return
	}
	Debug(l.prefix+msg, args...)
}

func (l *debugLogger) Info(msg string, args ...interface{}) {
	if l.Disabled() {
		return
	}
	Info(l.prefix+msg, args...)
}

func (l *debugLogger) Warning(msg string, args ...interface{}) {
	if l.Disabled() {
		return
	}
	Warning(l.prefix+msg, args...)
}

func (l *debugLogger) Error(msg string, args ...interface{}) {
	if l.Disabled() {
		return
	}
	Error(l.prefix+msg, args...)
}

func (l *debugLogger) Fatal(msg string, args ...interface{}) {
	Fatal(l.prefix+msg, args...)
}

func (l *debugLogger) Err(op string, err error) error {
	if !l.Disabled() && err != nil {
		l.Error("%s failed: %v", op, err)
	}
	return err
}

type emptyLogger struct{}

func (l *emptyLogger) Debug(msg string, args ...interface{})   {}
func (l *emptyLogger) Info(msg string, args ...interface{})    {}
func (l *emptyLogger) Warning(msg string, args ...interface{}) {}
func (l *emptyLogger) Error(msg string, args ...interface{})   {}
func (l *emptyLogger) Fatal(msg string, args ...interface{})   {}

func (l *emptyLogger) Enabled() bool { return false }

func (l *emptyLogger) Err(op string, err error) error {
	return err
}

type envlist struct {
	sync.Mutex
	name string
	data map[string]struct{}
}

func (t *envlist) has(key string) bool {
	t.Lock()
	defer t.Unlock()

	if t.data == nil {
		t.data = make(map[string]struct{})
		for _, name := range split(os.Getenv(t.name)) {
			t.data[name] = struct{}{}
		}
	}

	_, ok := t.data[key]
	return ok
}

func split(s string) []string {
	var a = strings.Split(s, ",")
	var result []string
	for _, t := range a {
		var k = strings.ToLower(strings.TrimSpace(t))
		if len(k) == 0 {
			continue
		}
		result = append(result, k)
	}
	return result
}
