package dvonn

import (
	"errors"
)

type HexNode struct {
	identifier  string
	stackLength int
	chipsStack  []Chip
	childNodes  map[EdgeType]*HexNode
}

func GetHexNode(id string) *HexNode {
	return &HexNode{id, 0, make([]Chip, 0), make(map[EdgeType]*HexNode)}
}

func (hn *HexNode) GetIdentifier() string {
	return hn.identifier
}

func (hn *HexNode) GetChildNodes() map[EdgeType]*HexNode {
	return hn.childNodes
}

func (hn *HexNode) SetChildNode(edgeType EdgeType, node *HexNode) {
	hn.childNodes[edgeType] = node
}

func (hn *HexNode) traverse() {
	// TODO: implement it
}

func (hn *HexNode) GetPossibleMovements() {
	// TODO: implement it
}

func (hn *HexNode) IsEmpty() bool {
	return len(hn.chipsStack) == 0
}

func (hn *HexNode) GetStackLength() int {
	return len(hn.chipsStack)
}

func (hn *HexNode) GetChipsStack() []Chip {
	return hn.chipsStack
}

func (hn *HexNode) GetTopChips() (Chip, error) {
	var res Chip
	if len(hn.chipsStack) > 0 {
		return hn.chipsStack[len(hn.chipsStack)-1], nil
	}
	return res, errors.New("no chips available on this node")
}

func (hn *HexNode) HasRedChips() bool {
	for _, chip := range hn.chipsStack {
		if chip.GetColor() == RED {
			return true
		}
	}
	return false
}

/*
 To Add chips on the box, accepts slice of chips (stack of chips)
*/
func (hn *HexNode) AddChips(chips []Chip) {
	hn.chipsStack = append(hn.chipsStack, chips...)
}

/*
 To make the Hex Box empty, this operation will be performed when player moves to other box
*/
func (hn *HexNode) Empty() {
	hn.chipsStack = make([]Chip, 0)
}

func (hn *HexNode) IsCompletelySurrounded() bool {
	return !(len(hn.GetChildNodes()) < 6)
}

func (hn *HexNode) HasFreeEdge() bool {
	return !hn.IsCompletelySurrounded()
}

/*
 returns: the list of Adjacent nodes which are in straight line from the current Node
 Adjacent nodes in straight line means: As there can be six adjacent nodes in a hexagonal node and those nodes
   can have another six adjacent nodes, but if you see not all of them will be in straight line.
*/
func (hn *HexNode) GetStraightAdjacentOnLevel(level int) []*HexNode {

	if level <= 0 {
		return nil
	}

	res := make([]*HexNode, 0)
	for edgeType, nodePtr := range hn.GetChildNodes() {
		_adjNode := _traverseInDirection(edgeType, level-1, nodePtr)
		if _adjNode != nil {
			res = append(res, _adjNode)
		}
	}
	return res
}

/*
 Summary: traverses the graph nodes only in a particular direction and returns the node which is exactly at
		  `depth` depth from the node passed in the argument. If it couldn't find any such node then it returns
		  nil.
 NOTE: the logic written here is strictly for getting adjacent nodes at some depth from some particular node,
		it should not be confused for finding possible moves from one HexNode to another HexNode coz movement of
		chips can be done across empty spaces.
*/
func _traverseInDirection(edgeType EdgeType, depth int, node *HexNode) *HexNode {

	if depth == 0 {
		return node
	}
	currentNode := node
	for depth > 0 {
		if val, ok := currentNode.childNodes[edgeType]; ok {
			currentNode = val
		} else { // return nil coz as we couldn't find node for that particular edge type even before depth is reached
			// so it's not possible to find it anyway after this point
			return nil
		}
		depth = depth - 1
	}
	return currentNode
}

func (hn *HexNode) presentIn(master []*HexNode) bool {
	for _, item := range master {
		if hn == item {
			return true
		}
	}
	return false
}
