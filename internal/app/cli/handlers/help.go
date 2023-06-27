package handlers

type HelpHandler struct{}

func (h HelpHandler) Identifier() string {
	return "help"
}

func (h HelpHandler) Help() string {
	return "help: describes all the available commands"
}

func (h HelpHandler) Process(args []string) string {
	panic("implement me")
}
