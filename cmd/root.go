package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	// "github.com/o8n/QueryCraft/pkg/pkg/service"
)

var rootCmd = &cobra.Command{
	Use:   "QueryCraft",
	Short: "QueryCraft is a CLI tool for generating SQL from CSV files",
	Long:  `QueryCraft is a CLI tool that reads member IDs from a CSV file and generates SQL statements based on user inputs.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("QueryCraft is running")
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringP("file", "f", "", "CSV file containing member IDs")
	rootCmd.PersistentFlags().BoolP("sql", "s", false, "Flag to generate SQL statement")
}
