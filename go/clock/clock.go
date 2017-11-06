package clock

import "fmt"

const testVersion = 4

type Clock struct {
	h, m int
}

func New(hour, minute int) Clock {
	return Clock{hour, minute}.normalizeTime()
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}

func (c Clock) Add(minutes int) Clock {
	c.m += minutes
	return c.normalizeTime()
}

func (c Clock) normalizeTime() Clock {
	c.h = (c.h + c.m/60) % 24
	c.m = c.m % 60

	if c.m < 0 {
		c.h -= 1
		c.m += 60
	}
	if c.h < 0 {
		c.h += 24
	}

	return c
}
