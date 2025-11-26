package clock

import "fmt"

// ClockHM is an alternative representation using hour and minute fields.
type ClockHM struct {
    hour, minute int
}

// NewHM constructs a ClockHM, normalizing internally.
func NewHM(h, m int) ClockHM {
    total := h*60 + m
    total = ((total % minutesPerDay) + minutesPerDay) % minutesPerDay
    return ClockHM{
        hour:   total / 60,
        minute: total % 60,
    }
}

// Add returns a new ClockHM with m minutes added.
func (c ClockHM) Add(m int) ClockHM {
    return NewHM(c.hour, c.minute+m)
}

// String formats ClockHM as HH:MM.
func (c ClockHM) String() string {
    return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}
