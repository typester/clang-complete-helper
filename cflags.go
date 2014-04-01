package main

import (
	"path/filepath"
	"errors"
	"os"
	"fmt"
	"io"
)

func PrintCflags(file string) {
	_printCfags(file, os.Stdout)
}

func _printCfags(file string, w io.Writer) {
	p, err := FindProject(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
		return
	}

	cflags := p.Cflags()
	for _, flag := range(cflags) {
		fmt.Fprintf(w, `-I"%s" `, flag)
	}

	if p.Type() == "iphoneos" {
		sdkroot, err := find_ios_sdkroot()
		if err != nil {
			fmt.Fprintf(os.Stderr, "iPhone SDK not found\n")
		} else {
			fmt.Fprintf(w, "-isysroot=%s -miphoneos-version-min=4.0 -arch=i386", sdkroot)
		}
	}

	fmt.Fprintf(w, "\n")
}


func find_ios_sdkroot() (string, error) {
	dir := "/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneSimulator.platform/Developer/SDKs"

	sdks, err := filepath.Glob(filepath.Join(dir, "iPhoneSimulator*.sdk"))
	if err != nil {
		return "", err
	}

	if len(sdks) == 0 {
		return "", errors.New("sdkroot not found")
	}

	return sdks[0], nil
}
