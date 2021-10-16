package zlog

import (
	"testing"
)

func TestZlog(t *testing.T) {
	l := NewLogger()
	l.Info("aaaa")
	l.Error("bbbb")
}
