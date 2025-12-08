**The Core Mental Model:**

Think of it as **decision trees** - at each step, you're making choices about what to include/exclude or how to arrange things.

## **The Big 3 Patterns:**

### **1. Combinations: "Choose K from N"**
```go
// Template: C(n,k) = C(n-1,k-1) + C(n-1,k)
// "Include current item" + "Skip current item"

func combine(n, k int) [][]int {
    var result [][]int
    var path []int
    
    var backtrack func(start int)
    backtrack = func(start int) {
        if len(path) == k {
            result = append(result, append([]int{}, path...))
            return
        }
        
        for i := start; i <= n; i++ {
            path = append(path, i)
            backtrack(i + 1)  // No repeats
            path = path[:len(path)-1]
        }
    }
    
    backtrack(1)
    return result
}
```

### **2. Permutations: "Arrange N items"**
```go
// Template: Try each unused item at current position

func permute(nums []int) [][]int {
    var result [][]int
    var path []int
    used := make([]bool, len(nums))
    
    var backtrack func()
    backtrack = func() {
        if len(path) == len(nums) {
            result = append(result, append([]int{}, path...))
            return
        }
        
        for i := 0; i < len(nums); i++ {
            if used[i] { continue }
            
            path = append(path, nums[i])
            used[i] = true
            backtrack()
            used[i] = false
            path = path[:len(path)-1]
        }
    }
    
    backtrack()
    return result
}
```

### **3. Subsets: "All possible selections"**
```go
// Template: For each item, decide include/exclude

func subsets(nums []int) [][]int {
    var result [][]int
    var path []int
    
    var backtrack func(start int)
    backtrack = func(start int) {
        result = append(result, append([]int{}, path...))
        
        for i := start; i < len(nums); i++ {
            path = append(path, nums[i])
            backtrack(i + 1)
            path = path[:len(path)-1]
        }
    }
    
    backtrack(0)
    return result
}
```

## **The Decision Framework:**

**Ask yourself:**
1. **Order matters?** → Permutation
2. **Pick K items?** → Combination  
3. **All possible groups?** → Subsets

**The backtracking recipe:**
1. **Add** current choice to path
2. **Recurse** with updated constraints
3. **Remove** choice (backtrack)

This covers 90% of contest problems. The hard part isn't the math - it's recognizing which pattern fits.
