package convolution

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvolution_Convolve(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	mod := 998244353
	c := NewConvolution(mod, 3)
	assert.Equal([]int{}, c.Convolve([]int{}, []int{}))
	assert.Equal([]int{}, c.Convolve([]int{}, []int{1, 2}))
	assert.Equal([]int{}, c.Convolve([]int{1, 2}, []int{}))
	assert.Equal([]int{}, c.Convolve([]int{1}, []int{}))

	n := 1234
	m := 2345
	a := make([]int, n)
	b := make([]int, m)
	for i := range a {
		a[i] = rand.Int() % mod
	}
	for i := range b {
		b[i] = rand.Int() % mod
	}
	assert.Equal(c.ConvolutionNaive(a, b), c.Convolve(a, b))

	mod = 998244353
	c = NewConvolution(mod, 3)
	for n := 1; n < 20; n++ {
		for m := 1; m < 20; m++ {
			a := make([]int, n)
			b := make([]int, m)
			for i := range a {
				a[i] = rand.Int() % mod
			}
			for i := range b {
				b[i] = rand.Int() % mod
			}
			assert.Equal(c.ConvolutionNaive(a, b), c.Convolve(a, b))
		}
	}

	mod = 924844033
	c = NewConvolution(mod, 3)
	for n := 1; n < 20; n++ {
		for m := 1; m < 20; m++ {
			a := make([]int, n)
			b := make([]int, m)
			for i := range a {
				a[i] = rand.Int() % mod
			}
			for i := range b {
				b[i] = rand.Int() % mod
			}
			assert.Equal(c.ConvolutionNaive(a, b), c.Convolve(a, b))
		}
	}

	mod = 2130706433
	c = NewConvolution(mod, 13)
	n = 1 << 5
	m = 1 << 6
	a = make([]int, n)
	b = make([]int, m)
	for i := range a {
		a[i] = rand.Intn(mod - 1)
	}
	for i := range b {
		b[i] = rand.Intn(mod - 1)
	}
	assert.Equal(c.ConvolutionNaive(a, b), c.Convolve(a, b))
}
