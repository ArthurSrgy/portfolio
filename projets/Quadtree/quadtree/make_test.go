package quadtree

//QuadTest
import (
	"fmt"
	"testing"
)

func coinhautgauche(q Quadtree) *node {
	node := q.root
	for node.content == -1 {
		node = node.topLeftNode
	}
	return node
}

func TestCoinQuadtree(t *testing.T) {
	floorContent := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0},
	}
	quadtest := MakeFromArray(floorContent, 0, 0)
	nodeCoin := coinhautgauche(quadtest)
	if nodeCoin.content != 0 || nodeCoin.width != 3 || nodeCoin.height != 2 {
		t.Fail()
	}

}

func TestToutQuadtree(t *testing.T) {
	floorContent := [][]int{
		{0, 0, 3, 3},
		{0, 0, 2, 2},
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}
	quadtest := MakeFromArray(floorContent, 0, 0)
	node := quadtest.root
	if node.content != -1 || node.width != 4 || node.height != 4 {
		fmt.Println("le noeu est mauvais")
		t.Fail()
	}
	node = quadtest.root.topLeftNode
	if node.content != 0 || node.width != 2 || node.height != 2 || node.topLeftX != 0 || node.topLeftY != 0 {
		fmt.Println("le noeu topLeft est mauvais")
		t.Fail()
	}
	node = quadtest.root.topRightNode
	if node.content != -1 || node.width != 2 || node.height != 2 || node.topLeftX != 2 || node.topLeftY != 0 {
		fmt.Println("le noeu topRight est mauvais")
		t.Fail()
	}
	node = quadtest.root.topRightNode.topLeftNode
	if node.content != 3 || node.width != 1 || node.height != 1 || node.topLeftX != 2 || node.topLeftY != 0 {
		fmt.Println("le noeu topRight topLeft est mauvais")
		t.Fail()
	}
	node = quadtest.root.topRightNode.topRightNode
	if node.content != 3 || node.width != 1 || node.height != 1 || node.topLeftX != 3 || node.topLeftY != 0 {
		fmt.Println("le noeu topRight topRight est mauvais")
		t.Fail()
	}
	node = quadtest.root.topRightNode.bottomLeftNode
	if node.content != 2 || node.width != 1 || node.height != 1 || node.topLeftX != 2 || node.topLeftY != 1 {
		fmt.Println("le noeu topRight BottomLeft est mauvais")
		t.Fail()
	}
	node = quadtest.root.topRightNode.bottomRightNode
	if node.content != 2 || node.width != 1 || node.height != 1 || node.topLeftX != 3 || node.topLeftY != 1 {
		fmt.Println("le noeu topRight BottomRight est mauvais")
		t.Fail()
	}
	node = quadtest.root.bottomLeftNode
	if node.content != -1 || node.width != 2 || node.height != 2 || node.topLeftX != 0 || node.topLeftY != 2 {
		fmt.Println("le noeu BottomLeft est mauvais")
		t.Fail()
	}
	node = quadtest.root.bottomLeftNode.topLeftNode
	if node.content != 1 || node.width != 1 || node.height != 1 || node.topLeftX != 0 || node.topLeftY != 2 {
		fmt.Println("le noeu BottomLeft topLeft est mauvais")
		t.Fail()
	}
	node = quadtest.root.bottomLeftNode.topRightNode
	if node.content != 2 || node.width != 1 || node.height != 1 || node.topLeftX != 1 || node.topLeftY != 2 {
		fmt.Println("le noeu BottomLeft topRight est mauvais")
		t.Fail()
	}
	node = quadtest.root.bottomLeftNode.bottomLeftNode
	if node.content != 5 || node.width != 1 || node.height != 1 || node.topLeftX != 0 || node.topLeftY != 3 {
		fmt.Println("le noeu BottomLeft bootomLeft est mauvais")
		t.Fail()
	}
	node = quadtest.root.bottomLeftNode.bottomRightNode
	if node.content != 6 || node.width != 1 || node.height != 1 || node.topLeftX != 1 || node.topLeftY != 3 {
		fmt.Println("le noeu BottomLeft bottomRight est mauvais")
		t.Fail()
	}
	node = quadtest.root.bottomRightNode
	if node.content != -1 || node.width != 2 || node.height != 2 || node.topLeftX != 2 || node.topLeftY != 2 {
		fmt.Println("le noeu BottomRight est mauvais")
		t.Fail()
	}
	node = quadtest.root.bottomRightNode.topLeftNode
	if node.content != 3 || node.width != 1 || node.height != 1 || node.topLeftX != 2 || node.topLeftY != 2 {
		fmt.Println("le noeu BottomRight topLeft est mauvais")
		t.Fail()
	}
	node = quadtest.root.bottomRightNode.topRightNode
	if node.content != 4 || node.width != 1 || node.height != 1 || node.topLeftX != 3 || node.topLeftY != 2 {
		fmt.Println("le noeu BottomRight bottomLeft est mauvais")
		t.Fail()
	}
	node = quadtest.root.bottomRightNode.bottomLeftNode
	if node.content != 7 || node.width != 1 || node.height != 1 || node.topLeftX != 2 || node.topLeftY != 3 {
		fmt.Println("le noeu BottomRight bottomLeft est mauvais")
		t.Fail()
	}
	node = quadtest.root.bottomRightNode.bottomRightNode
	if node.content != 8 || node.width != 1 || node.height != 1 || node.topLeftX != 3 || node.topLeftY != 3 {
		fmt.Println("le noeu BottomRight bottomRight est mauvais")
		t.Fail()
	}
}
