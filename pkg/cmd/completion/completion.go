package completion

import (
	"github.com/{{ .Owner }}/{{ .Repo }}/internal/factory"
	"github.com/{{ .Owner }}/{{ .Repo }}/pkg/cmd/completion/bash"
	"github.com/{{ .Owner }}/{{ .Repo }}/pkg/cmd/completion/fish"
	"github.com/{{ .Owner }}/{{ .Repo }}/pkg/cmd/completion/zsh"
	"github.com/spf13/cobra"
)

func NewCompletionCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   f.Localizer.LocalizeByID("completion.cmd.use"),
		Short: f.Localizer.LocalizeByID("completion.cmd.shortDescription"),
		Long:  f.Localizer.LocalizeByID("completion.cmd.longDescription"),
		Args:  cobra.ExactArgs(1),
	}

	cmd.AddCommand(
		bash.NewCommand(f),
		zsh.NewCommand(f),
		fish.NewCommand(f),
	)

	return cmd
}
