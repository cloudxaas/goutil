package cxutilalgosearch

import (
    "testing"
)

func TestBinarySearchInt(t *testing.T) {
    tests := []struct {
        name     string
        arr      []int
        target   int
        expected int
    }{
        {"Found in middle", []int{1, 3, 5, 7, 9}, 5, 2},
        {"Found at beginning", []int{1, 3, 5, 7, 9}, 1, 0},
        {"Found at end", []int{1, 3, 5, 7, 9}, 9, 4},
        {"Not found", []int{1, 3, 5, 7, 9}, 4, -1},
        {"Empty array", []int{}, 1, -1},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := BinarySearchInt(tt.arr, tt.target); got != tt.expected {
                t.Errorf("BinarySearchInt() = %v, want %v", got, tt.expected)
            }
        })
    }
}

func TestBinarySearchString(t *testing.T) {
    tests := []struct {
        name     string
        arr      []string
        target   string
        expected int
    }{
        {"Found in middle", []string{"apple", "banana", "cherry", "date", "elderberry"}, "cherry", 2},
        {"Found at beginning", []string{"apple", "banana", "cherry", "date", "elderberry"}, "apple", 0},
        {"Found at end", []string{"apple", "banana", "cherry", "date", "elderberry"}, "elderberry", 4},
        {"Not found", []string{"apple", "banana", "cherry", "date", "elderberry"}, "fig", -1},
        {"Empty array", []string{}, "apple", -1},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := BinarySearchString(tt.arr, tt.target); got != tt.expected {
                t.Errorf("BinarySearchString() = %v, want %v", got, tt.expected)
            }
        })
    }
}

func TestBinarySearchBytes(t *testing.T) {
    tests := []struct {
        name     string
        arr      [][]byte
        target   []byte
        expected int
    }{
        {"Found in middle", [][]byte{[]byte("apple"), []byte("banana"), []byte("cherry"), []byte("date"), []byte("elderberry")}, []byte("cherry"), 2},
        {"Found at beginning", [][]byte{[]byte("apple"), []byte("banana"), []byte("cherry"), []byte("date"), []byte("elderberry")}, []byte("apple"), 0},
        {"Found at end", [][]byte{[]byte("apple"), []byte("banana"), []byte("cherry"), []byte("date"), []byte("elderberry")}, []byte("elderberry"), 4},
        {"Not found", [][]byte{[]byte("apple"), []byte("banana"), []byte("cherry"), []byte("date"), []byte("elderberry")}, []byte("fig"), -1},
        {"Empty array", [][]byte{}, []byte("apple"), -1},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := BinarySearchBytes(tt.arr, tt.target); got != tt.expected {
                t.Errorf("BinarySearchBytes() = %v, want %v", got, tt.expected)
            }
            if got := BinarySearchBytesZeroAlloc(tt.arr, tt.target); got != tt.expected {
                t.Errorf("BinarySearchBytesZeroAlloc() = %v, want %v", got, tt.expected)
            }
        })
    }
}

func BenchmarkBinarySearchInt(b *testing.B) {
    arr := make([]int, 1000)
    for i := range arr {
        arr[i] = i
    }
    b.ResetTimer()
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        BinarySearchInt(arr, 500)
    }
}

func BenchmarkBinarySearchString(b *testing.B) {
    arr := []string{"apple", "banana", "cherry", "date", "elderberry"}
    b.ResetTimer()
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        BinarySearchString(arr, "cherry")
    }
}

func BenchmarkBinarySearchBytes(b *testing.B) {
    arr := [][]byte{
        []byte("apple"),
        []byte("banana"),
        []byte("cherry"),
        []byte("date"),
        []byte("elderberry"),
    }
    target := []byte("cherry")
    b.ResetTimer()
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        BinarySearchBytes(arr, target)
    }
}

func BenchmarkBinarySearchBytesZeroAlloc(b *testing.B) {
    arr := [][]byte{
        []byte("apple"),
        []byte("banana"),
        []byte("cherry"),
        []byte("date"),
        []byte("elderberry"),
    }
    target := []byte("cherry")
    b.ResetTimer()
    b.ReportAllocs()
    for i := 0; i < b.N; i++ {
        BinarySearchBytesZeroAlloc(arr, target)
    }
}
