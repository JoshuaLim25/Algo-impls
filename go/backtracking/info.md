**Backtracking is glorified trial and error with undo.**

## **The Core Concept:**
```
Try something → If it works, keep going → If it fails, UNDO and try next option
```

## **The Universal Template:**
```go
func backtrack(state) {
    // Base case: found solution or dead end
    if isComplete(state) {
        recordSolution(state)
        return
    }
    
    // Try each possible choice
    for choice := range getAllChoices(state) {
        makeChoice(choice)        // DO
        backtrack(newState)       // RECURSE  
        undoChoice(choice)        // UNDO
    }
}
```

## **Concrete Example - Generate All Subsets:**

**Problem:** Given `[1,2,3]`, generate all subsets.

**Think:** At each number, I have 2 choices: "include it" or "skip it"

```go
func subsets(nums []int) [][]int {
    var result [][]int
    var current []int  // This is our "state"
    
    var backtrack func(index int)
    backtrack = func(index int) {
        // BASE CASE: processed all numbers
        if index == len(nums) {
            result = append(result, append([]int{}, current...))
            return
        }
        
        // CHOICE 1: Include nums[index]
        current = append(current, nums[index])  // DO
        backtrack(index + 1)                    // RECURSE
        current = current[:len(current)-1]      // UNDO
        
        // CHOICE 2: Skip nums[index]  
        backtrack(index + 1)                    // Just recurse, no DO/UNDO needed
    }
    
    backtrack(0)
    return result
}
```

## **The Execution Tree for [1,2,3]:**
```
                    []
                  /    \
            include 1   skip 1
               [1]        []
              /   \      /   \
        include 2  skip 2   ...
           [1,2]    [1]
           /   \    /   \
          ...  ... ... ...
```

## **Why the UNDO step?**
Without it, `current` keeps growing and never resets:
```go
// WITHOUT UNDO:
current = [1]
current = [1,2]  
current = [1,2,3] // Gets recorded
// Now when we try "skip 3", current is still [1,2,3] - WRONG!
```

**The undo restores state for the next choice.**

## **Key Insight:**
Backtracking is just **DFS with cleanup**. You explore one path completely, then backtrack to try other paths.

The magic is in the **state management** - making sure each recursive call sees the correct state.
