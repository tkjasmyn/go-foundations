## mutex.go

    File: sync/mutex.go
    Date: 2026-07-21
    Time spent: 15 mins

    What I looked at:
    - The `mutex` struct ("bathroom lock")
    - The `Lock` method (block until you get the lock)
    - The `Unlock` method (release the lock)

    Key struct fields:
    - _nocopy — A dummy field that tells Go's `vet` tool: "Don't let anyone copy this struct." Because copying a lock is dangerous
    - mu isync.Mutex — actual lock mechanism, imlemented in Go's runtime

    What I learned:
    A `Mutex` is a single key to a single room. Only one goroutine holds it at a time. Everyone else waits. The zero value is an unlocked, ready-to-use lock that must never be copied. `Lock()` either grabs the lock immediately, or your goroutine goes to sleep until the current holder releases it — whichever comes first.
    `Unlock()` is a memory flush. `Lock()` is a memory refresh. The pair creates a pipeline where everything written before the flush is guaranteed visible after the refresh.

    What I don't understand yet:
    - What exactly is `goroutine`?
    - What is Go's `vet` tool and its relevance?

    Questions to answer later:
    - Read about `goroutine`
    - Read about Go's `vet` tool
