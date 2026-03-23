package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/Thundercloud12/go-cli-cache-tool.git/helper"
	"github.com/spf13/cobra"
)

func addtoJson(command Command) error {

	data, err := loadJsonFile()

	if err != nil {
		fmt.Println("Error reading JSON file:", err)
		return err
	}

	data = append(data, command)

	data_write, err := json.Marshal(data)

	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return err
	}

	err = os.WriteFile("./commands.json", data_write, 0644)
	if err != nil {
		fmt.Println("Error writing JSON file:", err)
		return err
	}

	return nil

}

var tags string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long:  `add karna hai`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Please provide a command to add")
			return
		}

		commandValue := args[0]

		// parse tags
		var tagList []string
		if tags != "" {
			tagList = strings.Split(tags, ",")
		}

		newCommand := Command{
			ID:    helper.GenerateID(),
			Value: commandValue,
			Tags:  tagList,
		}

		err := addtoJson(newCommand)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Command added successfully!")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&tags, "tags", "t", "", "Tags for the command")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
