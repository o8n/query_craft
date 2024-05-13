package service

import (
	"bufio"
	"fmt"
	"github.com/atotto/clipboard"
	"os"
	"strings"
)

// ReadIDs reads IDs from a given CSV file path.
func ReadIDs(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ids []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		// ignore empty lines
		if line != "" {
			ids = append(ids, line)
		}
	}
	return ids, scanner.Err()
}

func GenerateSQL(ids []string) (string, error) {
	// Join IDs with comma
	joinedIDs := strings.Join(ids, ", ")

	// Generate SQL based on user input
	var operation string
	var table string
	fmt.Println("Do you want to generate a SQL statement? [yes/no]")
	var generateSQL string
	fmt.Scanln(&generateSQL)

	if generateSQL != "yes" {
		return "", nil
	}

	fmt.Println("Please choose: select, update, or delete")
	fmt.Scanln(&operation)

	fmt.Println("Enter the table name:")
	fmt.Scanln(&table)

	switch operation {
	case "select":
		return fmt.Sprintf("SELECT * FROM %s WHERE user_id IN (%s);", table, joinedIDs), nil
	case "update":
		fmt.Println("Enter the set clause (e.g., `set column = value`):")
		var setClause string
		fmt.Scanln(&setClause)
		return fmt.Sprintf("UPDATE %s %s WHERE user_id IN (%s);", table, setClause, joinedIDs), nil
	case "delete":
		return fmt.Sprintf("DELETE FROM %s WHERE user_id IN (%s);", table, joinedIDs), nil
	default:
		return "", fmt.Errorf("invalid operation: %s", operation)
	}
}

func CopyToClipboard(content string) error {
	return clipboard.WriteAll(content)
}
