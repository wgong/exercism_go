package clock

import "fmt"

// Clock type
type Clock struct {
	h, m int
}

// New is Clock constructor
func New(hour, minute int) Clock {
	h, m := reset(hour, minute)
	return Clock{h, m}
}

// String returns string representation of clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

// Add increases offset by M mins
func (c Clock) Add(minutes int) Clock {
	p := &c
	p.h, p.m = reset(c.h, c.m+minutes)
	return *p
}

// Subtract decreases offset by M mins
func (c Clock) Subtract(minutes int) Clock {
	p := &c
	p.h, p.m = reset(c.h, c.m-minutes)
	return *p
}

func reset(h, m int) (int, int) {
	// renormalize to within one day
	mins := (h*60 + m) % 1440
	if mins < 0 {
		mins += 1440
	}
	return int(mins / 60), mins % 60
}

/*
// New is Clock constructor
func New(hour, minute int) *Clock {
	h, m := reset(hour, minute)
	return &Clock{h, m}
}

// String returns string representation of clock
func (c *Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

// Add increases offset by M mins
func (c *Clock) Add(minutes int) Clock {
	c.h, c.m = reset(c.h, c.m+minutes)
	return *c
}

// Subtract decreases offset by M mins
func (c *Clock) Subtract(minutes int) Clock {
	c.h, c.m = reset(c.h, c.m-minutes)
	return *c
}

*/
