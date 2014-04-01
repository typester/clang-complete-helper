package main

import (
	"testing"
	"bytes"
)

func TestPrintCflags(t *testing.T) {
	w := new(bytes.Buffer)

	_printCfags("testfiles/iosproject/iosproject/AppDelegate.m", w)
	t.Logf("flags: %s", w.String())

	w.Reset()
	_printCfags("testfiles/osxproject/osxproject/AppDelegate.m", w)
	t.Logf("flags: %s", w.String())
}
