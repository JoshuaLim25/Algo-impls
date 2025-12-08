# Go Patterns & Templates

## 1. Sliding Window
### Dynamic Sliding Window - Min Window
```go
func shortestWindow(nums []int, condition func() bool) int {
    i := 0
    minLength := math.MaxInt32

    for j := 0; j < len(nums); j++ {
        // Expand window - add nums[j] to current window logic

        // Shrink window while condition is met
        for condition() {
            if j-i+1 < minLength {
                minLength = j - i + 1
            }
            // Remove nums[i] from window logic
            i++
        }
    }

    if minLength == math.MaxInt32 {
        return -1
    }
    return minLength
}
```

### Dynamic Sliding Window - Max Window
```go
func longestWindow(nums []int, condition func() bool) int {
    i := 0
    maxLength := 0

    for j := 0; j < len(nums); j++ {
        // Expand window - add nums[j] to current window logic

        // Shrink window if condition is violated
        for !condition() {
            // Remove nums[i] from window logic
            i++
        }

        if j-i+1 > maxLength {
            maxLength = j - i + 1
        }
    }

    return maxLength
}
```

### Fixed Sliding Window
```go
func windowFixedSize(nums []int, k int) interface{} {
    i := 0
    for j := 0; j < len(nums); j++ {
        // Expand window - add nums[j] to current window logic

        // Ensure window has size k
        if j-i+1 < k {
            continue
        }

        // Process window of size k
        // Update result

        // Remove nums[i] from window, increment i
        i++
    }

    return nil // return result
}
```

## 2. Two Pointers

```go
func twoPointerTemplate(input []int) interface{} {
    i, j := 0, len(input)-1
    
    for i < j {
        // Process elements at both pointers
        
        // Adjust pointers based on conditions
        if /* condition */ {
            i++
        } else {
            j--
        }
    }
    
    return nil // return result
}
```

## 3. Slow and Fast Pointers

```go
type ListNode struct {
    Val  int
    Next *ListNode
}

func slowFastPointers(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }

    slow, fast := head, head

    // Move slow once, fast twice
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next

        // Custom logic (e.g., cycle detection)
        if fast == slow {
            // Cycle detected
            break
        }
    }

    return slow // or custom result
}
```

## 4. In-Place Linked List Reversal

```go
func reverseLinkedList(head *ListNode) *ListNode {
    var prev *ListNode
    curr := head

    for curr != nil {
        next := curr.Next
        curr.Next = prev
        prev = curr
        curr = next
    }

    return prev
}
```

## 5. Binary Search

### Classic Binary Search
```go
func binarySearch(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right-left)/2

        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return -1
}
```

### Rotated Array Search
```go
func searchRotated(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := (left + right) / 2

        if nums[mid] == target {
            return mid
        }

        // Left side sorted
        if nums[left] <= nums[mid] {
            if nums[left] <= target && target <= nums[mid] {
                right = mid - 1
            } else {
                left = mid + 1
            }
        } else {
            // Right side sorted
            if nums[mid] <= target && target <= nums[right] {
                left = mid + 1
            } else {
                right = mid - 1
            }
        }
    }

    return -1
}
```

## 6. Top K Elements

```go
import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

// Top K Largest (use min heap)
func topKLargest(nums []int, k int) []int {
    h := &IntHeap{}
    heap.Init(h)
    
    for _, num := range nums {
        heap.Push(h, num)
        if h.Len() > k {
            heap.Pop(h)
        }
    }
    
    return *h
}

// Top K Smallest (use max heap - negate values)
func topKSmallest(nums []int, k int) []int {
    h := &IntHeap{}
    heap.Init(h)
    
    for _, num := range nums {
        heap.Push(h, -num) // Negate for max heap
        if h.Len() > k {
            heap.Pop(h)
        }
    }
    
    result := make([]int, h.Len())
    for i := range result {
        result[i] = -(*h)[i] // Negate back
    }
    return result
}
```

## 7. Binary Tree Traversal

