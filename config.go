package r9k

import (
	"regexp"
	"strings"
	"time"
)

type Config struct {
	DatabaseURL     string
	InitialPenalty  time.Duration
	PenaltyFunction func(time.Duration) time.Duration
	HalfLife        time.Duration
	Normalizer      func(string) string
	SignalRatio     float32
	MinContent      int
}

// Everything not a-z or Nordic characters.
// Nordic = [Sweden, Norway, Denmark, Finland] in this context.
var PunctuationMatcher = regexp.MustCompile(`/[^\wäæöøå]/gim`)
var DefaultConfig = Config{
	DatabaseURL:    "r9k.sqlite",
	InitialPenalty: time.Second * 2,
	PenaltyFunction: func(d time.Duration) time.Duration {
		return d * 4
	},
	HalfLife: time.Hour * 6,
	Normalizer: func(s string) string {
		return PunctuationMatcher.ReplaceAllString(strings.ToLower(s), "")
	},
	SignalRatio: 0.9,
	MinContent:  10,
}
