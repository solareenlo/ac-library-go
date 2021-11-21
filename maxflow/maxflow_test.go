package maxflow

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxFrow_Simple(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	g := NewMfGraph(4)
	assert.Equal(0, g.AddEdge(0, 1, 1))
	assert.Equal(1, g.AddEdge(0, 2, 1))
	assert.Equal(2, g.AddEdge(1, 3, 1))
	assert.Equal(3, g.AddEdge(2, 3, 1))
	assert.Equal(4, g.AddEdge(1, 2, 1))
	assert.Equal(2, g.Flow(0, 3))

	var edgeEqual func(edge, edge)
	t.Helper()
	edgeEqual = func(expect, actual edge) {
		assert.Equal(expect.from, actual.from)
		assert.Equal(expect.to, actual.to)
		assert.Equal(expect.cap, actual.cap)
		assert.Equal(expect.flow, actual.flow)
	}

	var e edge
	e = edge{0, 1, 1, 1}
	edgeEqual(e, g.GetEdge(0))
	e = edge{0, 2, 1, 1}
	edgeEqual(e, g.GetEdge(1))
	e = edge{1, 3, 1, 1}
	edgeEqual(e, g.GetEdge(2))
	e = edge{2, 3, 1, 1}
	edgeEqual(e, g.GetEdge(3))
	e = edge{1, 2, 1, 0}
	edgeEqual(e, g.GetEdge(4))

	assert.Equal([]bool{true, false, false, false}, g.MinCut(0))
}

func TestMaxFrow_NotSimple(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	g := NewMfGraph(2)
	assert.Equal(0, g.AddEdge(0, 1, 1))
	assert.Equal(1, g.AddEdge(0, 1, 2))
	assert.Equal(2, g.AddEdge(0, 1, 3))
	assert.Equal(3, g.AddEdge(0, 1, 4))
	assert.Equal(4, g.AddEdge(0, 1, 5))
	assert.Equal(5, g.AddEdge(0, 0, 6))
	assert.Equal(6, g.AddEdge(1, 1, 7))
	assert.Equal(15, g.Flow(0, 1))

	var edgeEqual func(edge, edge)
	t.Helper()
	edgeEqual = func(expect, actual edge) {
		assert.Equal(expect.from, actual.from)
		assert.Equal(expect.to, actual.to)
		assert.Equal(expect.cap, actual.cap)
		assert.Equal(expect.flow, actual.flow)
	}

	var e edge
	e = edge{0, 1, 1, 1}
	edgeEqual(e, g.GetEdge(0))
	e = edge{0, 1, 2, 2}
	edgeEqual(e, g.GetEdge(1))
	e = edge{0, 1, 3, 3}
	edgeEqual(e, g.GetEdge(2))
	e = edge{0, 1, 4, 4}
	edgeEqual(e, g.GetEdge(3))
	e = edge{0, 1, 5, 5}
	edgeEqual(e, g.GetEdge(4))

	assert.Equal([]bool{true, false}, g.MinCut(0))
}

func TestMaxFrow_Cut(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	g := NewMfGraph(3)
	assert.Equal(0, g.AddEdge(0, 1, 2))
	assert.Equal(1, g.AddEdge(1, 2, 1))
	assert.Equal(1, g.Flow(0, 2))

	var edgeEqual func(edge, edge)
	t.Helper()
	edgeEqual = func(expect, actual edge) {
		assert.Equal(expect.from, actual.from)
		assert.Equal(expect.to, actual.to)
		assert.Equal(expect.cap, actual.cap)
		assert.Equal(expect.flow, actual.flow)
	}

	var e edge
	e = edge{0, 1, 2, 1}
	edgeEqual(e, g.GetEdge(0))
	e = edge{1, 2, 1, 1}
	edgeEqual(e, g.GetEdge(1))

	assert.Equal([]bool{true, true, false}, g.MinCut(0))
}

