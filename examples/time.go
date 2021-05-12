package examples

import (
	"fmt"
	"time"
)

func Time() {
	p := fmt.Println

	now := time.Now()
	p(now)
	then := time.Date(2000, 11, 17, 20, 34, 58, 12, time.UTC)
	p(then)
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff).Equal(now))
	p(then.Add(-diff).Before(now))
}
