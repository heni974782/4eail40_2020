package board

import (
	"github.com/rob195286/4eail40_2020/exercises/chess/chess/model/coord"
	"github.com/rob195286/4eail40_2020/exercises/chess/chess/model/piece"
	"github.com/rob195286/4eail40_2020/exercises/chess/chess/model/player"
	"testing"
)

type mockCoord int

// Coord returns x if n==0, y if n==1
func (c mockCoord) Coord(n int) (int, error) {
	return int(c), nil
}

func (c mockCoord) String() string {
	return "1"
}

func TestClassic_MovePiece(t *testing.T) {
	type args struct {
		from coord.ChessCoordinates
		to   coord.ChessCoordinates
	}
	tests := []struct {
		name    string
		c       *Classic
		args    args
		wantErr bool
	}{
		{
			"testmock",
			&Classic{},
			args{mockCoord(0), mockCoord(1)},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.MovePiece(tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("Classic.MovePiece() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

/*

func TestClassic_PlacePieceAt1(t *testing.T) {
	coords := coord.NewCartesian(1,2)
	piece := piece.NewEchecPiece(coords, player.Color(0))
	board := Classic{}
	//c1, _ := coords.Coord(0)
	//c2, _ := coords.Coord(1)
	//board[c1][c2] = piece
	got := board.PlacePieceAt(piece, coords)

	if nil != got{
		t.Errorf("expected nil, but got %d", got)
	}

}
*/
func TestClassic_PlacePieceAt(t *testing.T) {
	coords := coord.NewCartesian(1, 2)
	x, _ := coords.Coord(0)
	y, _ := coords.Coord(1)
	pieceChess := piece.NewEchecPiece(coords, player.Color(0))
	board := Classic{}
	board[x][y] = pieceChess

	type args struct {
		p  piece.Piece
		at coord.ChessCoordinates
	}
	tests := []struct {
		name    string
		c       Classic
		args    args
		wantErr bool
	}{
		{
			"ok",
			board,
			args{pieceChess, coords},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.PlacePieceAt(tt.args.p, tt.args.at); (err != nil) != tt.wantErr {
				t.Errorf("PlacePieceAt() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
