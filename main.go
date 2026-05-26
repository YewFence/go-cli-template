package main

import "github.com/example/your-cli/cmd"

var version = "dev"

func main() {
	cmd.Execute(version)
}
