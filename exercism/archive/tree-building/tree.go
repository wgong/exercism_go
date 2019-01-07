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
	mapNodeDefined := make(map[Record]bool)
	bHasRoot := false

	for _, rec := range records {

		if mapNodeDefined[rec] {
			return nil, errors.New("dup node")
		}

		id, parent := rec.ID, rec.Parent
		fmt.Println("\nid, parent = ", id, parent)

		switch {
		case id == 0 && id == parent: // root Node
			fmt.Println("set ROOT")
			if _, exist := mapIDNode[id]; !exist {
				mapIDNode[id] = &Node{ID: id, Children: nil}
			}
			bHasRoot = true
		case id > 0 && id > parent:

			pNode, existNode := mapIDNode[id]
			if !existNode {
				fmt.Println("Create Node")
				pNode = &Node{ID: id, Children: nil}
				mapIDNode[id] = pNode

			} else {

				fmt.Println("Exist Node")
			}

			pParNode, existParNode := mapIDNode[parent]
			if !existParNode {
				fmt.Println("Create parent")
				pParNode = &Node{ID: parent, Children: []*Node{pNode}}
				mapIDNode[parent] = pParNode

			} else {
				fmt.Println("Exist parent")
				pParNode.Children = addChildInOrder(pParNode.Children, pNode)
			}
		default:
			return nil, errors.New("invalid record")
		}
		fmt.Println(mapIDNode)

		mapNodeDefined[rec] = true

	}

	if !bHasRoot {
		return nil, errors.New("no root node")
	}
	return mapIDNode[0], nil
}

func addChildInOrder(children []*Node, child *Node) []*Node {
	if children == nil || len(children) == 0 {
		return []*Node{child}
	}

	mIDNode := make(map[int]*Node)
	mIDNode[child.ID] = child
	for _, c := range children {
		mIDNode[c.ID] = c
	}

	idList := []int{}
	for k := range mIDNode {
		idList = append(idList, k)
	}
	sort.Ints(idList)

	children2 := []*Node{}
	for i := 0; i < len(idList); i++ {
		children2 = append(children2, mIDNode[idList[i]])
	}
	return children2
}
