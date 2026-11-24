package commands

import (
	"slices"

	"github.com/disgoorg/disgo/handler"

	"github.com/git-cameronbryce/internal/commands/setup"
)

var Definitions = slices.Concat(
	setup.Definitions,
)

func NewRouter() *handler.Mux {
	r := handler.New()
	setup.Register(r)
	return r
}
