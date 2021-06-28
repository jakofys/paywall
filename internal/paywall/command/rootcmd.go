package command

import (
	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "paywall",
	Short: "Paywall is calculator with lightning network implementation",
}

func Execute() {
	root.AddCommand(cmdAddUp)
	root.AddCommand(cmdSubstract)
	root.Execute()
}
