package utils

import (
	"reflect"
	"testing"
)

func TestSplitLast(t *testing.T) {

	tests := []struct {
		name string
		s    string
		c    rune
		n    int
		want []string
	}{
		{
			s:    "test.prefix",
			c:    '.',
			n:    2,
			want: []string{"test", "prefix"},
		},
		{
			s:    "test.prefix.user",
			c:    '.',
			n:    2,
			want: []string{"test.prefix", "user"},
		},
		{
			s:    "test",
			c:    '.',
			n:    2,
			want: []string{"test"},
		},
		{
			s:    "test.abc.def.effe.db",
			c:    '.',
			n:    2,
			want: []string{"test.abc.def.effe", "db"},
		},
		{
			s:    "test.abc.def.effe.db",
			c:    '.',
			n:    3,
			want: []string{"test.abc.def", "effe", "db"},
		},
		{
			s:    "test.abc.def.effe.db",
			c:    '.',
			n:    4,
			want: []string{"test.abc", "def", "effe", "db"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitLast(tt.s, tt.c, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitLast() = %v, want %v", got, tt.want)
			}
		})
	}
}
