package cli

import (
	"http-server/internal/app/cli/handlers"
	"strings"
)

type Cli struct {
	handlers            []CommandHandler
	commandToHandlerMap map[string]CommandHandler
}

func BuildCli() Cli {
	commands := []CommandHandler{handlers.CurlHandler{}, handlers.HelpHandler{}}
	commandToHandlerMap := map[string]CommandHandler{}

	for _, handler := range commands {
		commandToHandlerMap[handler.Identifier()] = handler
	}

	return Cli{
		handlers:            commands,
		commandToHandlerMap: commandToHandlerMap,
	}
}

func (c Cli) Help() string {
	commandDescriptors := []string{}
	for _, commandHandler := range c.handlers {
		commandDescriptors = append(commandDescriptors, commandHandler.Help())
	}
	return "Available commands are: \n" + strings.Join(commandDescriptors, "\n")
}
func (c Cli) Run(params []string) string {
	if params == nil || len(params) == 0 {
		panic("Missing arguments, run with `help` to see list of available actions")
	}
	action := params[0]

	if handler, ok := c.commandToHandlerMap[action]; ok {
		return handler.Process(params[1:])
	} else {
		return "Unknown action specified. " + c.Help()
	}
}
