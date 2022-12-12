package graph

type Edge struct {
	From   int
	To     int
	Weight int
}

type Graph struct {
	starts    []int
	endID     int
	Adjacency map[int][]Edge
}

func NewGraph(input [][]rune) *Graph {
	vertexID := 0
	graph := Graph{Adjacency: make(map[int][]Edge), starts: make([]int, 0, 1)}
	graph.findStarts(input)
	graph.findEnd(input)
	for vPos := range input {
		for hPos := range input[vPos] {
			adj := graph.addNode(input, hPos, vPos, vertexID)
			graph.Adjacency[vertexID] = adj
			vertexID++
		}
	}

	return &graph
}

func (g *Graph) VerticesCount() int {
	return len(g.Adjacency)
}

func (g *Graph) Start() []int {
	return g.starts
}

func (g *Graph) End() int {
	return g.endID
}

func (g *Graph) addNode(input [][]rune, hPos int, vPos int, vertexID int) []Edge {
	vertex := input[vPos][hPos]
	adj := make([]Edge, 0, 4)

	g.tryAdd(&adj, vertex, vertexID, input, hPos-1, vPos)
	g.tryAdd(&adj, vertex, vertexID, input, hPos+1, vPos)
	g.tryAdd(&adj, vertex, vertexID, input, hPos, vPos-1)
	g.tryAdd(&adj, vertex, vertexID, input, hPos, vPos+1)

	return adj
}

func (g *Graph) tryAdd(adj *[]Edge, vertex rune, vertexID int, input [][]rune, hPos int, vPos int) {
	if vPos < 0 || vPos >= len(input) {
		return
	}

	if hPos < 0 || hPos >= len(input[vPos]) {
		return
	}

	if input[vPos][hPos]-vertex <= 1 {
		to := coordinatesToID(input, vPos, hPos)
		edge := Edge{From: vertexID, To: to, Weight: 1}
		*adj = append(*adj, edge)
	}
}

func (g *Graph) findStarts(input [][]rune) {
	for i := range input {
		for j := range input[i] {
			if input[i][j] == 'a' {
				g.starts = append(g.starts, coordinatesToID(input, i, j))
			}

			if input[i][j] == 'S' {
				l := len(g.starts)
				if l == 0 {
					g.starts = append(g.starts, coordinatesToID(input, i, j))
				} else {
					first := g.starts[0]
					g.starts[0] = coordinatesToID(input, i, j)
					g.starts = append(g.starts, first)
				}

				input[i][j] = 'a'
			}
		}
	}
}

func (g *Graph) findEnd(input [][]rune) {
	for i := range input {
		for j := range input[i] {
			if input[i][j] == 'E' {
				g.endID = coordinatesToID(input, i, j)
				input[i][j] = 'z'
			}
		}
	}

}

func coordinatesToID(input [][]rune, vPos int, hPos int) int {
	return len(input[0])*vPos + hPos
}
