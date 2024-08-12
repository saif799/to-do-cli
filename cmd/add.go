package cmd

import (
	"github.com/spf13/cobra"
)

var Addcmnd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"create", "write"},
	Short:   "adds a task",
	Long:    "takes a describtion of a tasks and adds it to the data.csv file ",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		Write(args[0])
	},
}

func init() {

	rootCmd.AddCommand(Addcmnd)
}
