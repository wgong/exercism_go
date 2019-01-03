package clock

import "fmt"

// Clock type
type Clock struct {
	h, m int
}

// New is Clock constructor
func New(hour, minute int) *Clock {
	h, m := reset(hour, minute)
	return &Clock{h, m}
}

// String returns string representation of clock
func (c *Clock) String() string {
	h, m := reset(c.h, c.m)
	return fmt.Sprintf("%02d:%02d", h, m)
}

// Add increases offset by M mins
func (c *Clock) Add(minutes int) Clock {
	c.m += minutes
	return *c
}

// Subtract decreases offset by M mins
func (c *Clock) Subtract(minutes int) Clock {
	c.m -= minutes
	return *c
}

func reset(h, m int) (int, int) {
	// renormalize to within one day
	mins := (h*60 + m) % 1440
	if mins < 0 {
		mins += 1440
	}
	return int(mins / 60), mins % 60
}
