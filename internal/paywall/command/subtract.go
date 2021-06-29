package command

import (
	"github.com/jakofys/paywall/internal/paywall/handler"
	"github.com/spf13/cobra"
)

var cmdSubstract = &cobra.Command{
	Use:      "subtract",
	Short:    "Substract two operation member",
	RunE:     handler.SubstractHandler,
	PostRunE: handler.AmendTransactionHandler,
	Args:     cobra.MinimumNArgs(2),
}
