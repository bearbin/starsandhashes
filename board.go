package main

// The width and height of the individual boards and the game board as a whole.
// This value must be odd, and values of 3 or 5 are recommended.
const XY int = 5

// The coorindates of the central cell of the board. Since the board arrays are
// zero-indexed, it is one less than expected.
const CENTRE int = (XY - 1) / 2

type cell interface {
	Owner() *player
}

type board interface {
	cell
	SelectCell(x int, y int, p *player) *cell
}

type singlecell struct {
	owner *player
}

func (c *singlecell) Owner() *player {
	return c.owner
}

type miniboard struct {
	contents [XY][XY]cell
	owner    *player
}

// Func Owner() calculates if there is a winner of the miniboard. The results
// of this function are undefined for a board layout with more than one winner.
func (brd *miniboard) Owner() *player {
	// First check to see for a cached winner.
	if brd.owner != nil {
		return brd.owner
	}
	var winner *player
	// Check for horizontal wins.
	for y := 0; y < XY; y++ {
		tempwinner := brd.contents[0][y].Owner()
		for x := 1; x < XY; x++ {
			if brd.contents[x][y].Owner() != tempwinner {
				tempwinner = nil
			}
		}
		if tempwinner != nil {
			winner = tempwinner
		}
	}
	// Check for vertical wins.
	for x := 0; x < XY; x++ {
		tempwinner := brd.contents[x][0].Owner()
		for y := 1; y < XY; y++ {
			if brd.contents[x][y].Owner() != tempwinner {
				tempwinner = nil
			}
		}
		if tempwinner != nil {
			winner = tempwinner
		}
	}
	// Check for diagonal wins.
	// There is no diagonal win without occupying the centre.
	if brd.contents[CENTRE][CENTRE].Owner() == nil {
		return nil
	}
	tempwinner := brd.contents[0][0].Owner()
	for x := 1; x < XY; x++ {
		if brd.contents[x][x].Owner() != tempwinner {
			tempwinner = nil
		}
	}
	if tempwinner != nil {
		winner = tempwinner
	}
	tempwinner = brd.contents[0][XY-1].Owner()
	for x := 1; x < XY; x++ {
		if brd.contents[x][(XY-1)-x].Owner() != tempwinner {
			tempwinner = nil
		}
	}
	if tempwinner != nil {
		winner = tempwinner
	}
	brd.owner = winner
	return winner
}

type gameboard struct {
	contents [XY][XY]miniboard
	owner    *player
}

// Func Owner() calculates if there is a winner of the gameboard. The results
// of this function are undefined for a board layout with more than one winner.
func (brd *gameboard) Owner() *player {
	// First check to see for a cached winner.
	if brd.owner != nil {
		return brd.owner
	}
	var winner *player
	// Check for horizontal wins.
	for y := 0; y < XY; y++ {
		tempwinner := brd.contents[0][y].Owner()
		for x := 1; x < XY; x++ {
			if brd.contents[x][y].Owner() != tempwinner {
				tempwinner = nil
			}
		}
		if tempwinner != nil {
			winner = tempwinner
		}
	}
	// Check for vertical wins.
	for x := 0; x < XY; x++ {
		tempwinner := brd.contents[x][0].Owner()
		for y := 1; y < XY; y++ {
			if brd.contents[x][y].Owner() != tempwinner {
				tempwinner = nil
			}
		}
		if tempwinner != nil {
			winner = tempwinner
		}
	}
	// Check for diagonal wins.
	// There is no diagonal win without occupying the centre.
	if brd.contents[CENTRE][CENTRE].Owner() == nil {
		return nil
	}
	tempwinner := brd.contents[0][0].Owner()
	for x := 1; x < XY; x++ {
		if brd.contents[x][x].Owner() != tempwinner {
			tempwinner = nil
		}
	}
	if tempwinner != nil {
		winner = tempwinner
	}
	tempwinner = brd.contents[0][XY-1].Owner()
	for x := 1; x < XY; x++ {
		if brd.contents[x][(XY-1)-x].Owner() != tempwinner {
			tempwinner = nil
		}
	}
	if tempwinner != nil {
		winner = tempwinner
	}
	brd.owner = winner
	return winner
}
