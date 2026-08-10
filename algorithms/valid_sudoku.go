package algorithms

func IsValidSudoku(board [][]byte) bool {
	rows := make([]map[byte]bool, 9)
	cols := make([]map[byte]bool, 9)
	boxes := make([]map[byte]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}
	for i, row := range board {
		for j, cell := range row {
			if cell == '.' {
				continue
			}
			boxindex := (i/3)*3 + (j/3)
			if _, ok := rows[i][cell]; ok{
				return false
			}
			if _, ok := cols[j][cell]; ok {
				return false
			}
			if _, ok := boxes[boxindex][cell]; ok {
				return false
			}
			rows[i][cell] = true
			cols[j][cell] = true
			boxes[boxindex][cell] = true
		}
	}
	return true
}