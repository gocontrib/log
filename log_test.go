package log

import . "testing"

func TestLog(t *T) {
	Debug("debug")
	Info("info")
	Warning("warning")
	Error("error")

	Debug("debug %s", "log")
	Info("info %s", "log")
	Warning("warning %s", "log")
	Error("error %s", "log")
	//Fatal("fatal %s", "log")
}
