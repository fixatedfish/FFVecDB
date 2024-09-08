package main

import (
	"FFVecDB/vecdb"
)

func main() {
	db := vecdb.NewFlatDatabase(vecdb.NewDatabaseConfig(3))

	vec := make([]float64, 3)
	vec[0] = 0.6
	vec[1] = 0.7
	vec[2] = 0.5
	db.AddVector(vec)

	vec = make([]float64, 3)
	vec[0] = 0.5
	vec[1] = 0.5
	vec[2] = 0.5
	db.AddVector(vec)

	searchVec := make([]float64, 3)
	searchVec[0] = 0.5
	searchVec[1] = 0.5
	searchVec[2] = 0.5
	result := db.Search(searchVec, 4)
	for i := 0; i < len(result); i++ {
		println(result[i])
	}
}
