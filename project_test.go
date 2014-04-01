package main

import (
	"path/filepath"
	"testing"
)

func Test_find_xcodeproj(t *testing.T) {
	f, e := find_xcodeproj("./testfiles/osxproject")

	if e != nil {
		t.Errorf("err shoud nil: %s", e.Error())
	}

	expected, e := filepath.Abs(filepath.Join("./testfiles/osxproject", "osxproject.xcodeproj"))
	if e != nil {
		t.Errorf("err shoud nil: %s", e.Error())
	}
	if f != expected {
		t.Errorf("file is wrong: %s != %s", f, expected)
	}
}

func Test_find_xcodeproj_fail(t *testing.T) {
	f, e := find_xcodeproj("./testfiles")
	if f != "" {
		t.Errorf("file is wrong: %s", f)
	}
	if e.Error() != "xcodeproj not found" {
		t.Error("error is wrong: %s", e.Error())
	}
}

func Test_FindProject(t *testing.T) {
	p, err := FindProject("./testfiles/osxproject/osxproject/AppDelegate.m")
	if err != nil {
		t.Errorf("error shoul be nil: %s", err.Error())
	}

	expected, err := filepath.Abs("./testfiles/osxproject/osxproject.xcodeproj")
	if err != nil {
		t.Errorf("err should be nil: %s", err.Error())
	}
	if p.file != expected {
		t.Error("unexpected project %s != %s", p, expected)
	}
}

func Test_Project_Type(t *testing.T) {
	p, err := FindProject("./testfiles/osxproject/osxproject/AppDelegate.m")
	if err != nil {
		t.Errorf("error shoul be nil: %s", err.Error())
	}

	if p.Type() != "macosx" {
		t.Errorf("unexpected type: %s", p.Type())
	}
}

func Test_Project_Cflags(t *testing.T) {
	p, _ := FindProject("./testfiles/iosproject/iosproject/AppDelegate.m")

	flags := p.Cflags()

	if len(flags) != 2 {
		t.Errorf("len(clags) should be 2 but %d", len(flags))
	}

	if flags[0] != "/Applications/Xcode.app/Contents/Developer/Toolchains/XcodeDefault.xctoolchain/usr/include" {
		t.Errorf("flags[0] is unexpected: %s", flags[0])
	}

	expected := filepath.Join(p.root, "foo/include")
	if flags[1] != expected {
		t.Errorf("flags[1] is unexpected: %s != %s", flags[1], expected)
	}
}

func Test_Project_Cflags_notfound(t *testing.T) {
	p, _ := FindProject("./testfiles/osxproject/osxproject/AppDelegate.m")

	flags := p.Cflags()

	if len(flags) != 0 {
		t.Error("flags should be empty")
	}
}
