package out

import "time"

func since(x time.Duration) time.Time {
	return time.Since(x)
}
