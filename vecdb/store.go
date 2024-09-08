package vecdb

// todo: implement
type VectorStore interface {
	AddVector(Vector)
	AddVectors([]Vector)
	Enumerate()
}

type Cluster struct {
	centeroid Vector
	points    []Vector // todo: could be a vector store?
}

type ClusteredVectorStore struct {
	clusters []Cluster
}
