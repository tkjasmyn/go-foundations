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

## sort.go

    File: sort/sort.go
    Date: 2026-07-15
    Time Spent: 1 hr

    What I looked at:
    - The `Interface` interface ("Contract" for sorting any data)
    - `Sort` function ("the orchestarator". Calls `Interface`)
    - The `IntSlice` type (behaves like []int but is a distinct type with its own methods like func (p IntSlice) Len() int   { return len(p) })

    Key interface methods:
    - Len() int — number of elements in the collection or list
    - Less(i, j int) bool — "the brain", to know if item at position i is smaller than item at position j
    - Swap(i, j int) — swap the elementswith indexes i and j

    What I learned:
    sort.Interface defines the three questions any sorting algorithm needs answered, and Sort() orchestrates a clever hybrid algorithm (Pattern-Defeating Quicksort (pdqsort)) that asks those questions to rearrange anything into order.
    sort.Sort never sees the actual data, It only calls Len, Less, and Swap.

    What I don't understand yet:
    - How does Pattern-Defeating Quicksort (pdqsort) work and what does it do?
    - What is the significance and importance of time complexity?

    Questions to answer later:
    - Read about pdqsort
    - Read about time complexity

## list.go

    File: container/list/list.go
    Date: 2026-07-16
    Time Spent: 1 hr

    What I looked at:
    - The `Element` struct (element of a linked list)
    - The `List` struct (doubly linked list)

    Key struct fields (Element):
    - `next, prev *Element` — pointers to other `Element` nodes (the element that comes before and after this one)
    - `list *List` — back-pointer to the `List` container that owns this element
    - `Value any` — actual data or the value stored with this element (`any` means any data type can be put here)

    Key struct fields (List):
    - `root Element` — a sentinel (a dummy, fake, placeholder element that lives inside the List struct itself)
    - `len int` — how many real elements are in the list, excluding (this) sentinel element

    What I learned:
    Go's implementation uses a clever trick where the list is secretly a ring (circle), not a straight line. This makes the code much simpler because with such implementation there is never a `nil` pointer and edge cases are avoided or solved. Go uses a dummy node called `root` that lives inside the `List` struct itself/ The list is arranged in a circle where `root.next` points to the first real element `(Front())`, `root.prev` points to the last real element `(Back())`, the first real element's prev points back to `root`, the last real element's next points forward to `root`. There are no `nil` pointers in the chain. Every pointer always points to something valid. This eliminates entire categories of bugs.
    Each element is made to know which list it belongs to in order to prevent dangerous operations like moving an element from one list to another without proper cleanup, or removing an element from a list it doesn't belong to.
    `List` is a container that owns a circular chain of elements, using a built-in dummy anchor (`root`) so that even an empty list is a valid, self-contained ring with no loose ends or nil pointers.

    What I don't understand yet:
    - How does `root` really work and connect with everything?
    - The zero-value magic?

    Questions to answer later:
    - Read about `root`
    - Read about `zero-value`
