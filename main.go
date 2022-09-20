package main

import (
	"container/heap"
	"fmt"
)

type Edge struct {
	Node int
	Weight int
}

type Graph struct {
	Nodes map[int][]Edge
}

func NewGraph() *Graph {
	return &Graph{ Nodes: make(map[int][]Edge) }
}

func (g *Graph) AddEdge(startNode, endNode, weight int) {
	g.Nodes[startNode] = append(g.Nodes[startNode], Edge{endNode, weight})
}

func (g *Graph) ShortestPath(startNode, endNode int) []int {
	var routes []int
	visited := make(map[int]bool)

	pq := NewPriorityQueue()
	startRoute := Route{0, []int{startNode}}
	pq.Push(startRoute)

	for pq.Len() > 0 {
		currentRoute := pq.Pop()
		currentNode := currentRoute.Nodes[len(currentRoute.Nodes)-1]

		if visited[currentNode] {
			continue
		}

		if currentNode == endNode {
			routes = currentRoute.Nodes
			break
		}

		visited[currentNode] = true

		for _, edge := range g.Nodes[currentNode] {
			if !visited[edge.Node] {
				totalDistance := currentRoute.Distance + edge.Weight
				newRoute := Route {
					Distance: totalDistance,
					Nodes: append([]int{}, append(currentRoute.Nodes, edge.Node)...),
				}

				pq.Push(newRoute)
			}
		}
	}

	return routes
}

func test_case(orders [][]int, routes [][]int) {
	// orderLength := len(orders)
	// routeLength := len(routes)

	// fmt.Println("Order: ", orderLength)
	// fmt.Println("Route: ", routeLength)

	graph := NewGraph()

	for _, route := range routes {
		// fmt.Println(route[0])
		startNode := route[0]
		endNode := route[1]
		weight := route[2]
		graph.AddEdge(startNode, endNode, weight)
	}

	var optimalRoute []int

	initialPosition := 1
	initialOrderPick := orders[0][0]
	optimalRoute = graph.ShortestPath(initialPosition,initialOrderPick)

	for _, order := range orders {
		pick := order[0]
		destination := order[1]
		route := graph.ShortestPath(pick, destination)

		optimalRoute = append(optimalRoute[:len(optimalRoute)-1], route...)
	}

	fmt.Println(optimalRoute)
}

func main() {
	var order int
	var route int

	fmt.Scanf("%d", &order)

	orderList := make([][]int, order)

	for i := 0; i < order; i++ {
		var pick, destination int
		fmt.Scanf("%d %d", &pick, &destination)

		orderRoute := make([]int, 2)
		orderRoute[0] = pick
		orderRoute[1] = destination

		orderList[i] = orderRoute
	}

	
	fmt.Scanf("%d", &route)

	routeList := make([][]int, route)

	for i := 0; i < route; i++ {
		var node1, node2, cost int
		fmt.Scanf("%d %d %d", &node1, &node2, &cost)

		routeCost := make([]int, 3)	
		routeCost[0] = node1
		routeCost[1] = node2
		routeCost[2] = cost

		routeList[i] = routeCost
	}

	// fmt.Println(orderList)
	// fmt.Println(routeList)

	test_case(orderList, routeList)
}

type PriorityQueue struct {
	Heap *Routes
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{&Routes{}}
}

func (pq *PriorityQueue) Len() int {
	return pq.Heap.Len()
}

func (pq *PriorityQueue) Push(route Route) {
	heap.Push(pq.Heap, route)
}

func (pq *PriorityQueue) Pop() Route {
	return heap.Pop(pq.Heap).(Route)
}

// func (pq *PriorityQueue) Routes() []int {
// 	return pq.Heap[0].Nodes
// }

type Route struct {
	Distance int
	Nodes []int
}

type Routes []Route

func (q Routes) Len() int { return len(q) }
func (q Routes) Less(i,j int) bool {
	// fmt.Println(q[i])
	// fmt.Println(q[0])
	// fmt.Println("Q[i]: ", q[i].Distance)
	// fmt.Println("Q[j]:", q[j].Distance)
	// return false
	// fmt.Println(fmt.Sprintf("Q[%d] Distance: %d <> Q[%d] Distance: %d ", i, q[i].Distance, j, q[j].Distance))
	return q[i].Distance < q[j].Distance
}

func (q Routes) Swap(i,j int) {
	q[i], q[j] = q[j], q[i]
}

func (q *Routes) Push(x interface{}) {
	// n := len(*q)
	item := x.(Route)
	*q = append(*q, item)
}

func (q *Routes) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	*q = old[0:n-1]
	// fmt.Println("Pop: ", item)
	return item
}