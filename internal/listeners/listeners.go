package listeners

import (
	"log/slog"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/events"
)

func All() []bot.EventListener {
	return []bot.EventListener{
		bot.NewListenerFunc(onReady),
	}
}

func onReady(e *events.Ready) {
	slog.Info("bot is ready", "user", e.User.Username)
}
