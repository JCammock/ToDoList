/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"strconv"

	"github.com/spf13/cobra"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update <taskId> <task name>",
	Short: "Update the name of a specified task",
	Long: `Update a specific task with a new name, e.g.
	ToDoList update 1 "Wash the car"`,
	Run: func(cmd *cobra.Command, args []string) {
		taskID, _ := strconv.Atoi(args[0])
		updateTask(taskID-1, args[1])
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
