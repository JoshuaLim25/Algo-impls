# Rune conversion
- `rune - '0'`: acts like `strconv.Atoi()` (i.e., get the actual num)
- `num + '0'`: acts like `strconv.Itoa()` (i.e., get the actual rune)
- Remember by: it's NOT "add" for `Atoi`
- **NOTE**: the x + '0' trick ONLY WORKS FOR SINGLE DIGITS (0-9)

## Hashmap alternative for tracking character frequency
- Initialize `freq := make([]int, 26)`  (26 letters in alphabet)
```go
for _, r := range s {
    freq[r-'a']++  // rune-'a' converts to appropriate index (e.g., a=0, b=1, etc)
}
```

# Integer manipulation
```go
// Reverse integer digits
for num > 0 {
    digit := num % 10
    reversed = reversed*10 + digit
    num /= 10
}

// Count digits
digits := int(math.Log10(float64(num))) + 1

// Extract kth digit from right (0-indexed)
kthDigit := (num / int(math.Pow(10, k))) % 10
```

# Bit manipulation essentials
```go
// Check if power of 2
isPowerOf2 := n > 0 && (n & (n-1)) == 0

// Count set bits
bits := bits.OnesCount(n)

// Toggle bit
n ^= (1 << k)  // Toggle kth bit

// XOR for finding unique element
result := 0
for _, num := range nums {
    result ^= num  // Duplicates cancel out
}
```

# Two pointers patterns
```go
// Remove duplicates in-place
j := 0
for i := 1; i < len(nums); i++ {
    if nums[i] != nums[j] {
        j++
        nums[j] = nums[i]
    }
}
```

# Sliding window
```go
left := 0
for right := 0; right < len(s); right++ {
    // Expand window
    for /* condition violated */ {
        // Shrink from left
        left++
    }
    // Update result
}
```

# Fast exponentiation
```go
func pow(base, exp int) int {
    result := 1
    for exp > 0 {
        if exp&1 == 1 {
            result *= base
        }
        base *= base
        exp >>= 1
    }
    return result
}
```
