/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

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

var listTags string

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "list",
	Short: "List all commands or filter by tags",
	Long:  `List all commands or filter by specific tags. Use -t or --tags to filter by comma-separated tags.`,
	Run: func(cmd *cobra.Command, args []string) {
		commands, err := loadJsonFile()
		if err != nil {
			fmt.Println("Error loading commands:", err)
			return
		}

		if len(commands) == 0 {
			fmt.Println("No commands found. Add one with: my-tool add <command> -t tag1,tag2")
			return
		}

		// Filter by tags if provided
		var filtered []Command
		if listTags != "" {
			requiredTags := strings.Split(listTags, ",")
			for _, c := range commands {
				if containsAnyTag(c.Tags, requiredTags) {
					filtered = append(filtered, c)
				}
			}
		} else {
			filtered = commands
		}

		if len(filtered) == 0 {
			fmt.Printf("No commands found with tags: %s\n", listTags)
			return
		}

		fmt.Println("\nCommands:")
		for i, c := range filtered {
			fmt.Printf("%d. %s\n", i+1, c.Value)
			if len(c.Tags) > 0 {
				fmt.Printf("   Tags: %s\n", strings.Join(c.Tags, ", "))
			}
		}
	},
}

func containsAnyTag(commandTags, requiredTags []string) bool {
	for _, req := range requiredTags {
		for _, tag := range commandTags {
			if tag == req {
				return true
			}
		}
	}
	return false
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().StringVarP(&listTags, "tags", "t", "", "Filter by tags (comma-separated)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
