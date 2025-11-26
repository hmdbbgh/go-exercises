# Clock – Notes

- Internal representation uses minutes-since-midnight for simplicity.
- Uses modular normalization: ((x % m) + m) % m to handle negative and overflow values.
- Add returns a new Clock, keeping the type effectively immutable.
