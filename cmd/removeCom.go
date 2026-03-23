/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func deleteFromJson(index int) error {
	data, err := loadJsonFile()
	if err != nil {
		return err
	}

	if index < 0 || index >= len(data) {
		return fmt.Errorf("invalid command index: %d", index)
	}

	// Remove the command at the specified index
	data = append(data[:index], data[index+1:]...)

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

var deleteCmd = &cobra.Command{
	Use:   "delete [index]",
	Short: "Delete a command by index",
	Long:  `Delete a command by its index number. Use 'list' command to see indexed commands.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Error: invalid index. Please provide a number.")
			return
		}

		// Convert to 0-based index
		index--

		err = deleteFromJson(index)
		if err != nil {
			fmt.Printf("Error deleting command: %v\n", err)
			return
		}

		fmt.Println("Command deleted successfully!")
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// removeComCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// removeComCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
