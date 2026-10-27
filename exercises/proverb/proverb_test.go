package proverb

import "testing"

func TestProverbBasic(t *testing.T) {
	input := []string{"nail", "shoe", "horse"}
	want := []string{
		"For want of a nail the shoe was lost.",
		"For want of a shoe the horse was lost.",
		"And all for the want of a nail.",
	}

	got := Proverb(input)
	if len(got) != len(want) {
		t.Fatalf("expected %d lines, got %d", len(want), len(got))
	}

	for i := range want {
		if got[i] != want[i] {
			t.Errorf("line %d = %q; want %q", i, got[i], want[i])
		}
	}
}

func TestEmpty(t *testing.T) {
	got := Proverb([]string{})
	if len(got) != 0:
		t.Errorf("expected empty slice, got %v", got)
	}
}

func TestOptimizedVersion(t *testing.T) {
	input := []string{"eye", "vision"}
	want := []string{
		"For want of a eye the vision was lost.",
		"And all for the want of a eye.",
	}

	got := ProverbOptimized(input)
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("optimized mismatch at %d: got %q want %q", i, got[i], want[i])
		}
	}
}
