package tictacgo

import "testing"

func TestDisplay (t *testing.T) {
  var board Board
  boardDrawing := "__|__|__\n__|__|__\n  |  |  "

  if board.Display() != boardDrawing {
    t.Fail() 
  }
}
