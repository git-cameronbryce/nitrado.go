package setup

import (
	"strings"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/handler"
	"github.com/disgoorg/omit"

	"github.com/git-cameronbryce/internal/nitrado"
)

var Definitions = []discord.ApplicationCommandCreate{
	clusterCommand,
}

func Register(r handler.Router) {
	r.SlashCommand("/ase-setup-cluster", handleCluster)
	r.ButtonComponent("/ase-setup-cluster/primary", handleSetupTokenButton)
	r.Modal("/ase-setup-cluster/modal", handleSetupTokenModal)
}

var clusterCommand = discord.SlashCommandCreate{
	Name:                     "ase-setup-cluster",
	Description:              "Set up your Nitrado Ark Survival Evolved cluster.",
	DefaultMemberPermissions: omit.NewPtr(discord.PermissionAdministrator),
}

func handleCluster(_ discord.SlashCommandInteractionData, e *handler.CommandEvent) error {
	embed := discord.NewEmbed().
		WithDescription("**Ark Survival Evolved**\n**Cluster Setup & Information**\nTo begin the setup process, you will need to obtain your token following this [video preview](https://example.com) allowing us to directly connect to your cluster.\n\n**Additional Information**\nEnsure [this](https://example.com) guild is a community server, allowing us to create forums and threads for logging data.").
		WithImage("https://i.imgur.com/2ZIHUgx.png").
		WithColor(0x2ecc71)

	msg := discord.NewMessageCreate().
		WithEphemeral(true).
		AddEmbeds(embed).
		AddActionRow(
			discord.NewSuccessButton("Setup Token", "/ase-setup-cluster/primary"),
			discord.NewLinkButton("Support Server", "https://example.com"),
		)

	return e.CreateMessage(msg)
}

func handleSetupTokenButton(_ discord.ButtonInteractionData, e *handler.ComponentEvent) error {
	tokenInput := discord.NewShortTextInput("token").
		WithPlaceholder("Enter your Nitrado API token...").
		WithRequired(true)

	modal := discord.NewModalCreate("/ase-setup-cluster/modal", "Setup Nitrado Token").
		AddLabel("Nitrado API Token", tokenInput)

	return e.Modal(modal)
}

func handleSetupTokenModal(e *handler.ModalEvent) error {
	token := strings.TrimSpace(e.Data.Text("token"))

	if err := nitrado.ValidateToken(e.Ctx, token); err != nil {
		return e.CreateMessage(
			discord.NewMessageCreate().
				WithContent("Invalid Nitrado token.").
				WithEphemeral(true),
		)
	}

	return e.CreateMessage(
		discord.NewMessageCreate().
			WithContent("Token received!").
			WithEphemeral(true),
	)
}
