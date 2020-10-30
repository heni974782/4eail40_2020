// Package model contains the gameplay logic for the game of chess
package model

//TODO Implement type Board
import (
	"fmt"
	"github.com/rob195286/4eail40_2020/exercises/chess/chess/model/coord"
	"github.com/rob195286/4eail40_2020/exercises/chess/chess/model/piece"
)

// Classic 8x8 Chess board
type Classic [8][8]piece.Piece

func (c *Classic) String() string {
	panic("not implemented") // TODO: Implement
}

// PieceAt retrievves piece at give coordinates.
// Returns nil if no piece was found.
func (c *Classic) PieceAt(at coord.ChessCoordinates) piece.Piece {
	x, _ := at.Coord(0)
	y, _ := at.Coord(1)
	if c[x][y] != nil {
		return c[x][y]
	} else {
		return nil
	}
}


func (c *Classic) MovePiece(from coord.ChessCoordinates, to coord.ChessCoordinates) error {

	newx, _ := to.Coord(0)
	newy, _ := to.Coord(1)
	newCoords := coord.NewCartesian(newx, newy)

	if c.PieceAt(from) != nil {
		pieceChess := c.PieceAt(newCoords)
		c[newx][newy] = pieceChess
		return nil
	} else {
		return fmt.Errorf("bad coord, a piece exist in coord {%d}, impossible move to", from)
	}
}


func (c *Classic) PlacePieceAt(p piece.Piece, at coord.ChessCoordinates) error {
	x, _ := at.Coord(0)
	y, _ := at.Coord(1)
	if x > 8 || x < 0 || y < 0 || y > 8 {
		return fmt.Errorf("bad coord, %d or %d are out of range", x, y)
	} else if c.PieceAt(at) != nil {
		return fmt.Errorf("bad coord, a piece exist in coord {%d, %d}", x, y)
	} else {
		c[x][y] = p
		return nil
	}
}
