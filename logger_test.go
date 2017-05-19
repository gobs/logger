package logger

import "testing"

func TestLogger(*testing.T) {
	l := GetLogger(INFO, "LOGGER")

	l.Log(INFO, "info")
	l.Log(WARNING, "warning")
	l.Log(ERROR, "error")
	l.Log(DEBUG, "debug")
}
