## Week 4 Review

### What Was Planned vs What Happened

**Planned:** File Copier with Progress Bar + goroutines/channels  
**Actual:** Focus shifted to algorithm consolidation. File Copier works (basic copy + progress bar). Goroutines deferred to when actually needed.

### Algorithms This Week

| Algorithm                                      | Status             | How It Was Solved                                                                  |
| ---------------------------------------------- | ------------------ | ---------------------------------------------------------------------------------- |
| Valid Palindrome II                            | ✅                 | Two pointers + helper function for plain palindrome check. Skip one char allowed.  |
| Container With Most Water                      | ✅                 | Two pointers from ends. Move pointer at shorter line. Track max area.              |
| Longest Substring Without Repeating Characters | ✅ (with struggle) | Sliding window with map[byte]bool. Expand right, shrink left when duplicate found. |
| Product of Array Except Self                   | ✅                 | Two passes: left products forward, right products backward, multiply.              |
| Review: Two Sum II                             | ✅ FROM MEMORY     | Wrote correctly without looking at old code. Pattern is sticking.                  |

### Key Pattern: Two Pointers

Every algorithm this week used two pointers in some form:

- **Valid Palindrome II:** left/right from ends, branch on mismatch
- **Container With Most Water:** left/right from ends, move shorter
- **Longest Substring:** left/right as sliding window, shrink on duplicate
- **Product Except Self:** Two passes (conceptually two pointers)
- **Two Sum II:** left/right from ends on sorted array

### What I Struggled With

1. **Loop structure:** Nested vs sequential loops. Keeping braces correct.
2. **Variable scope:** Declaring products inside vs outside loops.
3. **Slice allocation:** `make([]int, len)` vs `[]int{}`
4. **Feeling like I don't understand:** Even when code passes tests, I expect to "feel" confident. I don't yet. That is normal.

### What Actually Worked

- **Two Sum II from memory:** Proves that repetition works. I did not invent the pattern, but I remembered it after 2 weeks.
- **Product Except Self:** I fixed 4 bugs myself (allocation, loop order, variable scope, missing multiplication). That is debugging skill.
- **File Copier:** Built a progress bar using `Read`/`Write` chunks. Learned `io.EOF` handling.

### What I Deferred

- Goroutines + channels for File Copier: Not needed yet. Will learn when building concurrent systems.
- Source code reading: Read `time/time.go` partially. Not a priority.

### One Pattern I See

&gt; "Two pointers" is not one technique. It is a family:
&gt; - Both from ends, move conditionally (Container With Most Water, Two Sum II)
&gt; - Both from start, window expands/contracts (Longest Substring)
&gt; - One forward, one backward, separate passes (Product Except Self)

### Next Week (Week 5)

**Project:** Key-Value Store HTTP API  
**Focus:** More algorithm reviews from memory. Build confidence through repetition, not novelty.

### Honest Note

I spent a lot of this week feeling like I was not improving because I needed hints. But on Friday I wrote Two Sum II from memory without looking at anything. That is proof the patterns are sticking. The goal is not to solve hard problems instantly. The goal is to remember patterns after seeing them multiple times.
