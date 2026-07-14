## strings.go

    File: strings/strings.go
    Date: 2026-07-13
    Time spent: 30 min

    What I looked at:
    - The `Builder` struct
    - The `copycheck` method

    Key struct fields:
    - addr *Builder — holds the address of the builder? (used by `copycheck()` to detect if the builder was copied)
    - buf []byte — buffer holding the slice of the value? (the byte slice that grows as we write to the builder)

    What I learned:
    strings.Builder uses a slice of bytes and grows it
    similar to how slices work, avoiding repeated allocations.
    copycheck() method acts as a guard to prevent copying a non-zero builder and corrupting its address.

    What I don't understand yet:
    - How does `noescape` work for keeping builders stack-allocated when possible?

    Questions to answer later:
    - Read about noescape

## print.go

    File: fmt/print.go
    Date: 2026-07-14
    Time spent: 30 min

    What I looked at:
    - The `pp` struct (printer state)
    - `Print` function calls `Fprint(os.Stdout, a...)`
    - `Sprintf` creates a `pp`, formats into buffer, returns string

    Key struct fields:
    - `buf buffer` — holds the formatted output
    - `arg any` — the current item being formatted
    - `value reflect.Value` — for reflection-based formatting
    - `fmt fmt` — the actual formatter

    What I learned:
    fmt.Print is not simple. It delegates to Fprint, which creates a pp struct,
    formats into a buffer, and writes to stdout. The "simple" function hides
    a state machine.

    What I don't understand yet:
    - How does `sync.Pool` work for reusing pp structs?
    - What is `reflect.Value` exactly?

    Questions to answer later:
    - Read about sync.Pool
    - Read about reflect package basics
