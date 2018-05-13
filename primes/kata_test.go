package primes_test

import (
	"testing"

	"github.com/fernandoocampo/katas/primes"
)

func Test_step(t *testing.T) {
	// GIVEN the following parameters
	tables := []struct {
		x int
		y int
		z int
		n []int
	}{
		{2, 100, 110, []int{101, 103}},
		{4, 100, 110, []int{103, 107}},
		{6, 100, 110, []int{101, 107}},
		{8, 300, 400, []int{359, 367}},
		{10, 300, 400, []int{307, 317}},
		{11, 30000, 100000, nil},
	}

	// WHEN execute the step function
	for _, step := range tables {
		result := primes.Step(step.x, step.y, step.z)
		if step.n == nil && result != nil {
			t.Errorf("Step of (%d,%d,%d) was incorrect, got: [%d,%d], want: %d.", step.x, step.y, step.z, result[0], result[1], step.n)
		}

		if step.n != nil && result == nil {
			t.Errorf("Step of (%d,%d,%d) was incorrect, got: nil, want: %d.", step.x, step.y, step.z, step.n)
		}
		if step.n != nil && result != nil && step.n[0] != result[0] && step.n[1] != result[1] {
			t.Errorf("Step of (%d,%d,%d) was incorrect, got: [%d,%d], want: %d.", step.x, step.y, step.z, result[0], result[1], step.n)
		}
	}
}

func Test_step2(t *testing.T) {
	// GIVEN the following parameters
	tables := []struct {
		x int
		y int
		z int
		n []int
	}{
		{2, 100, 110, []int{101, 103}},
		{4, 100, 110, []int{103, 107}},
		{6, 100, 110, []int{101, 107}},
		{8, 300, 400, []int{359, 367}},
		{10, 300, 400, []int{307, 317}},
		{11, 30000, 100000, nil},
	}

	// WHEN execute the step function
	for _, step := range tables {
		result := primes.Step2(step.x, step.y, step.z)
		if step.n == nil && result != nil {
			t.Errorf("Step of (%d,%d,%d) was incorrect, got: [%d,%d], want: %d.", step.x, step.y, step.z, result[0], result[1], step.n)
		}

		if step.n != nil && result == nil {
			t.Errorf("Step of (%d,%d,%d) was incorrect, got: nil, want: %d.", step.x, step.y, step.z, step.n)
		}
		if step.n != nil && result != nil && step.n[0] != result[0] && step.n[1] != result[1] {
			t.Errorf("Step of (%d,%d,%d) was incorrect, got: [%d,%d], want: %d.", step.x, step.y, step.z, result[0], result[1], step.n)
		}
	}
}
