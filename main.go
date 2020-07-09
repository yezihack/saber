package main

import (
	"log"

	"github.com/spf13/cobra/doc"
	"github.com/yezihack/saber/cmd"
)

func main() {
	cmd.Execute()
}
func generateDoc() {
	err := doc.GenMarkdownTree(cmd.GetRoot(), "./")
	if err != nil {
		log.Fatal(err)
	}
}
