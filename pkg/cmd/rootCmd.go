package cmd

import (
	"github.com/spf13/cobra"
	"github.com/zulubit/mole/pkg/helpers"
)

var RootCmd = &cobra.Command{
	Use:   "mole",
	Short: "Micro-PaaS minimal in size and complexity.",
	Long:  "\n" + helpers.MoleAsciiArt() + "\nMole is a lightweight micro-PaaS solution optimized for Git-based deployments with Docker Compose and Caddy.\nFind more information at https://github.com/zulubit/mole.",
}
