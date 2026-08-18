## waitgroup.go

    File: sync/waitgroup.go
    Date: 2026-18-08
    Time: 30 mins

    What I looked at:
    - The `WaitGroup` struct
    - The `Add` method
    - The `Done` method
    - The `Wait` method

    Key struct fields:
    - `noCopy noCopy` — A safety tag to prevent copying
    - `state atomic.Uint64
    - `sema uint32`

    What I learned:
    A waitgroup internally holds a tracking counter (atomic.Uint64). `Add` adds `delta`(int) to the counter. `Done` decrements the counter by one(Add(-1)). `Wait` blocks until the counter is 0. `Done` is a separate method instead of a mually calling `Add(-1)` yourself because of readability andit matches the mental model of the pattern
