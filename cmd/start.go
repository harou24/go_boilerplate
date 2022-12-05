package cmd

import (
	"github.com/harou24/go_boilerplate/api"
	"github.com/spf13/cobra"
)

var port string

func init() {
	rootCmd.AddCommand(start)
	start.PersistentFlags().StringVarP(&port, "port", "p", "", "run on a different port, default is 5000")
}

var start = &cobra.Command{
	Use:   "start",
	Short: "start the API server",
	Long:  "Start the API server that will listen on port 5000.",
	Run: func(cmd *cobra.Command, args []string) {
		api.NewManager().Serve()
	},
}
