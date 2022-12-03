package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "hello",
		Short: "say hello",
		Long:  "hello is a very basic command that says hello",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
