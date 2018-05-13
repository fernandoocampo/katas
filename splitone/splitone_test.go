package splitone_test

import (
	"testing"

	"github.com/fernandoocampo/katas/splitone"
)

func TestSolution(t *testing.T) {
	// GIVEN the following parameters
	tables := []struct {
		a string
		b []string
	}{
		{"abc", []string{"ab", "c_"}},
		{"abcdef", []string{"ab", "cd", "ef"}},
		{"abcdefg", []string{"ab", "cd", "ef", "g_"}},
	}

	for _, item := range tables {
		result := splitone.Solution(item.a)

		if item.b == nil && result != nil {
			t.Errorf("expected nil, but it got: %v", result)
		}

		if item.b != nil && result == nil {
			t.Errorf("expected %v, but it got: nil", item.b)
		}

		if item.b != nil && result != nil {
			for i := 0; i < len(result); i++ {
				if item.b[i] != result[i] {
					t.Errorf("expected %v, but it got: %v", item.b, result)
				}
			}
		}
	}
}

func TestSolutionTwo(t *testing.T) {
	// GIVEN the following parameters
	tables := []struct {
		a string
		b []string
	}{
		{"abc", []string{"ab", "c_"}},
		{"abcdef", []string{"ab", "cd", "ef"}},
		{"abcdefg", []string{"ab", "cd", "ef", "g_"}},
	}

	for _, item := range tables {
		result := splitone.SolutionTwo(item.a)

		if item.b == nil && result != nil {
			t.Errorf("expected nil, but it got: %v", result)
		}

		if item.b != nil && result == nil {
			t.Errorf("expected %v, but it got: nil", item.b)
		}

		if item.b != nil && result != nil {
			for i := 0; i < len(result); i++ {
				if item.b[i] != result[i] {
					t.Errorf("expected %v, but it got: %v", item.b, result)
				}
			}
		}
	}
}

func TestSolutionThree(t *testing.T) {
	// GIVEN the following parameters
	tables := []struct {
		a string
		b []string
	}{
		{"abc", []string{"ab", "c_"}},
		{"abcdef", []string{"ab", "cd", "ef"}},
		{"abcdefg", []string{"ab", "cd", "ef", "g_"}},
	}

	for _, item := range tables {
		result := splitone.SolutionThree(item.a)

		if item.b == nil && result != nil {
			t.Errorf("expected nil, but it got: %v", result)
		}

		if item.b != nil && result == nil {
			t.Errorf("expected %v, but it got: nil", item.b)
		}

		if item.b != nil && result != nil {
			for i := 0; i < len(result); i++ {
				if item.b[i] != result[i] {
					t.Errorf("expected %v, but it got: %v", item.b, result)
				}
			}
		}
	}
}
