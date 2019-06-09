package tinymaze

import "errors"

type TinyMazeSolver struct {

}

func (solver *TinyMazeSolver) Solve(tinyMaze [][]string) ([][]string, error) {
	if len(tinyMaze) == 0 {
		return [][]string{}, errors.New("empty maze")
	}
	return [][]string{}, nil
}
