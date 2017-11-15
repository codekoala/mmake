package resolver_test

import (
	"testing"

	"github.com/tj/mmake/resolver"
)

func TestGetIncludePath(t *testing.T) {
	var cases = []struct {
		Args     []string
		Expected string
	}{
		{[]string{}, resolver.DefaultIncludePath},
		{[]string{"update", "-I", "./relative/path"}, "./relative/path"},
		{[]string{"update", "-I./other/path"}, "./other/path"},
		{[]string{"update", "-I./other/path", "-I", "multiple/"}, "./other/path"},
		{[]string{"-I", "multiple/"}, "multiple/"},
		{[]string{"-I"}, resolver.DefaultIncludePath},
		{[]string{"update"}, resolver.DefaultIncludePath},
	}

	for _, c := range cases {
		t.Run(c.Expected, func(t *testing.T) {
			out := resolver.GetIncludePath(c.Args)
			if out != c.Expected {
				t.Errorf("expected %q, got %q", c.Expected, out)
			}
		})
	}
}
