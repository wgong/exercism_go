package tree

import (
	"errors"
	"fmt"
	"sort"
)

// Record for input record
type Record struct {
	ID, Parent int
}

// Node for tree node
type Node struct {
	ID       int
	Children []*Node
}

// Build builds tree from input Records
func Build(records []Record) (*Node, error) {
	mapIDNode := make(map[int]*Node)
	var node Node

	for _, rec := range records {
		id, parent := rec.ID, rec.Parent

		switch {
		case id == 0: // root Node
			if _, exist := mapIDNode[id]; !exist {
				mapIDNode[id] = &Node{ID: 0, Children: nil}
			}
		case id > 0 && id > parent:
			if pNode, exist := mapIDNode[id]; exist {
				fmt.Println("node exists: ", id, parent)
				node = *pNode
			} else {
				node := Node{ID: id, Children: nil}
				mapIDNode[id] = &node
			}
			// check parent
			if pNode, exist := mapIDNode[parent]; exist {
				pNode.Children = addChildInOrder(pNode.Children, &node)
			} else {
				children := []*Node{&node}
				parentNode := Node{ID: parent, Children: children}
				mapIDNode[parent] = &parentNode
			}
		default:
			return nil, errors.New("Invalid record")
		}

	}

	return mapIDNode[0], nil
}

func addChildInOrder(children []*Node, child *Node) []*Node {
	mapIDNode := make(map[int]*Node)
	mapIDNode[child.ID] = child

	idList := []int{child.ID}
	for _, c := range children {
		idList = append(idList, c.ID)
		mapIDNode[c.ID] = c
	}
	sort.Ints(idList)

	children2 := []*Node{mapIDNode[idList[0]]}
	for i := 1; i < len(idList); i++ {
		children2 = append(children2, mapIDNode[idList[i]])
	}
	return children2
}
