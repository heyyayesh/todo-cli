/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// toggleCmd represents the toggle command
var toggleCmd = &cobra.Command{
	Use:   "toggle",
	Short: "Toggles the completion status of a todo item",
	Long: `Toggles the completion status of a todo item.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a todo id.\nUsage: todo-cli toggle <id>")
			return
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Please provide a valid todo id.\nUsage: todo-cli toggle <id>")
			return
		}

		err = todos.Toggle(id - 1) // todo ids are 1-indexed
		if err != nil {
			fmt.Println(err)
			return
		}

		err = SaveData(todos)
		if err != nil {
			fmt.Println(err)
			return
		}
		todos.Print()
	},
}

func init() {
	rootCmd.AddCommand(toggleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// toggleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// toggleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
