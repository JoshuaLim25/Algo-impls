```go
import "container/heap"

type IntHeap []int
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

// Usage:
h := &IntHeap{}
heap.Init(h)
heap.Push(h, 42)
min := heap.Pop(h).(int)
```

# Max heap (flip Less)
```go
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } // Only change
```

# Custom struct heap
```go
type Item struct { val, priority int }
type ItemHeap []Item
func (h ItemHeap) Less(i, j int) bool { return h[i].priority < h[j].priority }
// ... rest same pattern
```

- **Less() determines min/max** - `<` for min heap, `>` for max heap
- **Always use `&IntHeap{}`** and `heap.Init(h)` 
- **Type assert on Pop()** - `heap.Pop(h).(int)`
