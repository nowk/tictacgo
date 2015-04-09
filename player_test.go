package tictacgo

import "testing"

func TestMark(t *testing.T) {
  player := Player{mark: "X"}
  
  if player.mark != "X" {
    t.Fail()
  }
}