```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

// Preorder
func preorderTraversal(root *TreeNode) {
    if root == nil {
        return
    }
    
    // Process root
    fmt.Print(root.Val, " ")
    preorderTraversal(root.Left)
    preorderTraversal(root.Right)
}

// Inorder
func inorderTraversal(root *TreeNode) {
    if root == nil {
        return
    }
    
    inorderTraversal(root.Left)
    // Process root
    fmt.Print(root.Val, " ")
    inorderTraversal(root.Right)
}

// Postorder
func postorderTraversal(root *TreeNode) {
    if root == nil {
        return
    }
    
    postorderTraversal(root.Left)
    postorderTraversal(root.Right)
    // Process root
    fmt.Print(root.Val, " ")
}

// BFS
func bfsTraversal(root *TreeNode) {
    if root == nil {
        return
    }
    
    queue := []*TreeNode{root}
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        
        // Process node
        fmt.Print(node.Val, " ")
        
        if node.Left != nil {
            queue = append(queue, node.Left)
        }
        if node.Right != nil {
            queue = append(queue, node.Right)
        }
    }
}
```

## 8. Graph and Matrices

### DFS Graph
```go
func dfsGraph(graph map[int][]int, start int) []int {
    visited := make(map[int]bool)
    result := []int{}
    
    var dfs func(node int)
    dfs = func(node int) {
        visited[node] = true
        result = append(result, node)
        
        for _, neighbor := range graph[node] {
            if !visited[neighbor] {
                dfs(neighbor)
            }
        }
    }
    
    dfs(start)
    return result
}
```

### BFS Graph
```go
func bfsGraph(graph map[int][]int, start int) []int {
    visited := make(map[int]bool)
    result := []int{}
    queue := []int{start}
    
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        
        if visited[node] {
            continue
        }
        
        visited[node] = true
        result = append(result, node)
        
        for _, neighbor := range graph[node] {
            if !visited[neighbor] {
                queue = append(queue, neighbor)
            }
        }
    }
    
    return result
}
```

### DFS Matrix
```go
func dfsMatrix(matrix [][]int) []int {
    if len(matrix) == 0 {
        return []int{}
    }
    
    m, n := len(matrix), len(matrix[0])
    visited := make([][]bool, m)
    for i := range visited {
        visited[i] = make([]bool, n)
    }
    
    result := []int{}
    directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    
    var dfs func(i, j int)
    dfs = func(i, j int) {
        if i < 0 || i >= m || j < 0 || j >= n || visited[i][j] {
            return
        }
        
        visited[i][j] = true
        result = append(result, matrix[i][j])
        
        for _, dir := range directions {
            dfs(i+dir[0], j+dir[1])
        }
    }
    
    // Start DFS from all unvisited cells
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if !visited[i][j] {
                dfs(i, j)
            }
        }
    }
    
    return result
}
```

## 9. Backtracking

```go
func backtrackTemplate(candidates []int, target interface{}) [][]int {
    result := [][]int{}
    path := []int{}
    
    var backtrack func(start int)
    backtrack = func(start int) {
        // Base case
        if /* solution found condition */ {
            // Make copy of path
            temp := make([]int, len(path))
            copy(temp, path)
            result = append(result, temp)
            return
        }
        
        for i := start; i < len(candidates); i++ {
            // Skip invalid candidates
            if /* invalid condition */ {
                continue
            }
            
            // Choose
            path = append(path, candidates[i])
            
            // Explore
            backtrack(i + 1) // or backtrack(i) for combinations with repetition
            
            // Unchoose
            path = path[:len(path)-1]
        }
    }
    
    backtrack(0)
    return result
}
```

## 10. Dynamic Programming

### Top-Down with Memoization
```go
func fibMemo(n int, memo map[int]int) int {
    if val, exists := memo[n]; exists {
        return val
    }
    
    if n <= 1 {
        return n
    }
    
    memo[n] = fibMemo(n-1, memo) + fibMemo(n-2, memo)
    return memo[n]
}
```

### Bottom-Up with Array
```go
func fibArray(n int) int {
    if n <= 1 {
        return n
    }
    
    dp := make([]int, n+1)
    dp[0], dp[1] = 0, 1
    
    for i := 2; i <= n; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    
    return dp[n]
}
```

