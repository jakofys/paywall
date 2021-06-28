package tests

import (
	"testing"

	"github.com/jakofys/paywall/internal/paywall/handler"
)

//Test simple addition operation
func TestAddUp(t *testing.T) {
	memberOne := 5
	memberTwo := 10
	if result := handler.SumTwoMembers(memberOne, memberTwo); result != 15 {
		t.Errorf("AddUp function is corrupt, %d + %d got %d want %d", memberOne, memberTwo, result, 15)
	}
}

//Test simple substraction operation
func TestSubstract(t *testing.T) {
	memberOne := 10
	memberTwo := 5
	if result := handler.SubstractTwoMembers(memberOne, memberTwo); result != 5 {
		t.Errorf("Substract function is corrupt, %d - %d, got %d want %d", memberOne, memberTwo, result, 5)
	}
}
