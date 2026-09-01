## sleep.go

    File: time/sleep.go
    Date: 2026-09-02
    Time: 15 mins

    What I looked at:
    - The `Sleep` function
    - The `runtime` timer integration
    - How `Sleep` differs from a busy loop

    Key function signature:
    - `func Sleep(d Duration)`

    What I learned:
    `Sleep` does not burn CPU. It registers a timer with the Go runtime, parks the current goroutine (takes it off the CPU), and resumes it when the timer fires. A million sleeping goroutines use the same CPU as one — the runtime only tracks the next timer that needs to fire. `Sleep` is implemented as `<-NewTimer(d).C` internally — it creates a one-shot timer and blocks on its channel.

## tick.go

    File: time/tick.go
    Date: 2026-09-02
    Time: 15 mins

    What I looked at:
    - The `Ticker` struct
    - `NewTicker(d Duration)`
    - The `Stop` method
    - The `C` channel field

    Key struct fields:
    - `C <-chan Time` — the channel that receives the current time on each tick
    - `initTicker` — runtime hook that starts the repeating timer

    What I learned:
    A `Ticker` is a `Timer` that repeats. It sends the current `time.Time` on channel `C` every `d` duration. The critical difference from `Sleep`: a Ticker must be stopped with `Stop()` or it leaks a goroutine — the runtime keeps a background thread alive to fire it forever. `Sleep` is one-shot and self-cleaning. `Ticker` is repeating and manual-cleanup. This is why Thursday's recurring tasks need `ticker.Stop()` on exit.
