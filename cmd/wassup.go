package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(wassupCmd)
	// wassupCmd.Flags().BoolP("advanced", "a", false, "Get prompts for advanced features. [default:false]")
	// wassupCmd.Flags().StringP("name", "n", "", "Name of project to create. [default:\"\"]")
	// wassupCmd.MarkFlagRequired("name")
	// wassupCmd.Flags().VarP(&flagFramework, "framework", "f", fmt.Sprintf("Framework to use. Allowed values: %s", strings.Join(flags.AllowedProjectTypes, ", ")))
	// wassupCmd.Flags().VarP(&flagDBDriver, "driver", "d", fmt.Sprintf("Database drivers to use. Allowed values: %s", strings.Join(flags.AllowedDBDrivers, ", ")))
}

var wassupCmd = &cobra.Command{
	Use:   "wassup",
	Short: "custom wassup command",
	Long:  `Wassup cmd to test arguments and flags`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("wassup")
	},
}
