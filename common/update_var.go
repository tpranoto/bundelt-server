package common

import "time"

type (
	UpdateInt64 struct {
		Value  int64
		Update bool
	}

	UpdateInt struct {
		Value  int64
		Update bool
	}

	UpdateFloat64 struct {
		Value  float64
		Update bool
	}

	UpdateString struct {
		Value  string
		Update bool
	}

	UpdateBool struct {
		Value  bool
		Update bool
	}

	UpdateTime struct {
		Value  time.Time
		Update bool
	}
)
