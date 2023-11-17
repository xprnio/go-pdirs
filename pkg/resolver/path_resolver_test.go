package resolver_test

import (
	"fmt"
	"go-pdirs/pkg/resolver"
	"go-pdirs/testutils"
	"testing"
)

type ResolvePathTestCase struct {
	name, input, expected string
}

func TestResolvePath(t *testing.T) {
	home := "/home/user"
	environ := testutils.NewMockEnviron(map[string]string{
		"HOME": home,
	})

	cases := []ResolvePathTestCase{
		{
			name:     "Pure",
			input:    "path/to/file",
			expected: "path/to/file",
		},
		{
			name:     "With environment variable",
			input:    "$HOME/path/to/file",
			expected: fmt.Sprintf("%s/path/to/file", home),
		},
	}

	for i, c := range cases {
		t.Run(
			fmt.Sprintf("Test case %d", i),
			func(t *testing.T) {
				result, err := resolver.ResolvePath(c.input, environ)
				if err != nil {
					t.Error(err)
				}

				if c.expected != result {
					t.Errorf(
						"With input %s\n  Expected: %s\n  Got: %s\n",
						c.input,
						c.expected,
						result,
					)
				}
			},
		)
	}
}
