package simplifypath

import (
	"strings"
	"testing"
)

type Tables struct {
	path    string
	newPath string
}

func TestSimplifyPath(t *testing.T) {
	tables := []Tables{
		{"/home/", "/home"},
		{"/home/aland/", "/home/aland"},
		{"/../", "/"},
		{"/aa/./bb/../../c", "/c"},
		{"/a//b////c/d//././/..", "/a/b/c"},
	}

	var str string
	for _, v := range tables {
		if str = simplifyPath(v.path); strings.Compare(str, v.newPath) != 0 {
			t.Errorf("Get:%s, expect:%s, path:%s", str, v.newPath, v.path)
		}
	}
}
