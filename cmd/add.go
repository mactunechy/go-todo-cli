/*
Copyright © 2024 Dellan Muchengapadare <mactunechy@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/mactunechy/go-todo-cli/core"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds an item to the todo list",
	Long:  `Adds an item to the todo list by providing a description as an argument.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		todo := args[0]

		err := core.Save(todo)
		if err != nil {
			log.Fatalln("Error saving todo", err) // do not shoe the error to the user
		}
		fmt.Println("Todo added successfully")
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
