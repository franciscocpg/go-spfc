package service

import (
	"testing"
)

func TestPrintStatus(t *testing.T) {
	_, err := Status("bla")
	if err == nil {
		t.Error("Should give an error")
	}
	sr, err := Status("com.apple.ubd")
	if err != nil {
		t.Error(err.Error())
	}
	t.Logf("sr %t %b", sr.Running, sr.PID)
}
