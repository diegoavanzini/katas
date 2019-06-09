package tinymaze

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func Test_WithEmptyMaze_ShouldReturnError(t *testing.T) {

	maze := [][]string{}

	solver := TinyMazeSolver{}

	_, err := solver.Solve(maze)

	if err != nil {
		assert.Equal(t, "empty maze", err.Error())
	} else {
		assert.Fail(t, "expected empty maze error")
	}

}
