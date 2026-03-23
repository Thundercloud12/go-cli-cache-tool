/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var updateTags string

func updateInJson(index int, newValue string, newTags string) error {
	data, err := loadJsonFile()
	if err != nil {
		return err
	}

	if index < 0 || index >= len(data) {
		return fmt.Errorf("invalid command index: %d", index)
	}

	// Update the command
	data[index].Value = newValue

	// Update tags if provided
	if newTags != "" {
		data[index].Tags = strings.Split(newTags, ",")
	}

	// Marshal and write back to file
	data_write, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	err = os.WriteFile("./commands.json", data_write, 0644)
	if err != nil {
		return fmt.Errorf("error writing JSON file: %w", err)
	}

	return nil
}

var updateCmd = &cobra.Command{
	Use:   "update [index] [command]",
	Short: "Update a command",
	Long:  `Update a command by its index. Use 'list' to see indexed commands. Optionally update tags with -t or --tags flag.`,
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: invalid index. Please provide a number.")
			return
		}

		// Convert to 0-based index
		index--

		commandValue := strings.Join(args[1:], " ")

		err = updateInJson(index, commandValue, updateTags)
		if err != nil {
			fmt.Printf("Error updating command: %v\n", err)
			return
		}

		fmt.Println("Command updated successfully!")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&updateTags, "tags", "t", "", "Update tags (comma-separated)")
}
