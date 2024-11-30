/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/mactunechy/go-todo-cli/core"
	"github.com/spf13/cobra"
)

var (
	id          int
	status      string
	validDtatus = map[string]bool{"done": true, "progress": true, "todo": true}
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the status of a todo item",
	Long: `Update the status of an existing TODO item by providing its ID and the desired status.
					The valid statuses are:
						- done
						- progress
						- todo`,
	Run: func(cmd *cobra.Command, args []string) {
		if id <= 0 {
			log.Fatalln("Invalid ID. Please provide a positive integer for the -id flag.")
		}

		if !validDtatus[status] {
			log.Fatalf("Invalid status: %s. Valid statuses are: done, progress, todo. \n", status)
		}

		err := core.Update(id, status)
		if err != nil {
			log.Fatalf("Error updating TODO", err)
		}

		fmt.Printf("TODO item with ID %d updated to status %s successfully\n", id, status)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().IntVar(&id, "id", 0, "ID of the TODO item to update (required)")
	updateCmd.Flags().StringVar(&status, "status", "", "New status of the TODO item (done, progress, todo)")

	updateCmd.MarkFlagRequired("id")
	updateCmd.MarkFlagRequired("status")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
