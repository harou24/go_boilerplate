package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var flag string

func init() {
	rootCmd.AddCommand(sayHelloCmd)
	sayHelloCmd.PersistentFlags().StringVarP(&flag, "world", "w", "", "add world to hello")
}

var sayHelloCmd = &cobra.Command{
	Use:   "hello",
	Short: "say hello",
	Long:  "hello is a very basic command that says hello",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Hello")

		if flag != "" {
			fmt.Print(" World")
		}
		fmt.Println()
	},
}
