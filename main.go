package main

import (
	"log"
	"os"

	"github.com/spf13/cobra/doc"
	"github.com/yezihack/saber/cmd"
)

func main() {
	// 带参数 go run . doc
	args := os.Args
	if len(args) > 0 && args[1] == "doc" {
		generateDoc()
		return
	}
	cmd.Execute()
}

// 生成文档
func generateDoc() {
	err := doc.GenMarkdownTree(cmd.GetRoot(), "./")
	if err != nil {
		log.Fatal(err)
	}
}
