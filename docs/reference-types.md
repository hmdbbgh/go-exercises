# Reference Types in Go — Internals and Behavior

Go has four built-in reference types:
- slices
- maps
- channels
- functions

Copying these values copies a small header, not the underlying data structure. This makes sharing cheap but also introduces aliasing.
