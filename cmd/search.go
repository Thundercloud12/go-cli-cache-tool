package cmd

import (
	"fmt"
	"sort"
	"strings"

	helper "github.com/Thundercloud12/go-cli-cache-tool.git/helper"
	"github.com/spf13/cobra"
)

type SearchResult struct {
	Command Command
	Score   int
}

func scoreFuzzy(query, candidate string) int {
	return helper.FuzzySearcher(query, candidate)
}

func minWordDistance(query, commandValue string) int {
	words := strings.Fields(strings.ToLower(commandValue))
	query = strings.ToLower(strings.TrimSpace(query))

	if query == "" {
		return 0
	}
	if len(words) == 0 {
		return scoreFuzzy(query, "")
	}

	best := int(^uint(0) >> 1)
	for _, word := range words {
		score := scoreFuzzy(query, word)
		if score < best {
			best = score
		}
	}

	return best
}

var searchCmd = &cobra.Command{
	Use:   "search [query]",
	Short: "Search commands using fuzzy matching",
	Long:  "Search saved commands using fuzzy matching and rank results by nearest match.",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		query := args[0]
		commands, err := loadJsonFile()
		if err != nil {
			fmt.Println("Error loading commands:", err)
			return
		}

		if len(commands) == 0 {
			fmt.Println("No commands found. Add one first with: add \"<command>\" -t <tags>")
			return
		}

		results := make([]SearchResult, 0, len(commands))
		for _, commandItem := range commands {
			score := minWordDistance(query, commandItem.Value)
			results = append(results, SearchResult{Command: commandItem, Score: score})
		}

		sort.Slice(results, func(i, j int) bool {
			if results[i].Score == results[j].Score {
				return results[i].Command.Value < results[j].Command.Value
			}
			return results[i].Score < results[j].Score
		})

		maxResults := 5
		if len(results) < maxResults {
			maxResults = len(results)
		}

		fmt.Printf("Top matches for %q:\n", query)
		for i := 0; i < maxResults; i++ {
			fmt.Printf("%d. %s\n", i+1, results[i].Command.Value)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
