package components

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type Search func(query string) Result

func FakeSearcher(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("[%s] result for %q\n", kind, query))
	}
}
