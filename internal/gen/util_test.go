package gen

import (
	"testing"
)

func TestFindWildcard(t *testing.T) {
	wildcard, i, valid := FindWildcard("/test/:name/:last_name/*wild")
	t.Log(wildcard, i, valid)
	path := "/test/:name/:last_name/*wild"
	params := GetPathParameters(path)
	t.Log(params)
}
