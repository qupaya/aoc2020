package day10

// Digraph directed graph
type Digraph map[Vertex]Edges

// Vertex vertex of digraph
type Vertex int

// Edges list of outgoing edges of a vertex
type Edges []Vertex

// Path path between vertices
type Path []Vertex

// AddEdge add an edge between two vertices
func AddEdge(digraph *Digraph, fromVertex Vertex, toVertex Vertex) {
	(*digraph)[fromVertex] = append((*digraph)[fromVertex], toVertex)
}

func enqueue(queue *[]Vertex, v Vertex) {
	*queue = append(*queue, v)
}

func dequeue(queue *[]Vertex) Vertex {
	v := (*queue)[0]
	*queue = (*queue)[1:]
	return v
}

// CountPaths count all paths between two vertices
func CountPaths(digraph *Digraph, from Vertex, to Vertex) int {
	count := 0
	queue := []Vertex{from}

	for len(queue) > 0 {
		v := dequeue(&queue)
		if v == to {
			count++
			// queue = []Vertex{root} // restart
			continue
		}

		for _, next := range (*digraph)[v] {
			enqueue(&queue, next)
		}

	}

	return count //panic("did not find solution")
}
