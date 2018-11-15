package main

import (
	"os"

	"github.com/theoden9014/exercitu/cmd"
)

func main() {
	cmd.Execute(os.Args[1:])
}
