package checkers

type Board struct {
	cells [3][3]int
}

func (board *Board) play(col int, row int, player int) bool {
	v := board.cells[col][row]
	if v == 0 {
		board.cells[col][row] = player
	}
	return v == 0
}
