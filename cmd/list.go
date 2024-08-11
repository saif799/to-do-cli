package cmd

import (
	"github.com/spf13/cobra"
)

var listAll bool
var listcmnd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"read", "show"},
	Short:   "lists all the tasks",
	Long:    "prints all the undone tasks add the -a flag for all the tasks including the undone ones",
	Run: func(cmd *cobra.Command, args []string) {
		Read(listAll)
	},
}

func init() {
	listcmnd.Flags().BoolVarP(&listAll, "all", "a", false, "list all the tasks")

	rootCmd.AddCommand(listcmnd)
}
