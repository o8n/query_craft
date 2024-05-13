package main

import (
	"fmt"
	"os"

	"github.com/o8n/QueryCraft/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
