# Go vs Python

just writing down some thoughts comparing the two.

## Things Go does better
- **Speed:** oh my god it's so fast. compilation is instantaneous and binaries run immediately.
- **Concurrency:** goroutines and channels are infinitely better than Python's `asyncio` or threading.
- **Single Binary:** compiling to a single file and running it anywhere without pip install or venv is pure magic.

## Things Python does better
- **List comprehensions:** missing these a lot. `[x for x in list if x > 1]` is way cleaner than Go's for loops.
- **Dynamic typing:** go's compiler is extremely picky. sometimes you just want to dump data into a map and read it without defining structs or doing `i.(string)`.
- **Standard library:** python's stdlib has things like `requests` and simple json parsing that don't need struct tagging.

## Summary
Go feels like a sports car but with manual transmission. Python is a cozy automatic sedan.
