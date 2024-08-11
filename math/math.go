package cxutilmath

import (
    "golang.org/x/exp/constraints"
)

// Max returns the greater of two values of an ordered type.
func Max[T constraints.Ordered](a, b T) T {
    if a > b {
        return a
    }
    return b
}

// Min returns the lesser of two values of an ordered type.
func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

// IsWithinThreshold checks if the absolute difference between a and b is not more than threshold.
// It is generic over numeric types (integers and floats).
func IsWithinThreshold[T constraints.Float | constraints.Integer](a, b, threshold T) bool {
    diff := a - b
    if diff < 0 {
        diff = -diff
    }
    return diff <= threshold
}
