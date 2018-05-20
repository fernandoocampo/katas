package fight

import "math"

func Combat(health, damage float64) float64 {
	if (health - damage) < 0 {
		return .0
	}
	return (health - damage)
}

func CombatTwo(health, damage float64) float64 {
	return math.Max(health-damage, 0.0)
}
