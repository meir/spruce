package spruce

import "context"

type Node struct {
	States   []State
	Activate func(context.Context) (context.Context, bool)
}

var nodes = []Node{}

func RegisterNode(node Node) {
	nodes = append(nodes, node)
}

func GetNodes() []Node {
	return nodes
}

func GetNodesByState(state State) []Node {
	var res []Node
	for _, node := range nodes {
		for _, s := range node.States {
			if s == state {
				res = append(res, node)
				break
			}
		}
	}
	return res
}
