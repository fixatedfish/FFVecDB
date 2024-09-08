package vecdb

import (
	"math/rand"
	"reflect"
	"testing"
)

func getNonZeroVectors(numVectors int, dimensions int, initValue float64) []Vector {
	vectors := make([]Vector, numVectors)

	for i := 0; i < len(vectors); i++ {
		vectors[i] = make(Vector, dimensions)
		for j := 0; j < dimensions; j++ {
			vectors[i][j] = initValue
		}
	}

	return vectors
}

func getNonZeroVectorsInClusters(numVectors int, dimensions int, clusters int) []Vector {
	vectors := make([]Vector, numVectors)
	random := rand.New(rand.NewSource(1))

	vecCnt := 0
	for c := 0; c < clusters; c++ {
		for i := 0; i < numVectors/clusters; i++ {
			vector := make(Vector, dimensions)
			for j := 0; j < dimensions; j++ {
				vector[j] = float64(c*4) + random.Float64()
			}
			vectors[vecCnt] = vector
			vecCnt++
		}
	}

	return vectors
}

func TestKMeans_initializeCentroids(t *testing.T) {
	kmeans := NewKMeans(3, 100, 12)
	vectors := getNonZeroVectorsInClusters(9, 3, 3)
	kmeans.initializeCentroids(vectors)

	want := []Vector{
		{4.300911860585287, 4.5152126285020655, 4.813639960990097},
		{0.4377141871869802, 0.4246374970712657, 0.6868230728671094},
		{4.214263872582375, 4.380657189299686, 4.3180581743303295},
	}

	for i := 0; i < len(want); i++ {
		got := kmeans.Centroids[i].Center
		if !reflect.DeepEqual(got, want[i]) {
			t.Errorf("Expected centroid center to be %v, but got %v", want[i], got)
		}
	}
}

func TestFit(t *testing.T) {
	kmeans := NewKMeans(3, 10000, 0)
	vectors := getNonZeroVectorsInClusters(9, 3, 3)
	kmeans.Fit(vectors)

	want := 3
	for i := 0; i < len(kmeans.Centroids); i++ {
		got := len(kmeans.Centroids[i].Points)
		if got != 3 {
			t.Errorf("Expected centroids to have %v points, but got %v", want, got)
		}
	}
}
