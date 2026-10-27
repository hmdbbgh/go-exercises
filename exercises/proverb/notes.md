# Proverb — Notes

## Overview
Generate a proverb chain where each item depends on the next.

## Versions

### 1. Simple Version
Direct loop and append. Easy to read but less optimal.

### 2. Idiomatic Version
Uses append with preallocated capacity. Cleanest Go solution.

### 3. Optimized Version
Exact-size slice, no append, constant templates.
Ideal when memory behavior must be predictable.

## Complexity
Time: O(n)
Space: O(n)
