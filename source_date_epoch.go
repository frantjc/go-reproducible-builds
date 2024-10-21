package reproduciblebuilds

import (
	"os"
	"strconv"
	"time"
)

var SourceDateEpoch = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC).
	Add(func() time.Duration {
		sourceDateEpoch := os.Getenv("SOURCE_DATE_EPOCH")
		if sourceDateEpoch == "" {
			return 0
		}

		if seconds, err := strconv.Atoi(sourceDateEpoch); err == nil {
			return time.Second * time.Duration(seconds)
		}

		return 0
	}())
