package dvonn

import "log"

type DvonnBoard struct {
	cells map[string]*HexNode
}

func GetDvonnBoard() *DvonnBoard {
	return &DvonnBoard{initializeBoard()}

}

func initializeBoard() map[string]*HexNode {
	var cells = make(map[string]*HexNode)
	ids := []string{"a1", "a2", "a3", "b1", "b2", "b3", "b4", "c1", "c2", "c3", "c4", "c5", "d1", "d2", "d3", "d4",
		"d5", "e1", "e2", "e3", "e4", "e5", "f1", "f2", "f3", "f4", "f5", "g1", "g2", "g3", "g4", "g5", "h1", "h2",
		"h3", "h4", "h5", "j2", "j3", "j4", "j5", "k3", "k4", "k5"}

	for _, id := range ids {
		cells[id] = GetHexNode(id)
	}
	// TODO: connect the cells as required
	return cells
}


/*
 NOTE: This method returns the cells on the board, however we should not be using this method coz, anything which we need
 to query on the cells should be written as a method in this class
 */
func (db *DvonnBoard) GetCells() map[string]*HexNode {
	return db.cells;
}

/*
 Adds a stack of chips to the box whose id is passed
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

/*
 Returns cells which are not connected to the red chip Node by any link
 Solution:  - Get all the nodes where red chip are placed
			- Traverse from each of the above node where red chip is placed and store the visited node id
			- The ids which aren't visited will be the disconnected nodes
 */
func (db *DvonnBoard) GetDisconnectedCells() []string {
	redChipNodes := db.GetRedChipsIds()
	visitedCells := make([]*HexNode, 0)

	for _, nodeId := range redChipNodes {
		visitedCells = append(visitedCells, db.traverseConnectedNodes(db.cells[nodeId])...)
	}
	// todo: fetch ids and find difference from the total ids
	return nil
}

/*
 Returns the connected cells from the nodePtr passed
 connected means: traversal is possible b/w nodes which have chip kept on them
 */
func (db *DvonnBoard) traverseConnectedNodes(nodePtr *HexNode) []*HexNode {
	res := make([]*HexNode, 0)
	db._traverse(nodePtr, res)
	return res
}


/*
 doing DFS from the node
 */
func (db *DvonnBoard) _traverse(nodePtr *HexNode, visitedNode []*HexNode) {
	// exit condition
	if nodePtr == nil || nodePtr.presentIn(visitedNode) || nodePtr.IsEmpty() {
		return
	} else {
		visitedNode = append(visitedNode, nodePtr)
		for _, node := range nodePtr.GetChildNodes() {
			db._traverse(node, visitedNode)
		}
	}
}

