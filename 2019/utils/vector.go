package utils

import (
	"fmt"
	"math"
)

type Vector struct {
	X, Y int
}

// Not really a normalization, just smallest possible direction vector that is from integers
func (v *Vector) Normalize() error {
	gcd := GCD(v.X, v.Y)
	if gcd == 0 {
		return fmt.Errorf("cannot normalize a (0,0) vector")
	}
	v.X /= gcd
	v.Y /= gcd
	return nil
}

func (v Vector) AngleWith(w Vector) float64 {
	// cos(theta) = dot(v,w) / |v|*|w|
	right := float64(v.X*w.X+v.Y*w.Y) / (v.Magnitude() * w.Magnitude())
	return math.Acos(right)
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(float64(v.X), 2) + math.Pow(float64(v.Y), 2))
}
