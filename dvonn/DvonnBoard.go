package dvonn

import (
	"log"
)

type DvonnBoard struct {
	cells map[string]*HexNode
}

func GetDvonnBoard() *DvonnBoard {
	return &DvonnBoard{initializeBoard()}

}

func initializeBoard() map[string]*HexNode {
	var cells = make(map[string]*HexNode)
	ids := getBoardCellsIds()

	for _, id := range ids {
		cells[id] = GetHexNode(id)
	}
	connectCells(cells)
	return cells
}

func getBoardCellsIds() []string {
	return []string{"a1", "a2", "a3", "b1", "b2", "b3", "b4", "c1", "c2", "c3", "c4", "c5", "d1", "d2", "d3", "d4",
		"d5", "e1", "e2", "e3", "e4", "e5", "f1", "f2", "f3", "f4", "f5", "g1", "g2", "g3", "g4", "g5", "h1", "h2",
		"h3", "h4", "h5", "i1", "i2", "i3", "i4", "i5", "j2", "j3", "j4", "j5", "k3", "k4", "k5"}
}

/*
 NOTE: This method returns the cells on the board, however we should not be using this method coz, anything which we need
 to query on the cells should be written as a method in this class
*/
func (db *DvonnBoard) GetCells() map[string]*HexNode {
	return db.cells
}

func getNodesIdentifiers(nodes []*HexNode) []string {
	res := make([]string, 0)
	for _, node := range nodes {
		res = append(res, node.GetIdentifier())
	}
	return res
}

/*
 Adds a stack of chips to the box whose id is passed
 NOTE: we shouldn't add extra logic here, like if game in placement phase then don't allow filling multiple chips on
 stack, etc. This should be taken care by the implementer of this method or should be taken care at root level.
*/
func (db *DvonnBoard) Fill(id string, chips []Chip) {
	if node, ok := db.cells[id]; ok {
		node.AddChips(chips)
	} else {
		log.Fatal("[DvonnBoard.Fill()] wrong identifier passed while adding chip to board, id: ", id)
	}
}

/*
 Removes all chips from the Box whose id is passed
*/
func (db *DvonnBoard) DeFill(id string) {
	if node, ok := db.cells[id]; ok {
		node.Empty()
	} else {
		log.Fatal("[DvonnBoard.DeFill()] wrong identifier passed while removing chip from board, id: ", id)
	}
}

func (db *DvonnBoard) IsPlaceVacant(placeId string) bool {
	if val, ok := db.cells[placeId]; ok {
		return val.IsEmpty()
	}
	// TODO: put it as some error code || have a common method for this which can be called by other method
	log.Fatal("[DvonnBoard.IsPlaceVacant()] wrong identifier passed, id: ", placeId)
	return false
}

func (db *DvonnBoard) GetBoardSnapshot() {
	// TODO: implement it
}

/*
 Returns the identifiers of hexagonal nodes where a red chip is placed
 Solution: Iterates over all the board cells and check if red chip is present in chip stack.
		   Another alternative could be storing a list of ids for red chip addresses at the time Fill is called
		   but if we go with later approach we will need to add extra piece of code in Fill() method which will
		   violate single responsibility principle. So writing a extra method below.
*/
// TODO: write testcase
func (db *DvonnBoard) GetRedChipsIds() []string {
	res := make([]string, 0)
	for id, node := range db.cells {
		if node.HasRedChips() {
			res = append(res, id)
		}
	}
	if len(res) > 3 {
		log.Fatal("can not have more than 3 red chips in a game, please check the issue")
	}
	return res
}

func (db *DvonnBoard) GetCellIdsByStackColor(color ChipColor) []string {
	res := make([]string, 0)
	for id, node := range db.cells {
		topChip, err := node.GetTopChips()
		// there can be cases when board cell will be empty, handling that case by continuing the loop
		if err != nil {
			continue
		}
		if topChip.GetColor() == color {
			res = append(res, id)
		}
	}
	return res
}

