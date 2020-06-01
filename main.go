package main

import (
	"fmt"
	cmd "not_your_fathers_search_engine_crawler/cmd"
)

func init() {
	cmd.InitializeApp()
}

func main() {
	fmt.Println("Starting App...")
	cmd.StartApp()
}