### Bottom-Up Space Optimized
```go
func fibOptimized(n int) int {
    if n <= 1 {
        return n
    }
    
    prev2, prev1 := 0, 1
    
    for i := 2; i <= n; i++ {
        current := prev1 + prev2
        prev2 = prev1
        prev1 = current
    }
    
    return prev1
}
```

## 11. Bit Manipulation

```go
func bitOperations(a, b int) map[string]int {
    return map[string]int{
        "AND":        a & b,
        "OR":         a | b,
        "XOR":        a ^ b,
        "NOT":        ^a,
        "LeftShift":  a << b,
        "RightShift": a >> b,
        "LSB":        a & 1,
    }
}

// Count set bits
func countBits(n int) int {
    count := 0
    for n != 0 {
        count++
        n &= n - 1 // Remove rightmost set bit
    }
    return count
}
```

## 12. Overlapping Intervals

```go
func processIntervals(intervals [][]int) [][]int {
    if len(intervals) == 0 {
        return [][]int{}
    }
    
    // Sort by start time
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    result := [][]int{intervals[0]}
    
    for i := 1; i < len(intervals); i++ {
        current := intervals[i]
        last := result[len(result)-1]
        
        // No overlap
        if last[1] < current[0] {
            result = append(result, current)
        } else {
            // Merge overlapping intervals
            last[1] = max(last[1], current[1])
        }
    }
    
    return result
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
```

## 13. Monotonic Stack

### Monotonic Increasing Stack
```go
func monotonicIncreasing(nums []int) []int {
    stack := [][]int{} // [value, index]
    result := []int{}
    
    for i, num := range nums {
        // Pop elements that are greater than current
        for len(stack) > 0 && stack[len(stack)-1][0] > num {
            stack = stack[:len(stack)-1]
        }
        
        // Process with top of stack if needed
        if len(stack) > 0 {
            // Custom logic here
        }
        
        stack = append(stack, []int{num, i})
    }
    
    return result
}
```

### Monotonic Decreasing Stack
```go
func monotonicDecreasing(nums []int) []int {
    stack := [][]int{} // [value, index]
    result := []int{}
    
    for i, num := range nums {
        // Pop elements that are smaller than current
        for len(stack) > 0 && stack[len(stack)-1][0] < num {
            stack = stack[:len(stack)-1]
        }
        
        // Process with top of stack if needed
        if len(stack) > 0 {
            // Custom logic here
        }
        
        stack = append(stack, []int{num, i})
    }
    
    return result
}
```

## 14. Prefix Sum

```go
func buildPrefixSum(nums []int) []int {
    if len(nums) == 0 {
        return []int{}
    }
    
    prefix := make([]int, len(nums))
    prefix[0] = nums[0]
    
    for i := 1; i < len(nums); i++ {
        prefix[i] = prefix[i-1] + nums[i]
    }
    
    return prefix
}

func querySubarraySum(prefix []int, left, right int) int {
    if left == 0 {
        return prefix[right]
    }
    return prefix[right] - prefix[left-1]
}
```

## Learning Order Recommendation

1. **Two Pointers** - Foundation for many problems
2. **Sliding Window** - Builds on two pointers
3. **Binary Search** - Essential optimization technique
4. **Binary Tree Traversal** - Core tree algorithms
5. **Graph/Matrix Traversal** - DFS/BFS fundamentals
6. **Dynamic Programming** - Start with 1D, then 2D
7. **Backtracking** - Builds on recursion/DFS
8. **Prefix Sum** - Simple but powerful technique
9. **Monotonic Stack** - Specialized stack usage
10. **Fast/Slow Pointers** - Linked list specific
11. **Linked List Reversal** - Linked list manipulation
12. **Top K Elements** - Heap-based problems
13. **Overlapping Intervals** - Sorting + greedy
14. **Bit Manipulation** - Mathematical operations

Focus on mastering patterns 1-8 first, as they cover 80% of problems.
