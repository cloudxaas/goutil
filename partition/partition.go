// Package partition provides utilities for dividing totals into groups.
package partition

import (
	"errors"
	"math"
)

// DistributeEvenlyWithinLimit distributes a total number evenly into groups,
// respecting a maximum group size limit. It returns the number of groups and
// the size of each group. The total must be an even number.
//
// Parameters:
//   - total: The total number to be distributed (must be even)
//   - maxGroupSize: The maximum allowed size for each group
//
// Returns:
//   - numGroups: The number of groups
//   - groupSize: The size of each group
//   - error: An error if the total is not even
func DistributeEvenlyWithinLimit(total, maxGroupSize int) (numGroups, groupSize int, err error) {
	if total%2 != 0 {
		return 0, 0, errors.New("total must be an even number")
	}

	if total <= maxGroupSize {
		return 1, total, nil
	}

	numGroups = int(math.Ceil(float64(total) / float64(maxGroupSize)))
	groupSize = total / numGroups

	// Adjust group size and number of groups
	for groupSize > maxGroupSize || total%numGroups != 0 {
		numGroups++
		groupSize = total / numGroups
	}

	return numGroups, groupSize, nil
}
