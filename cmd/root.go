package cmd

import (
	"fmt"
	"os"

	"github.com/o8n/QueryCraft/pkg/service"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "QueryCraft",
	Short: "QueryCraft is a CLI tool for generating SQL from CSV files",
	Long:  `QueryCraft is a CLI tool that reads member IDs from a CSV file and generates SQL statements based on user inputs.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: CSV file path is required")
			os.Exit(1)
		}
		filePath := args[0]
		processFile(filePath)
	},
}

func processFile(filepath string)  {
	ids, err := service.ReadIDs(filepath)	
	if err != nil {
		fmt.Printf("Failed to read IDs from %s: %s\n", filepath, err)
		os.Exit(1)
	}
	sql, err := service.GenerateSQL(ids)
	if err != nil {
		fmt.Printf("Failed to generate SQL: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Generated SQL:", sql)
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringP("file", "f", "", "CSV file containing member IDs")
	rootCmd.PersistentFlags().BoolP("sql", "s", false, "Flag to generate SQL statement")
}
