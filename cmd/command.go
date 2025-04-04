package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vader",
	Short: "short: vader is a cli program tester app",
	// INFO: MANPAGE TOP
	Long: `LONG docs : vader is a cli program tester app. 
	Testing out the 'cobra' along with 'go-bootstrap'

	TODO:
	- arguments
	- flags
	- check is color possible in LONG field

	Formatting test
	# HEADING
	## HEADING2
	### HEADING3

	| test quote


	`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
