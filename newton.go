package main

import (
        "fmt"
        "math"
)

const NewtonIterations = 10
const SampleSize = 10000

type SqrtResult struct {
    Newtonian float64
    Librarian float64
}

func (r *SqrtResult) Diff() float64 {
        return math.Abs(r.Newtonian - r.Librarian)
}

func GetSqrtResultCustom(i int, iterations int) *SqrtResult {
    return &SqrtResult { SqrtCustom(float64(i), iterations), math.Sqrt(float64(i)) }
}
func GetSqrtResult(i int) *SqrtResult {
    return &SqrtResult { Sqrt(float64(i)), math.Sqrt(float64(i)) }
}
func SqrtCustom(x float64, iterations int) float64 {
    z := float64(1)
        for i := 1; i <= iterations; i++ {
                z = z - (z * z - x) / ( 2 * z)
        }
        return z
}

func Sqrt(x float64) float64 {
    z := float64(1)
        for i := 1; i <= NewtonIterations; i++ {
                z = z - (z * z - x) / ( 2 * z)
        }
        return z
}
func ComputeAggregatedDeltaCustom(iterations int) float64 {
    delta := 0.0
    for i := 1; i < SampleSize; i++ {
        delta += GetSqrtResultCustom(i, iterations).Diff()
    }
    return delta
}
func ComputeAggregatedDelta() float64 {
    delta := 0.0
    for i := 1; i < SampleSize; i++ {
        delta += GetSqrtResult(i).Diff()
    }
    return delta
}

func ComputeAndReportDeltaForCustom(iterations int, c chan float64) {
    v := ComputeAggregatedDeltaCustom(iterations)
    fmt.Printf("Total delta for iteration count [%d] is %f\n", iterations, v);
    c <- v
}
 
func main() {
    c := make(chan float64)
    go ComputeAndReportDeltaForCustom(8, c)
    go ComputeAndReportDeltaForCustom(9, c)
    go ComputeAndReportDeltaForCustom(10, c)
    <- c
    <- c
    <- c
    
}

