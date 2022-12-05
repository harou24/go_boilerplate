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
	Use:   "start",
	Short: "start the API server",
	Long:  "Start the API serve that will listen on port 5000.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Hello")

		if flag != "" {
			fmt.Print(" World")
		}
		fmt.Println()
	},
}
