package tictacgo

import "testing"

func TestNew(t *testing.T) {
  var game Game
  game.initPlayer()

  if game.players[0] == nil {
    t.Fail()
  }
}
