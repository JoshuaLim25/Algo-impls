Here are the essential Go string functions that'll cover 80% of your LeetCode string problems:

**Core manipulation:**
- `strings.Split(s, sep)` - Split by delimiter
- `strings.Join(slice, sep)` - Join slice with separator
- `strings.Replace(s, old, new, n)` / `strings.ReplaceAll(s, old, new)`
- `strings.ToLower(s)` / `strings.ToUpper(s)`

**Search & check:**
- `strings.Contains(s, substr)` - Substring check
- `strings.HasPrefix(s, prefix)` / `strings.HasSuffix(s, suffix)`
- `strings.Index(s, substr)` - First occurrence index (-1 if not found)
- `strings.Count(s, substr)` - Count occurrences

**Trimming:**
- `strings.TrimSpace(s)` - Remove leading/trailing whitespace
- `strings.Trim(s, cutset)` - Remove characters from both ends

**Builder for efficient concatenation:**
```go
var sb strings.Builder
sb.WriteString("hello")
sb.WriteByte('!')
result := sb.String()
```

**Rune conversion (critical for Unicode):**
```go
runes := []rune(s)  // Convert to runes for proper indexing
s = string(runes)   // Convert back
```

**Pro tips:**
- Use `strings.Builder` instead of `+=` for multiple concatenations
- Convert to `[]rune` when dealing with character manipulation (palindromes, reversals)
- `strings.Fields(s)` splits on any whitespace - cleaner than `Split(" ")`
- Remember Go strings are UTF-8, so `len(s)` gives bytes, not characters

These cover anagrams, palindromes, parsing, sliding window problems, and most string manipulation patterns you'll encounter.
