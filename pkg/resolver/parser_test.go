package resolver_test

import (
	"go-pdirs/pkg/resolver"
	"go-pdirs/testutils"
	"testing"
)

type ParseProjectsTestCase struct {
	name, input string
	expected    resolver.Projects
}

func TestParseProjects(t *testing.T) {
	home := "/home/user"
	environ := testutils.NewMockEnviron(map[string]string{
		"HOME": home,
	})

	cases := []ParseProjectsTestCase{
		{
			name:  "Single pure",
			input: "foo=path/to/foo\n",
			expected: resolver.Projects{
				"foo": "path/to/foo",
			},
		},
		{
			name: "Multiple pure",
			input: "" +
				"foo=path/to/foo\n" +
				"bar=path/to/bar\n" +
				"baz=path/to/baz\n",
			expected: resolver.Projects{
				"foo": "path/to/foo",
				"bar": "path/to/bar",
				"baz": "path/to/baz",
			},
		},
		{
			name:  "Single with environment variable",
			input: "foo=$HOME/path/to/foo\n",
			expected: resolver.Projects{
				"foo": home + "/path/to/foo",
			},
		},
		{
			name: "Multiple with environment variable",
			input: "" +
				"foo=$HOME/path/to/foo\n" +
				"bar=$HOME/path/to/bar\n" +
				"baz=$HOME/path/to/baz\n",
			expected: resolver.Projects{
				"foo": home + "/path/to/foo",
				"bar": home + "/path/to/bar",
				"baz": home + "/path/to/baz",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			results, err := resolver.ParseProjects(c.input, environ)

			if err != nil {
				t.Error(err)
			}

			if len(results) != len(c.expected) {
				t.Errorf(
					"Expected %d projects, got %d\n",
					len(c.expected),
					len(results),
				)
			}

			for name, result := range results {
				if c.expected[name] != result {
					t.Errorf(
						"For project %s:\n  Expected: %s\n  Got: %s\n",
						name,
						c.expected[name],
						result,
					)
				}
			}
		})
	}
}
