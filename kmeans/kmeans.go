
package kmeans

import (
    "github.com/ikuo0/go_number/n1df"
    "github.com/ikuo0/go_number/n1di"
    "github.com/ikuo0/go_number/n2df"
)

type Model struct {
    Clusters int
    Iteration int
    Threshold float64
    InitMeans n2df.N
    Means n2df.N
}

func New(clusters int, iteration int, threshold float64) Model {
    return Model {
        Clusters: clusters,
        Iteration: iteration,
        Threshold: threshold,
    }
}

func (me* Model) InitRandom(x n2df.N) {
    indexes := n1di.Arange(0, len(x))
    indexes = n1di.Shuffle(indexes)
    indexes = indexes[:me.Clusters]
    me.InitMeans = n2df.IndexingM(x, indexes)
}

func (me* Model) EStep(x n2df.N, means n2df.N) n1di.N {
    distances := n2df.New(len(x), me.Clusters)
    for m, row := range(x) {
        for cluster := 0; cluster < me.Clusters; cluster += 1 {
            diff := n1df.SubtractI(row, means[cluster])
            pow := n1df.Power(diff, 2)
            distance := n1df.Total(pow)
            distances[m][cluster] = distance
        }
    }
    
    predict := n1di.New(len(x))
    for m, distance := range(distances) {
        pred := n1df.ArgMin(distance)
        predict[m] = pred
    }
}

func (me* Model) MStep(x n2df.N, predict n1di.N) {
}

func (me* Model) Fit(x n2df.N) {
    me.InitRandom(x)
    
}

