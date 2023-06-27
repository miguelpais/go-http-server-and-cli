package cli

type CommandHandler interface {
	Identifier() string
	Help() string
	Process(args []string) string
}
