package main

import (
	"container/heap"
	"fmt"
	"math"
    "os"
)

type graph []*node



//false= directed graph, true = undirected graph
var undirected = false
var time int

type node struct {
	id           int64 //id node
	next         *node // next neighbor
	weight       int64 // weight of the edge
	father       *node // father of the node
	priority     int64 // The priority of the item in the queue.
	isUndirected bool  //false= directed graph, true = undirected graph
	letter       string
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}


/**********************************************************************************/

func main() {

	getGraph()

}

func (g graph) mstPrime(origin int64) {
	var total int64
	var v *node
	infinite := int64(math.MaxInt64)
	checkM := make(map[int64]*node)  //map check if a node is still in the queue
	path := make([]int64, 0, len(g)) //path from origin g[origin] until end

	//set all the vertex priority as infinite
	for i, u := range g {
		u.priority = infinite
		u.father = nil
		u.index = i
		checkM[u.id] = u
	}
	//g[origin].father = g[origin]
	g[origin].priority = 0
	checkM[g[origin].id] = g[origin]

	heap.Init(&g) //push the graph in a priority queue

	path = append(path, g[origin].id)

	sizeQ := g.Len()

	for sizeQ > 0 {

		u := heap.Pop(&g).(*node) //get the veertex in queue with the lowest priority

		delete(checkM, u.id) //delete node from the map

		sizeQ = g.Len() //update size of the queue

		if u.father != nil {
			total += u.priority
			//	fmt.Printf("%v ----> %v (%v)\n", string((u.father.id-1)+'a'), u.id, string((u.id-1)+'a'))
			path = append(path, u.id)
		}

		//for each vertex v ∈ G.adj[u] (all edges from the node u)
		adjU := u.next
		for adjU != nil {
			//if v∈Q
			if val, ok := checkM[adjU.id]; ok {
				v = val
				// if  w(u,v)<v.key
				if adjU.weight < v.priority {
					//update priority and father
					v.father = u
					v.priority = adjU.weight
					g.update(v)

				}
			}

			adjU = adjU.next
		}

	}
	fmt.Println(total)

}


/************************************************************************************/

// return new node (initialized to nil)
func getNode(v int64) *node {

	return &node{
		id:     v,
		letter: string((v - 1) + 'a'),
	}
}

// setNeighbor make a node with pointing to n
func setNeighbor(verNumber int64, n *node, weight int64, isUndirected bool) *node {

	return &node{
		id:           verNumber,
		next:         n,
		weight:       weight,
		isUndirected: isUndirected,
		letter:       string((verNumber - 1) + 'a'),
	}

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// build the graph
func getGraph() {
	var g graph

	var n, m int
	var v1, v2, weight, idOrigin int64
	f := os.Stdin
	_, err := fmt.Fscanf(f, "%d %d\n", &n, &m)
	check(err)
	lenG := n
	g = make(graph, lenG)

	//make n nodes
	for j := 0; j < n; j++ {
		id := int64(j + 1) //id start from 1
		g[j] = getNode(id)
	}
	//make  m paths
	for j := 0; j < m; j++ {
		_, err = fmt.Fscanf(f, "%d %d %d\n", &v1, &v2, &weight)
		check(err)
		g[v1-1].next = setNeighbor(v2, g[v1-1].next, weight, undirected)
		g[v2-1].next = setNeighbor(v1, g[v2-1].next, weight, undirected)

	}

	_, err = fmt.Fscanf(f, "%d\n", &idOrigin)
	check(err)

	g.mstPrime(idOrigin - 1)

}
/**********************************************************************************/


func (pq *graph) Len() int { return len(*pq) }

func (pq graph) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority <= pq[j].priority
}

func (pq graph) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j

}


func (pq *graph) Push(x interface{}) {

	n := len(*pq)
	item := x.(*node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *graph) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority of the node  in the queue.
func (pq *graph) update(v *node) {

	heap.Fix(pq, v.index)
}



