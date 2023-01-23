/*
Copyright © 2023 Rudra Saraswat <rs2009@ubuntu.com>
Licensed under GPL-3.0
*/
package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/blend-os/nearly/core"
	"github.com/spf13/cobra"
)

var runUsage = `(switching to rw is not recommended, as you cannot revert
any changes made while the system is read-write)

Usage:
    run [options] [command]

Options:
    --help/-h       show this message
    --verbose/-v    verbose output

Examples:
    nearly run pacman -Sy
`

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a command in read-write mode",
	RunE:  run_cmd,
}

func run_cmd(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(0)
	}

	core.CheckRoot()

	v, _ := cmd.Flags().GetBool("verbose")
	core.SwitchRo(false, v)

	// XXX: Move this bit to core
	_cmd := exec.Command("sh", "-c", args[0])
	_cmd.Stdout = os.Stdout
	_cmd.Stderr = os.Stderr
	_cmd.Stdin = os.Stdin

	// we want to keep the current env
	_cmd.Env = os.Environ()

	if err := _cmd.Run(); err != nil {
		fmt.Println(err)
	}

	core.SwitchRo(true, v)

	return nil
}

func init() {
	runCmd.Flags().BoolP("verbose", "v", false, "enable verbose output")

	rootCmd.AddCommand(runCmd)
	runCmd.SetUsageFunc(func(*cobra.Command) error { fmt.Println(enterUsage); return nil })
}
