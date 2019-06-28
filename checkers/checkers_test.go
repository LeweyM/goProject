package checkers

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)



func TestBoard_play(t *testing.T) {
	type args struct {
		col    int
		row    int
		player int
	}
	tests := []struct {
		name  string
		board *Board
		args  args
		expectedBoard *Board
		expectedVal bool
	}{
		{"should change the corner cell",
			&Board{ [3][3]int{} },
			args{0, 0, 1},
			&Board{[3][3]int{ {1, 0, 0},{0, 0, 0}, {0, 0 ,0} } },
			true,
		},
		{"should change the second cell",
			&Board{ [3][3]int{} },
			args{0, 1, 1},
			&Board{[3][3]int{ {0, 1, 0},{0, 0, 0}, {0, 0 ,0} } },
			true,
		},
		{"should change the second cell to 2",
			&Board{ [3][3]int{} },
			args{0, 1, 2},
			&Board{[3][3]int{ {0, 2, 0},{0, 0, 0}, {0, 0 ,0} } },
			true,
		},
		{"should change the second row cell to 1",
			&Board{ [3][3]int{} },
			args{1, 0, 1},
			&Board{[3][3]int{ {0, 0, 0},{1, 0, 0}, {0, 0 ,0} } },
			true,
		},
		{"should not an already written cell",
			&Board{ [3][3]int{ {1} } },
			args{0, 0, 2},
			&Board{[3][3]int{ {1, 0, 0},{0, 0, 0}, {0, 0 ,0} } },
			false,
		},
		{"should return false if attempting to overwrite",
			&Board{ [3][3]int{ {1} } },
			args{0, 0, 2},
			&Board{[3][3]int{ {1, 0, 0},{0, 0, 0}, {0, 0 ,0} } },
			false,
		},
	}

	RegisterFailHandler(Fail)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val := tt.board.play(tt.args.col, tt.args.row, tt.args.player)

			Ω(tt.board.cells).Should(Equal(tt.expectedBoard.cells))
			Ω(val).Should(Equal(tt.expectedVal))
		})
	}
}