func (db *DvonnBoard) GetCountOfPiecesControlledBy(color ChipColor) int {
	controlledIds := db.GetCellIdsByStackColor(color)
	count := 0
	for _, id := range controlledIds {
		count += db.GetCells()[id].GetStackLength()
	}
	return count
}

/*
 Returns cells which are not connected to the red chip Node by any link, if all are connected, then returns nil
 Solution:  - Get all the nodes where red chip are placed
			- Traverse from each of the above node where red chip is placed and store the visited node id
			- The ids which aren't visited will be the disconnected nodes
*/
func (db *DvonnBoard) GetDisconnectedCells() []string {
	redChipNodes := db.GetRedChipsIds()
	visitedCells := GetSet()

	for _, nodeId := range redChipNodes {
		visitedCells.AddMultiS(getNodesIdentifiers(db.traverseConnectedNodes(db.cells[nodeId])))

		if visitedCells.Size() == 49 { // if from one cell we could visit all nodes, then all cells are connected for sure
			return nil
		}
	}
	if visitedCells.Size() < 49 { // cells are disconnected
		cells := GetSet()
		cells.AddMultiS(getBoardCellsIds())
		return cells.DifferenceS(visitedCells)
	}
	return nil
}

/*
 for storing result to a mutable type, coz I was facing issue while updating result in slice from the recursive method
 for some reason (not sure what) slice isn't behaving like mutable type as its giving empty value out of the recursive
 method scope. So using a custom type and adding slice into it.
*/
type traverseStore struct {
	traverseRes []*HexNode
}

/*
 Returns the connected cells from the nodePtr passed
 connected means: traversal is possible b/w nodes which have chip kept on them
*/
func (db *DvonnBoard) traverseConnectedNodes(nodePtr *HexNode) []*HexNode {
	res := &traverseStore{make([]*HexNode, 0)}
	db._traverse(nodePtr, res)
	return res.traverseRes
}

/*
 doing DFS from the node
*/
func (db *DvonnBoard) _traverse(nodePtr *HexNode, res *traverseStore) {
	// exit condition
	if nodePtr == nil || nodePtr.presentIn(res.traverseRes) || nodePtr.IsEmpty() {
		return
	} else {
		res.traverseRes = append(res.traverseRes, nodePtr)
		for _, node := range nodePtr.GetChildNodes() {
			db._traverse(node, res)
		}
	}
}

func (db *DvonnBoard) RemoveDisconnectedCells() {
	disconnectedCells := db.GetDisconnectedCells()
	for _, cell := range disconnectedCells {
		db.DeFill(cell)
	}
}

func (db *DvonnBoard) IsCellEmpty(id string) bool {
	if node, ok := db.cells[id]; ok {
		return node.IsEmpty()
	}
	log.Fatal("[DvonnBoard.IsCellEmpty()] wrong identifier passed, id: ", id)
	return false
}

/*
 argument: source node id
 return: list of nodes where the chip stacks can be placed from the id passed in argument

 Solution: iterates over the node map in each of the six direction from the hexagonal node
		NOTE that maximum number of possible move for any Id can be 6, as there are 6 edges
		to the node, and nodes can be moved in any direction, but only along straight lines
*/
func (db *DvonnBoard) GetPossibleMoveFor(id string) []*HexNode {

	res := make([]*HexNode, 0)
	if node, ok := db.cells[id]; ok {
		// get possible adjacent nodes
		adjacentNodes := node.GetStraightAdjacentOnLevel(db.cells[id].GetStackLength())

		// check those places should not be empty
		for _, adjNode := range adjacentNodes {
			if !adjNode.IsEmpty() {
				res = append(res, adjNode)
			}
		}
		return res
	}
	log.Fatal("[DvonnBoard.GetPossibleMoveFor()] wrong identifier passed, id: ", id)
	return nil
}
