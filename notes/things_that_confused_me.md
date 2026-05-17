# Things that confused me in Go

- **`:=` vs `var`:** why does both exist?
- **Unused variables:** the compiler yells at me for importing something or declaring a variable but not using it immediately. Highly annoying.
- **Slices vs Arrays:** arrays have fixed sizes, slices don't. But slices are backed by arrays.
- **Pointers:** `&` vs `*`. I keep typing `*` when I want `&`. Value receivers vs Pointer receivers on methods.
- **Goroutine Leaks:** if a goroutine blocks on reading a channel, but no one ever writes to it, it leaks! Found out the hard way.
