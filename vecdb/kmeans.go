package vecdb

import (
	"math"
	"math/rand"
)

type Centroid struct {
	Center Vector
	Points []Vector
	sum    Vector
}

type KMeans struct {
	NumClusters   int
	MaxIterations int
	RandomState   int64
	Centroids     []Centroid
	random        *rand.Rand
}

func NewKMeans(numClusters int, maxIterations int, randomState int64) KMeans {
	return KMeans{
		NumClusters:   numClusters,
		MaxIterations: maxIterations,
		RandomState:   randomState,
		random:        rand.New(rand.NewSource(randomState)),
	}
}

func (k *KMeans) Fit(vectors []Vector) {
	k.initializeCentroids(vectors)
	for i := 0; i < k.MaxIterations; i++ {
		k.estimateCentroids(vectors)
	}
}

func (k *KMeans) initializeCentroids(vectors []Vector) {
	k.Centroids = make([]Centroid, k.NumClusters)

	takenIndices := make(map[int]bool, k.NumClusters)
	for i := 0; i < k.NumClusters; {
		randomIdx := k.random.Intn(len(vectors))
		if !takenIndices[randomIdx] {
			takenIndices[randomIdx] = true
			k.Centroids[i].Center = vectors[randomIdx]
			k.Centroids[i].Points = make([]Vector, 0)
			i++
		}
	}
}

func (k *KMeans) estimateCentroids(vectors []Vector) {
	k.clearCentroidPoints()

	for i := 0; i < len(vectors); i++ {
		distance := math.MaxFloat64
		centroidIdx := -1

		for j := 0; j < len(k.Centroids); j++ {
			curDistance := vectors[i].EuclideanDistanceSquared(k.Centroids[j].Center)
			if math.Abs(curDistance) < math.Abs(distance) {
				distance = curDistance
				centroidIdx = j
			}
		}

		if centroidIdx >= 0 {
			centroid := &k.Centroids[centroidIdx]
			centroid.Points = append(k.Centroids[centroidIdx].Points, vectors[i])
			for dim := 0; dim < len(vectors[i]); dim++ {
				centroid.sum[dim] += vectors[i][dim]
			}
		} else {
			// todo: this shouldn't happen, but we should handle it
			panic("WTF happened?")
		}
	}

	for i := 0; i < len(k.Centroids); i++ {
		centroid := &k.Centroids[i]
		if len(centroid.Points) > 0 {
			for dim := 0; dim < len(centroid.Points[0]); dim++ {
				centroid.Center[dim] = centroid.sum[dim] / float64(len(centroid.Points))
			}
		}
	}
}

func (k *KMeans) clearCentroidPoints() {
	for i := 0; i < len(k.Centroids); i++ {
		k.Centroids[i].Points = make([]Vector, 0)
		k.Centroids[i].sum = make(Vector, 3)
	}
}
