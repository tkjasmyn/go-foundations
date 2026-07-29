## time.go

    File: time/time.go, time/format.go
    Date: 2026-07-29
    Time:

    What I looked at:
    - The `Time` struct
    - The `Now() Time` function (gives a `Time` value representing right now)
    - The `Format(layout string)` method

    What I learned:
    `Now()` returns `Time` value representing now. `Format` turns time into a readable string. Go uses the specific date Mon Jan 2 15:04:05 MST 2006 as a reference to teach `Format` what each component means.

    What I don't understand yet:
    - How the monotonic clock reading in `ext` works (not urgent)

    Questions to answer later:
    - What other time layouts exist besides RFC3339?
