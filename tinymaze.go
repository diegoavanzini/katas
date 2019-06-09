package tinymaze

import (
	"errors"
)

type TinyMazeSolver struct {

}

const wall = "1"
const start = ":S"
const end = ":E"

type Maze [][]string

func (solver *TinyMazeSolver) Solve(tinyMaze Maze) (Maze, error) {
	if len(tinyMaze) == 0 {
		return nil, errors.New("empty maze")
	}
	if tinyMaze.findStartAndEnd() < 2  {
		return nil, errors.New("maze should have Start and End")
	}
	var solvedMaze Maze
	solvedMaze = tinyMaze
	for rowNum, _ := range solvedMaze {
		err := solvedMaze.scanRow(rowNum)
		if err != nil {
			return nil, err
		}
	}

	return tinyMaze, nil
}

func (m *Maze) scanRow(rowNum int) error {
	foundWalls := 0
	row := (*m)[rowNum]
	for colNum, value := range row {
		if value == wall || value == start || value == end {
			foundWalls++
		}
		if value != wall {
			m.markCell(rowNum, colNum)
		}
	}
	if len(row) == foundWalls && foundWalls > 2 {
		return errors.New("maze without solution")
	}
	return nil
}

func (m *Maze) markCell(rowNum int, colNum int) {
	(*m)[rowNum][colNum] = ":x"
}

func (m *Maze) findStartAndEnd() int {
	found := 0
	for _, row := range *m {
		for _, value := range row {
			if value == start || value == end {
				found++
			}
		}
	}
	return found
}
