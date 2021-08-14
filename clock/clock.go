package clock

import "fmt"

const minutesInDay int = 1440
const hoursInDay int = 24
const minutesInHour int = 60

type Clock struct {
	h, m int
}

func convertToClock(minutes int) Clock {
	minutes = minutes % minutesInDay
	if minutes < 0 {
		minutes =  minutesInDay + minutes
	}

	return Clock{
		(minutes / minutesInHour) % hoursInDay,
		minutes % minutesInHour,
	}
}

func New(h, m int) Clock {
	return convertToClock(h * minutesInHour + m)
}

func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock.h, clock.m)
}

func (clock Clock) convertToMinutes() int {
	return clock.h * minutesInHour + clock.m
}

func (clock Clock) Add(min int) Clock {
	return convertToClock(clock.convertToMinutes() + min)
}

func (clock Clock) Subtract(min int) Clock {
	return convertToClock(clock.convertToMinutes() - min)
}
