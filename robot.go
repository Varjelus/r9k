package r9k

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const DB_DRIVER = "sqlite3" // Only support this for now

type robot struct {
	db  *sql.DB
	cfg Config
}

func (r *robot) Process(uid, data string) (penalize bool, penalty penalty) {
	s := r.cfg.Normalizer(data)
	l := len(data)

	// If no content
	if l == 0 {
		goto fail
	}

	// If enough content but too much noise
	if l > r.cfg.MinContent && float32(len(s))/float32(l) < r.cfg.SignalRatio {
		goto fail
	}

	// If content is original
	if r.Original(s) {
		return
	}

fail:
	penalize = true
	penalty = r.Penalty(uid)
	return
}

func (r *robot) Original(s string) bool {
	// TODO:
	// if original, add to db and return true
	// else return false
	return false
}

func (r *robot) Penalty(uid string) penalty {
	// TODO:
	// if no previous penalty:
	//      add r.cfg.InitialPenalty for uid in db, calc, return penalty obj
	// else:
	//      run uid penalty through r.cfg.PenaltyFunction, update db, calc, return penalty obj
	return penalty{}
}

func New(cfg Config) (r *robot, err error) {
	r.db, err = sql.Open(DB_DRIVER, cfg.DatabaseURL)
	r.cfg = cfg
	return
}
