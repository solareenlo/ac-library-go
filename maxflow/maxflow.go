package maxflow

type preEdge struct{ to, rev, cap int }
type pair struct{ x, y int }
type MfGraph struct {
	n   int
	pos []pair
	g   [][]preEdge
}

func NewMfGraph(n int) *MfGraph {
	pos := make([]pair, 0)
	g := make([][]preEdge, n)
	return &MfGraph{n, pos, g}
}

func (mfg *MfGraph) AddEdge(from, to int, cap int) int {
	m := len(mfg.pos)
	formId := len(mfg.g[from])
	toId := len(mfg.g[to])
	mfg.pos = append(mfg.pos, pair{from, formId})
	if from == to {
		toId++
	}
	mfg.g[from] = append(mfg.g[from], preEdge{to, toId, cap})
	mfg.g[to] = append(mfg.g[to], preEdge{from, formId, 0})
	return m
}

type edge struct {
	from, to  int
	cap, flow int
}

func (mfg *MfGraph) GetEdge(i int) edge {
	e := mfg.g[mfg.pos[i].x][mfg.pos[i].y]
	re := mfg.g[e.to][e.rev]
	return edge{mfg.pos[i].x, e.to, e.cap + re.cap, re.cap}
}

func (mfg *MfGraph) Edges() []edge {
	m := len(mfg.pos)
	result := make([]edge, 0)
	for i := 0; i < m; i++ {
		result = append(result, mfg.GetEdge(i))
	}
	return result
}

func (mfg *MfGraph) ChangeEdge(i int, newCap, newFlow int) {
	e := &(mfg.g[mfg.pos[i].x][mfg.pos[i].y])
	re := &(mfg.g[e.to][e.rev])
	e.cap = newCap - newFlow
	re.cap = newFlow
}

func (mfg *MfGraph) Flow(s, t int) int {
	return mfg.FlowCapped(s, t, int(^uint(0)>>1))
}

func (mfg *MfGraph) FlowCapped(s, t int, flowLimit int) int {
	level := make([]int, mfg.n)
	var bfs func()
	bfs = func() {
		for i := 0; i < len(level); i++ {
			level[i] = -1
		}
		level[s] = 0
		que := make([]int, 0, mfg.n)
		que = append(que, s)
		for len(que) > 0 {
			v := que[0]
			que = que[1:]
			for _, e := range mfg.g[v] {
				if e.cap == 0 || level[e.to] >= 0 {
					continue
				}
				level[e.to] = level[v] + 1
				if e.to == t {
					return
				}
				que = append(que, e.to)
			}
		}
	}
	iter := make([]int, mfg.n)
	var dfs func(int, int) int
	dfs = func(v, up int) int {
		if v == s {
			return up
		}
		res := 0
		levelV := level[v]
		for i := iter[v]; i < len(mfg.g[v]); i++ {
			e := mfg.g[v][i]
			cap := mfg.g[e.to][e.rev].cap
			if levelV <= level[e.to] || cap == 0 {
				continue
			}
			nextUp := up - res
			if cap < up-res {
				nextUp = cap
			}
			d := dfs(e.to, nextUp)
			if d <= 0 {
				continue
			}
			mfg.g[v][i].cap += d
			mfg.g[e.to][e.rev].cap -= d
			res += d
			if res == up {
				return res
			}
		}
		level[v] = mfg.n
		return res
	}
	flow := 0
	for flow < flowLimit {
		bfs()
		if level[t] == -1 {
			break
		}
		for i := 0; i < len(iter); i++ {
			iter[i] = 0
		}
		f := dfs(t, flowLimit-flow)
		if f == 0 {
			break
		}
		flow += f
	}
	return flow
}

func (mfg *MfGraph) MinCut(s int) []bool {
	visited := make([]bool, mfg.n)
	que := make([]int, 0, mfg.n)
	que = append(que, s)
	for len(que) > 0 {
		p := que[0]
		que = que[1:]
		visited[p] = true
		for _, e := range mfg.g[p] {
			if e.cap > 0 && !visited[e.to] {
				visited[e.to] = true
				que = append(que, e.to)
			}
		}
	}
	return visited
}
