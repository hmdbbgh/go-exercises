# Value vs Pointer Semantics in Go — Deep Dive

Go is a value-oriented language. By default, values are copied on assignment or when passed to functions. Pointers allow you to share and mutate state explicitly.

This document explains:
- When to use values vs pointers
- How this compares to Python
- How this affects APIs and method receivers
