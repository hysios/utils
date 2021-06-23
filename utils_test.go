package utils

import (
	"reflect"
	"testing"
	"time"

	"github.com/tj/assert"
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

func TestParseDatestr(t *testing.T) {

	got, err := ParseDatestr("20210621140431974", 3)
	assert.NoError(t, err)
	assert.NotNil(t, got)
	t.Logf("got %s", got)
	t.Logf("date %s", time.Date(2021, 6, 21, 14, 04, 31, 974*1000000, time.UTC))
	assert.Equal(t, got, time.Date(2021, 6, 21, 14, 04, 31, 974*1000000, time.UTC))
}
