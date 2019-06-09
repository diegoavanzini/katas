package tinymaze

import "errors"

type TinyMazeSolver struct {

}

func (solver *TinyMazeSolver) Solve(tinyMaze [][]string) ([][]string, error) {
	if len(tinyMaze) == 0 {
		return [][]string{}, errors.New("empty maze")
	}
	if len(tinyMaze) == 1 && len(tinyMaze[0]) == 1 {
		return [][]string{}, errors.New("maze should have Start and End")
	}
	return [][]string{}, nil
}
