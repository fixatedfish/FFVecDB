package vecdb

import "math"

type DatabaseConfig struct {
	dimensions  int
	Name        string
	VectorStore []Vector
}

type Database interface {
	AddVector(Vector)
	AddVectors([]Vector)
	Search(Vector, int)
}

func NewDatabaseConfig(dimensions int) DatabaseConfig {
	return DatabaseConfig{dimensions: dimensions, Name: "Test"}
}

type FlatDatabase struct {
	DatabaseConfig
}

func NewFlatDatabase(config DatabaseConfig) FlatDatabase {
	return FlatDatabase{DatabaseConfig: config}
}

func (db *FlatDatabase) AddVector(v Vector) {
	db.VectorStore = append(db.VectorStore, v)
}

func (db *FlatDatabase) AddVectors(vecs []Vector) {
	db.VectorStore = append(db.VectorStore, vecs...)
}

func (db *FlatDatabase) Search(vec Vector, i int) Vector {
	minDistance := math.MaxFloat64
	minVector := make(Vector, 3)
	for _, vector := range db.VectorStore {
		//dist := vector.EuclideanDistanceSquared(vec)
		dist := vector.CosineSimilarity(vec)
		dist = math.Abs(1 - dist)
		if dist < minDistance {
			minDistance = dist
			minVector = vector
		}
	}

	return minVector
}

// --------- CLUSTERED DATABASE ----------

type ClusteredDatabase struct {
	DatabaseConfig
	flatDatabase FlatDatabase
}

func NewClusteredDatabase(config DatabaseConfig, flatDatabase FlatDatabase) ClusteredDatabase {
	return ClusteredDatabase{DatabaseConfig: config, flatDatabase: flatDatabase}
}

func (db *ClusteredDatabase) Train() {

}
