// https://leetcode.com/problems/number-of-restricted-paths-from-first-to-last-node/
// Not my code.

package problem1786

import (
	"container/heap"
	"sort"
)

const (
	inf = 4_000_000_999
)

type edge struct {
	v, w int
}

func countRestrictedPaths(n int, edges [][]int) int {
	adj := make([][]edge, n)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		u-- // re-label
		v--
		adj[u] = append(adj[u], edge{v: v, w: w})
		adj[v] = append(adj[v], edge{v: u, w: w})
	}

	return dijkstra(adj, n)
}

func dijkstra(adj [][]edge, n int) int {
	dh := &distHeap{
		h:    make([]int, n),
		pos:  make([]int, n),
		dist: make([]int, n),
	}
	for i := range dh.dist {
		dh.dist[i] = inf
		dh.h[i] = i
		dh.pos[i] = i
	}
	dh.Swap(0, n-1)
	dh.dist[n-1] = 0

	// now start from n-1
	for dh.Len() > 0 {
		u := heap.Pop(dh).(int)

		for _, e := range adj[u] {
			if dh.dist[u]+e.w < dh.dist[e.v] {
				dh.dist[e.v] = dh.dist[u] + e.w
				heap.Fix(dh, dh.pos[e.v])
			}
		}
	}

	dist := dh.dist

	// reuse
	total := dh.pos
	for i := range total {
		total[i] = 0 // reset
	}
	total[n-1] = 1

	// reuse
	ind := dh.h[:n]
	for i := range ind {
		ind[i] = i
	}

	// sort by last distance
	sort.Slice(ind, func(i, j int) bool {
		return dist[ind[i]] < dist[ind[j]]
	})

	for _, u := range ind {
		if u == n-1 {
			continue
		}

		if dist[u] == inf {
			break
		}

		du := dist[u]
		for _, e := range adj[u] {
			if dist[e.v] < du {
				total[u] += total[e.v]
				total[u] %= 1_000_000_007
			}
		}
	}

	return total[0]
}

type distHeap struct {
	h    []int // nodes
	pos  []int // position in heap
	dist []int // distance
}

func (d *distHeap) Len() int { return len(d.h) }
func (d *distHeap) Less(i, j int) bool {
	return d.dist[d.h[i]] < d.dist[d.h[j]]
}
func (d *distHeap) Swap(i, j int) {
	u, v := d.h[i], d.h[j]
	d.h[i], d.h[j] = v, u
	d.pos[u], d.pos[v] = j, i
}
func (d *distHeap) Push(x any) {
	// no push, only pop
}
func (d *distHeap) Pop() any {
	n := len(d.h)
	x := d.h[n-1]
	d.h = d.h[:n-1]
	return x
}
