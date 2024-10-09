/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list <status>",
	Short: "List all tasks or optionally a list of tasks matching a status",
	Long: `List all tasks or specify a status to list tasks matching the status.
	Parameters include "todo", "in-progress", and "done" e.g.
	ToDoList list
	ToDoList list in-progress`,
	Run: func(cmd *cobra.Command, args []string) {
		listTasks(args[0])
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
