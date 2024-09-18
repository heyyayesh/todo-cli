/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// clearCmd represents the clear command
var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		force, err := cmd.Flags().GetBool("force")
		if err != nil {
			fmt.Println(err)
			return
		}

		if force {
			err = ClearData()
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		if getConfirmation() {
			err = ClearData()
			if err != nil {
				fmt.Println(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(clearCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clearCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	clearCmd.Flags().BoolP("force", "f", false, "Force clear all todos")
}

func getConfirmation() bool {
	var c rune
	fmt.Print("This will delete all the todos. Are you sure? (y/n) ")
	fmt.Scanf("%c\n", &c)
	return c == 'y'
}
