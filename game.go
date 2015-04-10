package tictacgo

type Game struct {
  players [2]*Player
}

func (g *Game) initPlayer() {
  g.players[0] = new(Player)
}
