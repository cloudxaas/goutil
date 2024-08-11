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

// IsWithinDifference checks if the difference between a and b is not more than c.
// It is generic over any ordered type.
func IsWithinDifference[T constraints.Ordered](a, b, c T) bool {
    if a > b {
        return a-b <= c
    }
    return b-a <= c
}
