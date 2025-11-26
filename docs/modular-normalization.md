# Modular Normalization — The Universal Formula

Explains why Go's `%` with negatives can be surprising and how to normalize values into a fixed range using:

    ((x % m) + m) % m

with applications to:
- clocks
- angles
- circular buffers
- hashing
