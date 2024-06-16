package cxutilfmtreadable

import (
	"strconv"
)

// AppendBytes converts bytes to a human-readable string with units and appends to the provided buffer
func AppendBytes(buf []byte, b uint64) []byte {
	const unit = 1024
	if b < unit {
		return strconv.AppendUint(buf, b, 10)
	}
	div, exp := unit, 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	buf = strconv.AppendFloat(buf, float64(b)/float64(div), 'f', 1, 64)
	return append(buf, "KMGTPE"[exp], 'B')
}
