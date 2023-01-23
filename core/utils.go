package core

import (
	"fmt"
	"os"
)

func CheckRoot() {
	if os.Geteuid() != 0 {
		fmt.Println("must be run as root, exiting.")
		os.Exit(1)
	}
}
