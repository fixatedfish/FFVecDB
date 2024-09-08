package vecdb

import (
	"math"
)

type Vector []float64

func (v Vector) EuclideanDistance(v2 Vector) float64 {
	return EuclideanDistance(v, v2)
}

func EuclideanDistance(v1 Vector, v2 Vector) float64 {
	if len(v1) != len(v2) {
		return math.MaxFloat64
	}

	distance := float64(0)
	for i := 0; i < len(v1); i++ {
		distance += v1[i] - v2[i]*v1[i] - v2[i]
	}

	return math.Sqrt(distance)
}

func (v Vector) EuclideanDistanceSquared(v2 Vector) float64 {
	return EuclideanDistanceSquared(v, v2)
}

func EuclideanDistanceSquared(v1 Vector, v2 Vector) float64 {
	if len(v1) != len(v2) {
		return math.MaxFloat64
	}

	distance := float64(0)
	for i := 0; i < len(v1); i++ {
		distance += (v1[i] - v2[i]) * (v1[i] - v2[i])
	}

	return distance
}

func (v1 Vector) CosineSimilarity(v2 Vector) float64 {
	return CosineSimilarity(v1, v2)
}

func CosineSimilarity(v1 Vector, v2 Vector) float64 {
	return v1.DotProduct(v2) / (Magnitude(v1) * Magnitude(v2))
}

func (v Vector) DotProduct(v2 Vector) float64 {
	return DotProduct(v, v2)
}

func DotProduct(v1 Vector, v2 Vector) float64 {
	if len(v1) != len(v2) {
		return math.MaxFloat64
	}

	product := float64(0)
	for i := 0; i < len(v1); i++ {
		product += v1[i] * v2[i]
	}

	return product
}

func (v Vector) Magnitude() float64 {
	return Magnitude(v)
}

func Magnitude(v Vector) float64 {
	mag := float64(0)
	for i := 0; i < len(v); i++ {
		mag += v[i] * v[i]
	}
	return math.Sqrt(mag)
}
