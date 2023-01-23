/*
Copyright © 2023 Rudra Saraswat <rs2009@ubuntu.com>
Licensed under GPL-3.0
*/
package core

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

var nearly_paths = []string{
	"/usr", // since blendOS uses the merged-usr layout
}

func SwitchRo(ro bool, verbose bool) {
	if ro {
		fmt.Println("Enabling immutability.")
	} else {
		fmt.Println("Disabling immutability.")
	}
	for _, path := range nearly_paths {
		nearly_files_l1, _ := filepath.Glob(filepath.Join(path, "**"))

		for _, path_l1 := range nearly_files_l1 {
			/* if ro && GetImmutable(path_l1) {
				return
			} else if ro == false && !GetImmutable(path_l1) {
				return
			} */

			nearly_files_l2, _ := filepath.Glob(filepath.Join(path_l1, "**"))

			var wg sync.WaitGroup

			for _, file := range nearly_files_l2 {
				wg.Add(1)
				go func(file string, ro bool, verbose bool) {
					defer wg.Done()
					immutable(file, ro, verbose)
				}(file, ro, verbose)
				immutable(file, ro, verbose)
			}

			wg.Add(1)
			go func(file string, ro bool, verbose bool) {
				defer wg.Done()
				immutable(file, ro, verbose)
			}(path_l1, ro, verbose)
			immutable(path_l1, ro, verbose)
		}
	}
	if ro {
		fmt.Println("Immutability has been enabled.")
	} else {
		fmt.Println("Immutability has been disabled.")
	}
}

func immutable(file string, ro bool, verbose bool) {
	if ro && GetImmutable(file) {
		return
	} else if ro == false && !GetImmutable(file) {
		return
	}

	if verbose {
		print("handing file: " + file + "\n")
	}

	f, err := os.OpenFile(file, os.O_RDONLY, 0755)

	leg := false

	if err != nil {
		if verbose {
			fmt.Printf("Error opening %s: %v, try setting using the legacy chattr tool\n", file, err)
			leg = true
		}

		f, err = os.Open(file)
		if err != nil {
			if verbose {
				fmt.Printf("Skipping %s: %s\n", file, err.Error())
				return
			}
		}
	}

	switch {
	case ro && leg:
		err := LegacySetAttr(file, "i")
		if err != nil && verbose {
			fmt.Printf("(legacy) Error setting immutable flag on %s: %s\n", file, err.Error())
		}
	case ro:
		err := SetAttr(f, FS_IMMUTABLE_FL)
		if err != nil && verbose {
			fmt.Println("Error while setting immutable flag: ", err)
		}
	case leg:
		err := LegacyUnsetAttr(file, "i")
		if err != nil && verbose {
			fmt.Printf("(legacy) Error removing immutable flag on %s: %s\n", file, err.Error())
		}
	default:
		err := UnsetAttr(f, FS_IMMUTABLE_FL)
		if err != nil && verbose {
			fmt.Println("Error while setting immutable flag: ", err)
		}
	}
}
