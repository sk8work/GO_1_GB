package triangle

import "math"

func GetHep(cat1, cat2 float64) float64 {
	return math.Sqrt(cat1*cat1 + cat2*cat2)
}

func GetSquare(cat1, cat2 float64) float64 {
	return ((cat1 * cat2) / 2)
}

func GetPerimetr(cat1, cat2 float64) float64 {
	return GetHep(cat1, cat2) + cat1 + cat2
}
