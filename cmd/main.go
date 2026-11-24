package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/snowflake/v2"
	"github.com/joho/godotenv"

	"github.com/git-cameronbryce/internal/commands"
	"github.com/git-cameronbryce/internal/listeners"
)

func main() {
	_ = godotenv.Load()

	guildID, err := snowflake.Parse(os.Getenv("GUILD_ID"))
	if err != nil {
		panic(err)
	}

	eventListeners := append(listeners.All(), commands.NewRouter())

	client, err := disgo.New(os.Getenv("TOKEN"),
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(
				gateway.IntentGuilds,
			),
		),
		bot.WithEventListeners(eventListeners...),
	)
	if err != nil {
		panic(err)
	}

	if err = handler.SyncCommands(client, commands.Definitions, []snowflake.ID{guildID}); err != nil {
		panic(err)
	}

	if err = client.OpenGateway(context.TODO()); err != nil {
		panic(err)
	}

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
	<-s
}
