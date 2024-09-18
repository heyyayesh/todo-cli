/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a todo item to the list",
	Long: `Adds a todo item to the list. The title should be in double quotes if it contains multiple words.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a todo title.\nUsage: todo-cli add <todo title>")
			return
		}

		if len(args) > 1 {
			fmt.Println("Too many arguements. Provide the todo title in double quotes if it has multiple words.")
			return
		}
		
		err := todos.Add(args[0])
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
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
