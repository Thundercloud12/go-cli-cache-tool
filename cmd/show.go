/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"os"

	"github.com/spf13/cobra"
)

type Command struct {
	ID    string   `json:"id"`
	Value string   `json:"value"`
	Tags  []string `json:"tags"`
}

func loadJsonFile() ([]Command, error) {
	// This function will load the JSON file and parse the commands.
	data, err := os.ReadFile("./commands.json")
	if err != nil {
		// Return empty slice if file doesn't exist (first time use)
		if os.IsNotExist(err) {
			return []Command{}, nil
		}
		fmt.Println("Error reading JSON file:", err)
		return nil, err
	}

	// Handle empty file
	if len(data) == 0 {
		return []Command{}, nil
	}

	var commands []Command
	err = json.Unmarshal(data, &commands)
	if err != nil {
		return nil, err
	}
	return commands, nil
}

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "A brief description of your command",
	Long:  `A function to load the commands from the json file and display them in a user-friendly format. This function will read the JSON file, parse the commands, and print them to the console in a structured way.`,
	Run: func(cmd *cobra.Command, args []string) {
		commands, err := loadJsonFile()
		if err != nil {
			fmt.Println("Error hai", err)
			return
		}

		for i, cmd := range commands {
			fmt.Println(i+1, cmd.Value, cmd.Tags)
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
