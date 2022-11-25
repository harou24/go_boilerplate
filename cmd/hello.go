package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sayHelloCmd)
}

var sayHelloCmd = &cobra.Command{
	Use:   "hello",
	Short: "say hello",
	Long:  "hello is a very basic command that says hello",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello")
	},
}
