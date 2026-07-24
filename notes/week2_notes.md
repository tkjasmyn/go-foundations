## mutex.go

    File: sync/mutex.go
    Date: 2026-07-21
    Time spent: 15 mins

    What I looked at:
    - The `mutex` struct ("bathroom lock")
    - The `Lock` method (block until you get the lock)
    - The `Unlock` method (release the lock)

    Key struct fields:
    - `_nocopy` — A dummy field that tells Go's `vet` tool: "Don't let anyone copy this struct." Because copying a lock is dangerous
    - `mu isync.Mutex` — actual lock mechanism, imlemented in Go's runtime

    What I learned:
    A `Mutex` is a lock. Only one goroutine holds it at a time. Everyone else waits. The zero value is an unlocked, ready-to-use lock that must never be copied. `Lock()` locks `m`, a pointer to the mutex. `Unlock()` unlocks `m`.

    What I don't understand yet:
    - What exactly is `goroutine`?
    - What is Go's `vet` tool and its relevance?

    Questions to answer later:
    - Read about `goroutine`
    - Read about Go's `vet` tool

## buffer.go

    File: bytes/buffer.go
    Date: 2026-07-24

    What I looked at:
    - The `Buffer` struct
    - The `Write` method (appends bytes to the buffer and increasing the size of the buffer as needed)
    - The `Read` method (reads the next bytes from the buffer or until the buffer is drained)
    - The `String` method (returns the contents of unread portion of the buffer as a string)

    Key struct fields:
    - `buf []byte` — a slice of bytes
    - `off int` —
    - `lastRead readOp` —

    What I learned:
    byte.Buffer is the computer memory storage of bytes with `Buffer.Read` and `Buffer.Write` methods. bytes.Buffer helps makes strings mutable.
    It also makes string concatenation faster and easier by just appending instead of copying

    What I don't understand yet:
    - What does `off int` represent?
    - What is `readOp`?

    Questions to answer later:
    - Read about `readOp`
