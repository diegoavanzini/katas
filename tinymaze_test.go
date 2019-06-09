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

func Test_MazeOneRowOneColumns_ShouldReturnError(t *testing.T){
	maze := [][]string{
		{":S"},
	}

	solver := TinyMazeSolver{}

	_, err := solver.Solve(maze)

	if err != nil {
		assert.Equal(t, "maze should have Start and End", err.Error())
	} else {
		assert.Fail(t, "expected Start and End maze error")
	}
}


func Test_MazeWithOnlyStartAndEnd(t *testing.T){
	maze := [][]string{
		{":S"},
		{":E"},
	}

	solver := TinyMazeSolver{}

	solvedMaze, err := solver.Solve(maze)

	if err != nil {
		t.Error(err)
	}

	expected := Maze{{":x"}, {":x"}}
	assert.Equal(t, expected ,solvedMaze)
}

func Test_MazeWithOnlyStartAndEndInOneRow(t *testing.T){
	maze := [][]string{
		{":S", ":E"},
	}

	solver := TinyMazeSolver{}

	solvedMaze, err := solver.Solve(maze)

	if err != nil {
		t.Error(err)
	}

	expected := Maze{{":x", ":x"}}
	assert.Equal(t, expected ,solvedMaze)
}

func Test_MazeWithStartEndAndAWallBetween_ShouldReturnError(t *testing.T){
	maze := [][]string{
		{":S", "1", ":E"},
	}

	solver := TinyMazeSolver{}

	_, err := solver.Solve(maze)

	if err != nil {
		assert.Equal(t, "maze without solution", err.Error())
	} else {
		assert.Fail(t, "expected without error")
	}
}
