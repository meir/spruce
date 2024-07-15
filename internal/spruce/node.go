package spruce

import "context"

type Node struct {
	Priority uint8
	States   []State
	Activate func(context.Context) (context.Context, bool)
}

var nodes = []Node{}

func RegisterNode(node Node) {
	nodes = append(nodes, node)

	// Sort nodes by priority
	for i := 0; i < len(nodes); i++ {
		for j := i; j < len(nodes); j++ {
			if nodes[i].Priority < nodes[j].Priority {
				nodes[i], nodes[j] = nodes[j], nodes[i]
			}
		}
	}
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
