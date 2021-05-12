package examples

import (
	"fmt"
	"time"
)

func TimeFormatingParsing() {
	p := fmt.Println

	t := time.Now()

	p(t.Format(time.RFC3339))
	p(t.Format(time.RFC1123Z))

	t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")

	p(t1.Hour())

	p(t.Format("3:04PM"))
	p(t.Format("Mon Jan _2 15:04:05 2006"))
	p(t.Format("2006-01-01T15:04"))

	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	ansic := "Mon Jan _2 15:04:07 2006"
	_, e := time.Parse(ansic, "8:41PM")
	p(e)
}
