package r9k

import (
	"time"
)

type penalty struct {
	UID      string
	Issued   time.Time
	Duration time.Duration
	Expires  time.Time
}
