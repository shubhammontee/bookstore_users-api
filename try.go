package main

type node struct {
	parent     *node
	data       int
	leftChild  *node
	rightChild *node
	infected   bool
}

func infectedNodePesrSec(infectedNode *node, time int) int {
	if (*(*infectedNode).parent).infected && (*(*infectedNode).rightChild).infected && (*(*infectedNode).leftChild).infected {
		return time
	}

	//infecting all the child and parent node
	(*infectedNode).infected = true
	(*(*infectedNode).parent).infected = true
	(*(*infectedNode).rightChild).infected = true
	(*(*infectedNode).leftChild).infected = true

	infectedNodePesrSec((*infectedNode).parent, time+1)
	if (*infectedNode).rightChild != nil {
		infectedNodePesrSec((*infectedNode).rightChild, time+1)
	}
	if (*infectedNode).leftChild != nil {
		infectedNodePesrSec((*infectedNode).leftChild, time+1)
	}

	return time

}

func main() {

}
