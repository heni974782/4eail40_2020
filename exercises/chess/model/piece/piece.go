package piece

import (
	"fmt"
	"github.com/rob195286/4eail40_2020/exercises/chess/chess/model/coord"
	"github.com/rob195286/4eail40_2020/exercises/chess/chess/model/player"
)

type EchecPiece struct {
	coord coord.Cartesian
	color player.Color
}

func NewEchecPiece(coord coord.Cartesian, color player.Color) EchecPiece {
	return EchecPiece{coord, color}
}

func (e EchecPiece) Color() player.Color {
	return e.color
}

func (e EchecPiece) Moves(isCapture bool) map[coord.ChessCoordinates]bool {
	return map[coord.ChessCoordinates]bool{}
}
func (e EchecPiece) String() string {
	return fmt.Sprintf("coord : {%d}, color : %d", e.coord, e.color)
}

func (e EchecPiece) Coord(n int) (int, error) {
	switch n {
	case 0:
		return 0, nil
	case 1:
		return 1, nil

	}
	return 0, fmt.Errorf("unknown coord component %d", n)
}
