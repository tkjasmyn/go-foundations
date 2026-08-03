## Pattern: Two Pointers with Skip

**Problem:** Valid Palindrome II — check if string is palindrome after deleting at most one char.

**Approach:**

- Two pointers from both ends
- When chars match: move both inward
- When chars mismatch: branch into two checks
  - Check substring without left char
  - Check substring without right char
- Use helper function for plain palindrome check (no skips allowed)

**Helper signature:**

```go
func isPalindrome(s string, left, right int) bool
```
