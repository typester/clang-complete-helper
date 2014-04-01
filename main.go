package main

import (
	"fmt"
	"os"
)

func print_usage(program string) {
	fmt.Printf("Usage: %s cflags|prefix-header filename\n", program)
}

func main() {
	if len(os.Args) != 3 {
		print_usage(os.Args[0])
		os.Exit(1)
	}

	switch os.Args[1] {
	case "cflags":
		PrintCflags(os.Args[2])
	case "prefix-header":
		// TODO
	default:
		print_usage(os.Args[0])
	}
}
