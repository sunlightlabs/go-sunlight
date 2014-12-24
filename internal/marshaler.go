package internal

import (
	"time"
)

type Time struct {
	time.Time
}

func (t *Time) UnmarshalJSON(in []byte) (err error) {
	out, err := time.Parse("\"2006-01-02 15:04:05\"", string(in))
	*t = Time{out}
	return
}
