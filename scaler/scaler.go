
package scaler

import (
    "github.com/ikuo0/go_number/n1df"
    "github.com/ikuo0/go_number/n2df"
)

type StandardizationParams struct {
    Mean n1df.N
    Variance n1df.N
    Std n1df.N
}

func Standardization() (StandardizationParams) {
    return StandardizationParams{}
}

func (me* StandardizationParams) Fit(x n2df.N) {
    me.Mean = n2df.MeanM(x)
    me.Variance = n2df.VarianceM(x, 1)
    me.Std = n1df.Sqrt(me.Variance)
}

func (me* StandardizationParams) Transform(x n2df.N) n2df.N {
    x = n2df.SubtractM(x, me.Mean)
    x = n2df.DivisionM(x, me.Std)
    return x
}

func (me* StandardizationParams) FitTransform(x n2df.N) n2df.N {
    me.Fit(x)
    return me.Transform(x)
}
