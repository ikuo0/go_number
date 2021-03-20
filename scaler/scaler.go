
package scaler

import (
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

func (me* StandardizationParams) Fit() {
}

func (me* StandardizationParams) Transform(x n2df.N) {
    me.Mean = n2df.Mean(x)
    me.Variance = n2df.Variance(x, 1)
    me.Std = n2df.Sqrt(me.Variance)
}
