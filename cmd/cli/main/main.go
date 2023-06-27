package main

import (
	"fmt"
	"http-server/internal/app/cli"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	var cli = cli.BuildCli()
	fmt.Println(cli.Run(argsWithoutProg))
}
