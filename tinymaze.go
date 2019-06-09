package tinymaze

import "errors"

type TinyMazeSolver struct {

}

type Maze [][]string

func (solver *TinyMazeSolver) Solve(tinyMaze Maze) (Maze, error) {
	if len(tinyMaze) == 0 {
		return [][]string{}, errors.New("empty maze")
	}
	if len(tinyMaze) == 1 && len(tinyMaze[0]) == 1 {
		return [][]string{}, errors.New("maze should have Start and End")
	}
	var solvedMaze Maze
	solvedMaze = tinyMaze
	for rowNum, _ := range solvedMaze {
		solvedMaze.scanRow(rowNum)
	}

	return tinyMaze, nil
}

func (m *Maze) scanRow(rowNum int) {
	for colNum, _ := range (*m)[rowNum] {
		m.markCell(rowNum, colNum)
	}
}

func (m *Maze) markCell(rowNum int, colNum int) {
	(*m)[rowNum][colNum] = ":x"
}
