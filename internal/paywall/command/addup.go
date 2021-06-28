package command

import (
	"github.com/jakofys/paywall/internal/paywall/handler"
	"github.com/spf13/cobra"
)

var cmdAddUp = &cobra.Command{
	Use:     "add",
	Short:   "Sum two operation member",
	RunE:    handler.AddUpHandler,
	PostRun: handler.AmendTransactionHandler,
	Args:    cobra.MinimumNArgs(2),
}
