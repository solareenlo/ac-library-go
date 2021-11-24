package segtree

type E func() string
type Op func(a, b string) string
type Compare func(v string) bool
type Segtree struct {
	n    int
	size int
	log  int
	d    []string
	e    E
	op   Op
}

func NewSegtree(v []string, op Op, e E) *Segtree {
	seg := new(Segtree)
	seg.n = len(v)
	seg.log = seg.ceilPow2(seg.n)
	seg.size = 1 << uint(seg.log)
	seg.d = make([]string, 2*seg.size)
	seg.e = e
	seg.op = op
	for i := range seg.d {
		seg.d[i] = seg.e()
	}
	for i := 0; i < seg.n; i++ {
		// val, _ := strconv.Atoi(v[i])
		seg.d[seg.size+i] = v[i]
	}
	for i := seg.size - 1; i >= 1; i-- {
		seg.update(i)
	}
	return seg
}

func (seg *Segtree) update(k int) {
	seg.d[k] = seg.op(seg.d[2*k], seg.d[2*k+1])
}

func (seg *Segtree) Set(p int, x string) {
	if p < 0 || seg.n <= p {
		panic("error: Set")
	}
	p += seg.size
	seg.d[p] = x
	for i := 1; i <= seg.log; i++ {
		seg.update(p >> uint(i))
	}
}

func (seg *Segtree) Get(p int) string {
	if p < 0 || seg.n <= p {
		panic("error: Get")
	}
	return seg.d[p+seg.size]
}

func (seg *Segtree) Prod(l, r int) string {
	if l < 0 || r < l || seg.n < r {
		panic("error: Prod")
	}
	sml, smr := seg.e(), seg.e()
	l += seg.size
	r += seg.size
	for l < r {
		if (l & 1) == 1 {
			sml = seg.op(sml, seg.d[l])
			l++
		}
		if (r & 1) == 1 {
			r--
			smr = seg.op(seg.d[r], smr)
		}
		l >>= 1
		r >>= 1
	}
	return seg.op(sml, smr)
}

func (seg *Segtree) AllProd() string {
	return seg.d[1]
}

func (seg *Segtree) MaxRight(l int, cmp Compare) int {
	if l < 0 || seg.n < l {
		panic("error: MaxRight")
	}
	if !cmp(seg.e()) {
		panic("error: MaxRight")
	}
	if l == seg.n {
		return seg.n
	}
	l += seg.size
	sm := seg.e()
	for {
		for l%2 == 0 {
			l >>= 1
		}
		if !cmp(seg.op(sm, seg.d[l])) {
			for l < seg.size {
				l = 2 * l
				if cmp(seg.op(sm, seg.d[l])) {
					sm = seg.op(sm, seg.d[l])
					l++
				}
			}
			return l - seg.size
		}
		sm = seg.op(sm, seg.d[l])
		l++
		if l&-l == l {
			break
		}
	}
	return seg.n
}

func (seg *Segtree) MinLeft(r int, cmp Compare) int {
	if r < 0 || seg.n < r {
		panic("error: MinLeft")
	}
	if !cmp(seg.e()) {
		panic("error: MinLeft")
	}
	if r == 0 {
		return 0
	}
	r += seg.size
	sm := seg.e()
	for {
		r--
		for r > 1 && r%2 != 0 {
			r >>= 1
		}
		if !cmp(seg.op(seg.d[r], sm)) {
			for r < seg.size {
				r = 2*r + 1
				if cmp(seg.op(seg.d[r], sm)) {
					sm = seg.op(seg.d[r], sm)
					r--
				}
			}
			return r + 1 - seg.size
		}
		sm = seg.op(seg.d[r], sm)
		if r&-r == r {
			break
		}
	}
	return 0
}

func (seg *Segtree) ceilPow2(n int) int {
	x := 0
	for (uint(1) << x) < uint(n) {
		x++
	}
	return x
}
