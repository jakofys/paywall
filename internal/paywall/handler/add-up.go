package handler

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// AddUp sum to member
func SumTwoMembers(member1 int, member2 int) int {
	return member1 + member2
}

//Handler to trigger on "add" call command
func AddUpHandler(cmd *cobra.Command, args []string) error {

	// Convert each member of operation from string to int
	member1, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	member2, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}
	fmt.Printf("Sum of %d + %d equal to %d \n", member1, member2, SumTwoMembers(member1, member2))
	return nil
}
