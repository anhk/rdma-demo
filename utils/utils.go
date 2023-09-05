package utils

import (
	"fmt"
	"os"
	"runtime/debug"
)

func Must(e any) {
	if e != nil {
		fmt.Printf("%s\n", e)
		fmt.Printf("%s\n", debug.Stack())
		os.Exit(-1)
	}
}
