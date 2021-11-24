package segtree

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func op(a, b string) string {
	if a == "$" {
		return b
	}
	if b == "$" {
		return a
	}
	return a + b
}

func e() string {
	return "$"
}

func cmp(v string) bool {
	return v != "$"
}

func TestSegTree_Zero(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	s := []string{}
	seg := NewSegtree(s, op, e)
	assert.Equal("$", seg.AllProd())
}

func TestSegTree_Invalid(t *testing.T) {
	t.Parallel()

	s := make([]string, 10)
	seg := NewSegtree(s, op, e)

	assert.Panics(t, func() { seg.Get(-1) })
	assert.Panics(t, func() { seg.Get(10) })

	assert.Panics(t, func() { seg.Prod(-1, -1) })
	assert.Panics(t, func() { seg.Prod(3, 2) })
	assert.Panics(t, func() { seg.Prod(0, 11) })
	assert.Panics(t, func() { seg.Prod(-1, 11) })

	assert.Panics(t, func() { seg.MaxRight(11, cmp) })
	assert.Panics(t, func() { seg.MinLeft(-1, cmp) })
	assert.Panics(t, func() { seg.MaxRight(0, cmp) })
}

func TestSegTree_One(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	s := make([]string, 1)
	s[0] = e()
	seg := NewSegtree(s, op, e)
	assert.Equal("$", seg.AllProd())
	assert.Equal("$", seg.Get(0))
	assert.Equal("$", seg.Prod(0, 1))
	seg.Set(0, "dummy")
	assert.Equal("dummy", seg.Get(0))
	assert.Equal("$", seg.Prod(0, 0))
	assert.Equal("dummy", seg.Prod(0, 1))
	assert.Equal("$", seg.Prod(1, 1))
}

var y string

func leqY(x string) bool {
	return len(x) <= len(y)
}

func TestSegTree_CompareNaive(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	for n := 0; n < 30; n++ {
		s0 := make([]string, n)
		for i := range s0 {
			s0[i] = e()
		}
		s1 := make([]string, n)
		for i := range s1 {
			s1[i] = e()
		}
		seg0 := NewSegtree(s0, op, e)
		seg1 := NewSegtree(s1, op, e)
		for i := 0; i < n; i++ {
			s := ""
			s += "a" + strconv.Itoa(i)
			seg0.Set(i, s)
			seg1.Set(i, s)
		}

		for l := 0; l <= n; l++ {
			for r := l; r <= n; r++ {
				assert.Equal(seg0.Prod(l, r), seg1.Prod(l, r))
			}
		}

		for l := 0; l <= n; l++ {
			for r := l; r <= n; r++ {
				y = seg1.Prod(l, r)
				assert.Equal(seg0.MaxRight(l, leqY), seg1.MaxRight(l, leqY))
				assert.Equal(seg0.MaxRight(l, leqY), seg1.MaxRight(l, func(x string) bool {
					return len(x) <= len(y)
				}))
			}
		}

		for l := 0; l <= n; l++ {
			for r := l; r <= n; r++ {
				y = seg1.Prod(l, r)
				assert.Equal(seg0.MinLeft(l, leqY), seg1.MinLeft(l, leqY))
				assert.Equal(seg0.MinLeft(l, leqY), seg1.MinLeft(l, func(x string) bool {
					return len(x) <= len(y)
				}))
			}
		}
	}
}
