package tests

import "testing"

func TestSubstract(t *testing.T) {
	memberOne := 10
	memberTwo := 5
	if result := Substract(memberOne, memberTwo); result != 5 {
		t.Errorf("Substract function is corrupt, %d - %d got %d, want %d", memberOne, memberTwo, result, 5)
	}
}
