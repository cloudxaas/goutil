package rps

import (
    "sync/atomic"
    "time"
    "unsafe"

    "github.com/panjf2000/ants/v2"
)

// Counter represents the RPS counter
type Counter struct {
    requestCount int32
    ticker       *time.Ticker
    done         chan struct{}
    buffer       []byte
    identifier   string
}

// New creates a new RPS counter with the given identifier
func New(identifier string) *Counter {
    return &Counter{
        done:       make(chan struct{}),
        buffer:     make([]byte, 128), // Increased buffer size to accommodate identifier
        identifier: identifier,
    }
}

// Increment increases the request count by 1
func (c *Counter) Increment() {
    atomic.AddInt32(&c.requestCount, 1)
}

// Start begins logging the requests per second
func (c *Counter) Start() error {
    c.ticker = time.NewTicker(1 * time.Second)
    return ants.Submit(func() {
        for {
            select {
            case <-c.ticker.C:
                count := atomic.SwapInt32(&c.requestCount, 0)
                c.formatAndWrite(count)
            case <-c.done:
                return
            }
        }
    })
}

// Stop halts the RPS counter
func (c *Counter) Stop() {
    if c.ticker != nil {
        c.ticker.Stop()
    }
    close(c.done)
}

// formatAndWrite formats the RPS count with identifier and writes it to stdout without allocations
func (c *Counter) formatAndWrite(count int32) {
    // Reset buffer
    buf := c.buffer[:0]

    // Append identifier
    buf = append(buf, c.identifier...)
    buf = append(buf, " RPS: "...)

    // Convert count to string and append
    buf = c.appendInt(buf, int64(count))

    // Append newline
    buf = append(buf, '\n')

    // Write to stdout
    _, _ = stdout.Write(buf)
}

// appendInt appends an integer to a byte slice without allocations
func (c *Counter) appendInt(buf []byte, x int64) []byte {
    if x < 0 {
        buf = append(buf, '-')
        x = -x
    }
    return c.appendUint(buf, uint64(x))
}

// appendUint appends an unsigned integer to a byte slice without allocations
func (c *Counter) appendUint(buf []byte, x uint64) []byte {
    if x < 10 {
        return append(buf, byte(x)+'0')
    }
    return c.appendUint(append(buf, byte(x%10)+'0'), x/10)
}

// stdout is a wrapper for os.Stdout that doesn't allocate
var stdout = struct {
    Write func(p []byte) (n int, err error)
}{
    Write: func(p []byte) (n int, err error) {
        return write(1, p)
    },
}

// write is a syscall.Write on Unix-like systems, and syscall.WriteFile on Windows
func write(fd int, p []byte) (n int, err error) {
    return syscall.Write(fd, p)
}
