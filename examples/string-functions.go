package examples

import (
	"fmt"
	st "strings"
)

var p = fmt.Println

func StringFunctions() {
	p("Contains", st.Contains("test", "es"))
	p("Count", st.Count("test", "t"))
	p("HasPrefix", st.HasPrefix("test", "te"))
	p("HasSuffix", st.HasSuffix("test", "st"))
	p("Index", st.Index("test", "e"))
	p("Join", st.Join([]string{"a", "b", "c"}, "-"))
	p("Repeat", st.Repeat("a", 3))
	p("Replace", st.Replace("test", "e", "E", -1))
	p("Split", st.Split("a-b-c-d", "-"))
	p("ToLower", st.ToLower("TEST"))
	p("ToUpper", st.ToLower("test"))
}
