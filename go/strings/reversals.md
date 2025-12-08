# Prepending/reversing
- *Note*: avoid `slices.Insert(repr, 0, r)` since it's O(n^2)
```go
// 1. reverse
// for _, r := range a {
// 	tmpA = append([]rune{r}, tmpA...)
// }
// 2. reverse
// for _, r := range a {
// 	slices.Insert(tmpA, 0, r)
// }
// for _, r := range a {
// 	slices.Insert(tmpB, 0, r)
// }
// 3. reverse
repr := []rune(a)
for i, j := 0, len(repr)-1; i < j; i, j = i+1, j-1 {
    repr[i], repr[j] = repr[j], repr[i]
}
```
- For strings, `strings.Builder` is the goto choice for concatenation within a loop: really, whenever you think of repeatedly doing `s += ...` in a loop, think of the builder for a more efficient impl.
```go
var sb strings.Builder
reprA := []rune(a)
tmp := []rune{}
for i := len(reprA)-1; i >= 0; i-- {
    sb.WriteRune(reprA[i])
}
for i := len(reprA)-1; i >= 0; i-- {
    tmp = append(tmp, reprA[i])
}
fmt.Printf("%+v\n", sb.String())
fmt.Printf("%s\n", string(tmp))
```

```go
// Core methods
var sb strings.Builder

// Writing
sb.WriteString("hello")     // Write string
sb.WriteByte('x')          // Write single byte
sb.WriteRune('€')          // Write Unicode rune
sb.Write([]byte("data"))   // Write byte slice

// Getting result
s := sb.String()           // Convert to string
len := sb.Len()           // Current length

// Optimization
sb.Grow(1000)             // Pre-allocate capacity
sb.Reset()                // Clear for reuse
```
