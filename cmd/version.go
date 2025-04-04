package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of my test CLI app ",
	Long:  `All software has versions. This is supposed to be version manpage?`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version 0.01")
	},
}
