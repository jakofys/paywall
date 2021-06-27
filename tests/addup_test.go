package tests

import "testing"

func TestAddUp(t *testing.T) {
	memberOne := 5
	memberTwo := 10
	if result := AddUp(memberOne, memberTwo); result != 15 {
		t.Errorf("AddUp function is corrupt, %d + %d got %d, want %d", memberOne, memberTwo, result, 15)
	}
}
