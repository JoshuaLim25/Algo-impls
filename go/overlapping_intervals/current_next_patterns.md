# Current + Next Pattern Guide

## Pattern Template
```go
// Sort data appropriately first
sort.Slice(data, sortFunction)

// Initialize
cur := 0
result := 0

// Main loop: cur vs next
for next := 1; next < len(data); next++ {
    current := data[cur]
    nextItem := data[next]
    
    if conditionMet(current, nextItem) {
        // Handle the condition (count, remove, etc.)
        result++
        // Decision: which pointer to advance?
        // Option A: Keep cur same (when current is "better")
        // Option B: Move cur = next (when next is "better")
    } else {
        // No condition met, move cur forward
        cur = next
    }
}
```

## Problem: Non-Overlapping Intervals
**Your way - Current + Next pattern:**

```go
func eraseOverlapIntervals(intervals [][]int) int {
    if len(intervals) <= 1 {
        return 0
    }
    
    // Sort by end time (greedy: keep earliest ending)
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1]
    })
    
    cur := 0
    erased := 0
    
    for next := 1; next < len(intervals); next++ {
        current := intervals[cur]
        nextInterval := intervals[next]
        
        if current[1] > nextInterval[0] {  // Overlap detected
            erased++
            // Don't move cur - current ends earlier, so keep it
        } else {
            // No overlap, move cur to next
            cur = next
        }
    }
    
    return erased
}
```

## Problem: Merge Intervals
**Current + Next pattern:**

```go
func merge(intervals [][]int) [][]int {
    if len(intervals) <= 1 {
        return intervals
    }
    
    // Sort by start time
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    
    result := [][]int{}
    cur := 0
    
    for next := 1; next < len(intervals); next++ {
        current := intervals[cur]
        nextInterval := intervals[next]
        
        if current[1] >= nextInterval[0] {  // Overlap - merge them
            // Create merged interval
            merged := []int{current[0], max(current[1], nextInterval[1])}
            intervals[next] = merged  // Update next with merged result
            // Don't add current to result, move cur to next
            cur = next
        } else {
            // No overlap, add current to result
            result = append(result, current)
            cur = next
        }
    }
    
    // Add the last interval
    result = append(result, intervals[cur])
    return result
}
```

## Problem: Remove Duplicates from Sorted Array
**Current + Next pattern:**

```go
func removeDuplicates(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }
    
    cur := 0  // Position for unique elements
    
    for next := 1; next < len(nums); next++ {
        if nums[cur] != nums[next] {  // Found new unique element
            cur++
            nums[cur] = nums[next]  // Move unique element forward
        }
        // If duplicate, just continue (skip it)
    }
    
    return cur + 1
}
```

## Problem: Two Sum (Sorted Array)
**Current + Next pattern:**

```go
func twoSum(numbers []int, target int) []int {
    cur := 0
    next := len(numbers) - 1
    
    for cur < next {
        current := numbers[cur]
        nextNum := numbers[next]
        sum := current + nextNum
        
        if sum == target {
            return []int{cur + 1, next + 1}  // 1-indexed
        } else if sum < target {
            cur++  // Need bigger sum
        } else {
            next--  // Need smaller sum
        }
    }
    
    return nil
}
```

## Problem: Container With Most Water
**Current + Next pattern:**

```go
func maxArea(height []int) int {
    cur := 0
    next := len(height) - 1
    maxWater := 0
    
    for cur < next {
        currentHeight := height[cur]
        nextHeight := height[next]
        
        // Calculate area
        width := next - cur
        minHeight := min(currentHeight, nextHeight)
        area := width * minHeight
        maxWater = max(maxWater, area)
        
        // Move the shorter line (greedy: taller line might help later)
        if currentHeight < nextHeight {
            cur++
        } else {
            next--
        }
    }
    
    return maxWater
}
```

## Key Decision Points in Current + Next Pattern:

### 1. **Which pointer to move?**
- **Condition met**: Move based on strategy (keep better element, both elements, etc.)
- **No condition**: Usually `cur = next` (advance to next position)

### 2. **When sorting is needed:**
- **Interval problems**: Sort by start (merging) or end (non-overlapping)
- **Two pointer problems**: Often pre-sorted or sort by key attribute

### 3. **What to track:**
- **Count problems**: Increment counter when condition met
- **Collection problems**: Add to result array when condition met
- **Optimization problems**: Track best value seen so far

## Mental Model Summary:
1. **Sort** data if needed for the strategy to work
2. **Compare** `current` vs `next` 
3. **Decide** what to do based on comparison
4. **Advance** pointers based on strategy (who to keep/move)
5. **Track** the result (count, array, max value, etc.)

This pattern works for most two-pointer and interval problems. The key is deciding your advancement strategy based on the problem's optimization goal.