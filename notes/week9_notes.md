## net.go

    File: net/net.go
    Date: 2026-09-08
    Time: 30 mins

    What I looked at:
    - The `Conn` interface
    - The `Read` method
    - The `Write` method
    - The `Close` method

    Key interface methods:
    - `Read`
    - `Write`
    - `Close`
    - `LocalAddr`
    - 'RemoteAddr'
    - `SetDeadline`
    - `SetReadDeadline`
    - `SetWriteDeadline`

    What I learned:
    `Conn` is a generic stream-oriented network connection and multiple goroutines may simultaneously call methods on it.
    `Read` (from `conn` struct) implements the `Conn` `Read` method and same goes for other methods.
    `SetDeadline` methods set the deadines for future read and write calls. After the deadline is met it will timeout and stop operations

    What I don't understand yet:
    How do the methods on the struct work and implement the methods on the interface?
