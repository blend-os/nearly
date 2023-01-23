/*
Copyright © 2023 Rudra Saraswat <rs2009@ubuntu.com>
Licensed under GPL-3.0
*/
package cmd

import (
	"fmt"

	"github.com/blend-os/nearly/core"
	"github.com/spf13/cobra"
)

var version = "1.0"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nearly " + core.NearlyVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
