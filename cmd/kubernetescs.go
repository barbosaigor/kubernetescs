package main

import (
	"fmt"

	"github.com/barbosaigor/kubernetescs/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