func TestMaxFrow_Twice(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	g := NewMfGraph(3)
	assert.Equal(0, g.AddEdge(0, 1, 1))
	assert.Equal(1, g.AddEdge(0, 2, 1))
	assert.Equal(2, g.AddEdge(1, 2, 1))
	assert.Equal(2, g.Flow(0, 2))

	var edgeEqual func(edge, edge)
	t.Helper()
	edgeEqual = func(expect, actual edge) {
		assert.Equal(expect.from, actual.from)
		assert.Equal(expect.to, actual.to)
		assert.Equal(expect.cap, actual.cap)
		assert.Equal(expect.flow, actual.flow)
	}

	var e edge
	e = edge{0, 1, 1, 1}
	edgeEqual(e, g.GetEdge(0))
	e = edge{0, 2, 1, 1}
	edgeEqual(e, g.GetEdge(1))
	e = edge{1, 2, 1, 1}
	edgeEqual(e, g.GetEdge(2))

	g.ChangeEdge(0, 100, 10)
	e = edge{0, 1, 100, 10}
	edgeEqual(e, g.GetEdge(0))

	assert.Equal(0, g.Flow(0, 2))
	assert.Equal(90, g.Flow(0, 1))

	e = edge{0, 1, 100, 100}
	edgeEqual(e, g.GetEdge(0))
	e = edge{0, 2, 1, 1}
	edgeEqual(e, g.GetEdge(1))
	e = edge{1, 2, 1, 1}
	edgeEqual(e, g.GetEdge(2))

	assert.Equal(2, g.Flow(2, 0))

	e = edge{0, 1, 100, 99}
	edgeEqual(e, g.GetEdge(0))
	e = edge{0, 2, 1, 0}
	edgeEqual(e, g.GetEdge(1))
	e = edge{1, 2, 1, 0}
	edgeEqual(e, g.GetEdge(2))
}

func TestMaxFrow_Bound(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	INF := int(^uint(0) >> 1)
	g := NewMfGraph(3)
	assert.Equal(0, g.AddEdge(0, 1, INF))
	assert.Equal(1, g.AddEdge(1, 0, INF))
	assert.Equal(2, g.AddEdge(0, 2, INF))
	assert.Equal(INF, g.Flow(0, 2))

	var edgeEqual func(edge, edge)
	t.Helper()
	edgeEqual = func(expect, actual edge) {
		assert.Equal(expect.from, actual.from)
		assert.Equal(expect.to, actual.to)
		assert.Equal(expect.cap, actual.cap)
		assert.Equal(expect.flow, actual.flow)
	}

	var e edge
	e = edge{0, 1, INF, 0}
	edgeEqual(e, g.GetEdge(0))
	e = edge{1, 0, INF, 0}
	edgeEqual(e, g.GetEdge(1))
	e = edge{0, 2, INF, INF}
	edgeEqual(e, g.GetEdge(2))
}

func TestMaxFrow_SelfLoop(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	g := NewMfGraph(3)
	assert.Equal(0, g.AddEdge(0, 0, 100))

	var edgeEqual func(edge, edge)
	t.Helper()
	edgeEqual = func(expect, actual edge) {
		assert.Equal(expect.from, actual.from)
		assert.Equal(expect.to, actual.to)
		assert.Equal(expect.cap, actual.cap)
		assert.Equal(expect.flow, actual.flow)
	}

	var e edge
	e = edge{0, 0, 100, 0}
	edgeEqual(e, g.GetEdge(0))
}

func TestMaxFrow_Stress(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	for phase := 0; phase < 10000; phase++ {
		n := rand.Intn(20) + 2
		m := rand.Intn(100) + 1
		s, tt := rand.Intn(n-1), rand.Intn(n-1)
		if s == tt {
			tt += 1
		}

		g := NewMfGraph(n)
		for i := 0; i < m; i++ {
			u := rand.Intn(n - 1)
			v := rand.Intn(n - 1)
			c := rand.Intn(10000)
			g.AddEdge(u, v, c)
		}
		flow := g.Flow(s, tt)
		dual := 0
		cut := g.MinCut(s)
		v_flow := make([]int, n)
		for _, e := range g.Edges() {
			v_flow[e.from] -= e.flow
			v_flow[e.to] += e.flow
			if cut[e.from] && !cut[e.to] {
				dual += e.cap
			}
		}
		assert.Equal(flow, dual)
		assert.Equal(-flow, v_flow[s])
		assert.Equal(flow, v_flow[tt])
		for i := 0; i < n; i++ {
			if i == s || i == tt {
				continue
			}
			assert.Equal(0, v_flow[i])
		}
	}
}
