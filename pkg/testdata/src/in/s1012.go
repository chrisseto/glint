package in

import "time"

func since(x time.Time) time.Duration {
	return time.Now().Sub(x)
}
