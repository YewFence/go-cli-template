package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var appVersion = "dev"

var rootCmd = &cobra.Command{
	Use:   "your-cli",
	Short: "Your CLI description",
	Long:  "Your CLI description.",
	RunE: func(cmd *cobra.Command, args []string) error {
		_, err := fmt.Fprintln(cmd.OutOrStdout(), "Hello from your-cli")
		return err
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func SetVersion(version string) {
	appVersion = version
}
