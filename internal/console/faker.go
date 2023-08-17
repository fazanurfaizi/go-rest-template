package console

import (
	"github.com/fazanurfaizi/go-rest-template/database/faker"
	"github.com/fazanurfaizi/go-rest-template/pkg/command"
	"github.com/spf13/cobra"
)

type FakerCommand struct{}

func (s *FakerCommand) Short() string {
	return "Generate faker data"
}

func (s *FakerCommand) Setup(cmd *cobra.Command) {}

func (s *FakerCommand) Run() command.CommandRunner {
	return func(
		faker faker.Fakers,
	) {
		faker.Setup()
	}
}

func NewFakerCommand() *FakerCommand {
	return &FakerCommand{}
}
