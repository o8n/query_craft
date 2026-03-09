package service

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/atotto/clipboard"
	"io"
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
	return generateSQL(ids, os.Stdin, os.Stdout)
}

func CopyToClipboard(content string) error {
	return clipboard.WriteAll(content)
}

func generateSQL(ids []string, input io.Reader, output io.Writer) (string, error) {
	if len(ids) == 0 {
		return "", errors.New("no IDs found in the input file")
	}

	reader := bufio.NewReader(input)
	joinedIDs := strings.Join(ids, ", ")

	fmt.Fprintln(output, "Do you want to generate a SQL statement? [yes/no]")
	generateSQL, err := readLine(reader)
	if err != nil {
		return "", err
	}
	if !strings.EqualFold(generateSQL, "yes") {
		return "", nil
	}

	fmt.Fprintln(output, "Please choose: select, update, or delete")
	operation, err := readLine(reader)
	if err != nil {
		return "", err
	}

	fmt.Fprintln(output, "Enter the table name:")
	table, err := readLine(reader)
	if err != nil {
		return "", err
	}
	if table == "" {
		return "", errors.New("table name is required")
	}

	switch strings.ToLower(operation) {
	case "select":
		return fmt.Sprintf("SELECT * FROM %s WHERE user_id IN (%s);", table, joinedIDs), nil
	case "update":
		fmt.Fprintln(output, "Enter the set clause (e.g., `set column = value`):")
		setClause, err := readLine(reader)
		if err != nil {
			return "", err
		}
		if setClause == "" {
			return "", errors.New("set clause is required for update")
		}
		return fmt.Sprintf("UPDATE %s %s WHERE user_id IN (%s);", table, setClause, joinedIDs), nil
	case "delete":
		return fmt.Sprintf("DELETE FROM %s WHERE user_id IN (%s);", table, joinedIDs), nil
	default:
		return "", fmt.Errorf("invalid operation: %s", operation)
	}
}

func readLine(reader *bufio.Reader) (string, error) {
	line, err := reader.ReadString('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		return "", err
	}
	line = strings.TrimSpace(line)
	if err != nil && errors.Is(err, io.EOF) && line == "" {
		return "", io.EOF
	}
	return line, nil
}
