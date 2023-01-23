/*
Copyright © 2023 Rudra Saraswat <rs2009@ubuntu.com>
Licensed under GPL-3.0
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/blend-os/nearly/core"
	"github.com/spf13/cobra"
)

var enterUsage = `(switching to rw is not recommended, as you cannot revert
any changes made while the system is read-write)

Usage:
    enter [options] [command]

Options:
    --help/-h       show this message
    --verbose/-v    verbose output

Commands:
    ro              set the filesystem as read-only
    rw              set the filesystem as read-write
    default         set the filesystem as defined in the configuration file

Examples:
    nearly enter ro
    nearly enter rw
`

var enterCmd = &cobra.Command{
	Use:   "enter",
	Short: "Toggle immutability",
	RunE:  enter_mode,
}

func enter_mode(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		cmd.Help()
		os.Exit(0)
	}

	core.CheckRoot()

	v, _ := cmd.Flags().GetBool("verbose")

	switch args[0] {
	case "ro":
		core.SwitchRo(true, v)
	case "rw":
		core.SwitchRo(false, v)
	default:
		// return fmt.Errorf("unknown command: %s", args[0])
	}

	return nil
}

func init() {
	enterCmd.Flags().BoolP("verbose", "v", false, "enable verbose output")

	rootCmd.AddCommand(enterCmd)
	enterCmd.SetUsageFunc(func(*cobra.Command) error { fmt.Println(enterUsage); return nil })
}
