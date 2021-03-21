
package kmeans

import (
    "github.com/ikuo0/go_number/n1df"
    "github.com/ikuo0/go_number/n2df"
)

type KMenas struct {
    Clusters int
    Threshold float64
    InitMeans n2df.N
    Means n2df.N
}

func New(clusters int, threshold float64) {
    return KMenas{
        Clusters: clusters,
        Threshold: threshold,
    }
}

func (me* KMenas) InitRandom(x n2df.N) {
    indexes := n1di.Arange(0, len(x))
    indexes = n1di.Shuffle(indexes)
    indexes = indexes[:me.Clusters]
    me.InitMeans = n2df.IndexingM(x, indexes)
}

func (me* KMenas) Fit(x n2df.N) {
    me.InitRandom()
    
}

