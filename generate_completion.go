package urakil

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func GenerateCompletion(flag *pflag.FlagSet) error {
	command := &cobra.Command{
		Use: "completions",
	}
	command.Flags().AddFlagSet(flag)
	os.Mkdir("cmd/urakil/completions/", 0755)
	os.Mkdir("cmd/urakil/completions/bash", 0755)
	os.Mkdir("cmd/urakil/completions/zsh", 0755)
	os.Mkdir("cmd/urakil/completions/fish", 0755)
	os.Mkdir("cmd/urakil/completions/powershell", 0755)
	command.GenBashCompletionFileV2("cmd/urakil/completions/bash/urakil", true)
	command.GenZshCompletionFile("cmd/urakil/completions/zsh/urakil")
	command.GenFishCompletionFile("cmd/urakil/completions/fish/urakil", true)
	command.GenPowerShellCompletionFile("cmd/urakil/completions/ps1/urakil")
	return nil
}
