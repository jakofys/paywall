package handler

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func AddUpHandler(cmd *cobra.Command, args []string) error {
	member1, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	member2, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}
	fmt.Printf("Sum of %d + %d equal to %d \n", member1, member2, member1+member2)
	return nil
}
