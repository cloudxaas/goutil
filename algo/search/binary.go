package cxutilalgosearch

import "bytes"

// BinarySearchInt performs a binary search on a sorted integer slice.
// It returns the index of the target if found, or -1 if not found.
func BinaryInt(arr []int, target int) int {
    lo, hi := 0, len(arr)-1
    for lo <= hi {
        mid := int(uint(lo+hi) >> 1)
        switch {
        case arr[mid] < target:
            lo = mid + 1
        case arr[mid] > target:
            hi = mid - 1
        default:
            return mid
        }
    }
    return -1 // not found
}

// BinarySearchString performs a binary search on a sorted string slice.
// It returns the index of the target if found, or -1 if not found.
func BinaryString(arr []string, target string) int {
    lo, hi := 0, len(arr)-1
    for lo <= hi {
        mid := int(uint(lo+hi) >> 1)
        switch {
        case arr[mid] < target:
            lo = mid + 1
        case arr[mid] > target:
            hi = mid - 1
        default:
            return mid
        }
    }
    return -1 // not found
}

// BinarySearchBytes performs a binary search on a sorted []byte slice using bytes.Compare.
// It returns the index of the target if found, or -1 if not found.
// Note: This version may allocate memory in some cases.
func BinaryBytes1(arr [][]byte, target []byte) int {
    lo, hi := 0, len(arr)-1
    for lo <= hi {
        mid := int(uint(lo+hi) >> 1)
        cmp := bytes.Compare(arr[mid], target)
        switch {
        case cmp < 0:
            lo = mid + 1
        case cmp > 0:
            hi = mid - 1
        default:
            return mid
        }
    }
    return -1 // not found
}

// BinarySearchBytesZeroAlloc performs a binary search on a sorted []byte slice.
// It returns the index of the target if found, or -1 if not found.
// This version guarantees zero allocation.
func BinaryBytes(arr [][]byte, target []byte) int {
    lo, hi := 0, len(arr)-1
    for lo <= hi {
        mid := int(uint(lo+hi) >> 1)
        cmp := compareBytes(arr[mid], target)
        switch {
        case cmp < 0:
            lo = mid + 1
        case cmp > 0:
            hi = mid - 1
        default:
            return mid
        }
    }
    return -1 // not found
}

// compareBytes compares two byte slices without allocation.
// Returns -1 if a < b, 0 if a == b, and 1 if a > b.
func compareBytes(a, b []byte) int {
    minLen := len(a)
    if len(b) < minLen {
        minLen = len(b)
    }
    for i := 0; i < minLen; i++ {
        switch {
        case a[i] < b[i]:
            return -1
        case a[i] > b[i]:
            return 1
        }
    }
    switch {
    case len(a) < len(b):
        return -1
    case len(a) > len(b):
        return 1
    default:
        return 0
    }
}
