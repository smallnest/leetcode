package tree

import (
	"fmt"
	"testing"
)

func TestTree_BFS(t *testing.T) {
	root := &Node{
		100,
		[]*Node{
			{20, []*Node{
				{77, nil},
			}},
			{57, nil},
			{210, []*Node{
				{98, nil},
				{10, nil},
			}},
		},
	}

	tree := &Tree{root}

	ans := tree.BFS()

	fmt.Println(ans)
}

func TestTree_DFS(t *testing.T) {
	root := &Node{
		100,
		[]*Node{
			{20, []*Node{
				{77, nil},
			}},
			{57, nil},
			{210, []*Node{
				{98, nil},
				{10, nil},
			}},
		},
	}

	tree := &Tree{root}

	ans := tree.DFS()

	fmt.Println(ans)
}
