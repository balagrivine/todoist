/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/balagrivine/todoist/todo"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: addRun,
}

func init() {
	rootCmd.AddCommand(addCmd)

	var priority int

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "priority: 1, 2, 3")
}

func addRun(cmd *cobra.Command, args []string) {

	var (
		priority int
	)

	items, err := todo.ReadItems("./.todos.json"); if err != nil {
		log.Printf("Error: %v\n", err)
	}

	for _, val := range args {
		item := todo.Item{Text: val}
		item.SetPriority(priority)
		items = append(items, item)
	}
	if err := todo.SaveItems("./.todos.json", items); err != nil {
		fmt.Errorf("Error: %v\n", err)
	}
}
