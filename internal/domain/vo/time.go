package vo

import (
	"fmt"
	"strconv"
	"time"
)

func ParseTimeFromUnix(s string) (time.Time, error) {
	loc, _ := time.LoadLocation("Europe/Madrid")
	segs, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse time: %w", err)
	}
	t := time.Unix(segs, 0).In(loc)
	return t, nil
}
