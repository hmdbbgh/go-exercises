package clock

import "fmt"

const minutesPerDay = 24 * 60

// Clock represents a 24-hour clock as minutes since midnight.
type Clock struct {
    minutes int
}

// New constructs a new clock normalized to the 24-hour range.
func New(h, m int) Clock {
    total := h*60 + m
    total = ((total % minutesPerDay) + minutesPerDay) % minutesPerDay
    return Clock{minutes: total}
}

// Add returns a new Clock with m minutes added.
func (c Clock) Add(m int) Clock {
    return New(0, c.minutes+m)
}

// String formats the clock as HH:MM.
func (c Clock) String() string {
    h := c.minutes / 60
    m := c.minutes % 60
    return fmt.Sprintf("%02d:%02d", h, m)
}
