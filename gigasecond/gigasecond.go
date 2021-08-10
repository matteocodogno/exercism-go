package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
	const gigSec = 1e9
	t = t.Add(time.Second * gigSec)

	return t
}
