package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "go_boilerplate",
		Short: "A boilerplate to start an API server",
		Long:  "",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
