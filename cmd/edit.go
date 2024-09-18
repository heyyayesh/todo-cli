/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Please provide a todo id and a new title.\nUsage: todo-cli edit <id> <new title>")
			return
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Please provide a valid todo id.\nUsage: todo-cli edit <id> <new title>")
			return
		}

		err = todos.Edit(id - 1, args[1]) // todo ids are 1-indexed
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
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
