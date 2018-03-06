package main



import (
	"fmt"
	"math"
    "os"
)

type graph []*node

func main() {

	getGraph()

}



func (g graph) bfs(s int64) {
	sorgent := g[s-1]

	totalNumNode := len(g)
	infinite := int64(math.MaxInt64)

	//for each vertex u ∈ G.V- {s}
	for _, v := range g {
		if v == sorgent {
			continue
		}

		v.color = 'w'
		v.distance = infinite
		v.father = nil
	}

	sorgent.color = 'g'
	sorgent.distance = 0

	//Queue
	sq := make([]*node, 0, totalNumNode)
	sq = append(sq, sorgent) //add the s to to queue
	numberNodeQ := 1         //number node in the queue

	for numberNodeQ > 0 { //until the queue is empty

		u := sq[0]  //get a node from queue
		sq = sq[1:] //move the queue to next element
		numberNodeQ--

		adjU := u.next
		for adjU != nil {
			v := g[adjU.id-1]

			if v.color == 'w' {
				v.color = 'g'
				v.distance = u.distance + adjU.distance
				v.father = u
				sq = append(sq, v)
				numberNodeQ++

			}

			adjU = adjU.next
		}

		u.color = 'b'
	}

}

//print the path from S to D in the graph if exist
func (g graph) printDistance(s int64) {
	infinite := int64(math.MaxInt64)
	for _, v := range g {
		if v.id == s {
			continue
		}
		if v.distance == infinite {
			fmt.Printf("%d ", -1)

		} else {

			fmt.Printf("%d ", v.distance)
		}

	}
	fmt.Printf("\n")

}






//false= directed graph, true = undirected graph
var undirected = false

type node struct {
	id       int64 //
	next     *node //
	color    rune  //color  of the node
	father   *node
	distance int64 // distance from the source to vertex u
}

// return new node (initialized to nil)
func getNode(v int64) *node {

	return &node{
		id: v,
	}
}

// return new node pointing to n
func setNeighbor(verNumber int64, n *node) *node {

	return &node{
		id:       verNumber,
		distance: 6,
		next:     n,
	}

}

// printGraph all the graph
func (g graph) printGraph() {
	for i, v := range g {
		fmt.Printf("%d (%v)", i+1, v.distance)
		//fmt.Printf("%d", i)
		v.next.printNeighbors()

	}
}

//printNeighbors all the neighbor of the current Node
func (n *node) printNeighbors() {
	tmp := n
	for tmp != nil {
		fmt.Printf("---->%d(%v)", tmp.id, tmp.distance)
		tmp = tmp.next
	}
	fmt.Println()

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// build the graph
func getGraph() {
	var g graph
	var q int
	var n, m int
	var v1, v2, idOrigin int64

	f := os.Stdin
	

    _, err := fmt.Fscanf(f, "%d\n", &q)
	check(err)

	for i := 0; i < q; i++ {
		_, err = fmt.Fscanf(f, "%d %d\n", &n, &m)
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
			_, err = fmt.Fscanf(f, "%d %d\n", &v1, &v2)
			check(err)
			g[v1-1].next = setNeighbor(v2, g[v1-1].next)
			g[v2-1].next = setNeighbor(v1, g[v2-1].next)
		}

		_, err = fmt.Fscanf(f, "%d\n", &idOrigin)
		check(err)
		// fmt.Println("\n original graph")
		// g.printGraph()
		g.bfs(idOrigin)
		// fmt.Println("\n new graph")
		// g.printGraph()
		// fmt.Println("\nshortest path")
		g.printDistance(idOrigin)

	}

}

