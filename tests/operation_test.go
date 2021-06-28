package tests

import (
	"testing"

	"github.com/jakofys/paywall/internal/paywall/handler"
)

//Test simple addition operation
func TestAddUp(t *testing.T) {
	memberOne := "5"
	memberTwo := "10"
	if err := handler.AddUpHandler(nil, []string{memberOne, memberTwo}); err != nil {
		t.Errorf("AddUp function is corrupt, %s + %s got want %d", memberOne, memberTwo, 15)
	}
}

//Test simple substraction operation
func TestSubstract(t *testing.T) {
	memberOne := "10"
	memberTwo := "5"
	if err := handler.SubstractHandler(nil, []string{memberOne, memberTwo}); err != nil {
		t.Errorf("Substract function is corrupt, %s - %s, want %d", memberOne, memberTwo, 5)
	}
}
