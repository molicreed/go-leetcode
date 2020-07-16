package p0785

var p map[int]int

func isBipartite(graph [][]int) bool {
	p = make(map[int]int, len(graph))
	for i := range graph {
		if p[i] != 0 {
			continue
		}
		p[i] = 1
		if !dfs(graph, i, 2) {
			return false
		}
	}
	return true
}

func dfs(graph [][]int, i int, target int) bool {
	if len(graph[i]) == 0 {
		return true
	}
	for _, j := range graph[i] {
		if p[j] != 0 {
			if p[j] != target {
				return false
			}
			continue
		}
		p[j] = target
		var ok bool
		if target == 1 {
			ok = dfs(graph, j, 2)
		} else {
			ok = dfs(graph, j, 1)
		}
		if !ok {
			return false
		}
	}
	return true
}
