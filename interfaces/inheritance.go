package main

import "fmt"

type OneMan interface {
	Play() string
	Manage() string
}

type PlayingManger struct {
	*Player
}

func (plm *PlayingManger) Manage() string {
	return "manage"
}

type Player struct {}

func (ply *Player) Play() string {
	return "play"
}

func main() {
	s := new(PlayingManger)
	fmt.Println(s.Play())

	playAndManage(s)
}

func playAndManage(o OneMan) {
	fmt.Println("Play and manage!!")
	fmt.Println(o.Play())
	fmt.Println(o.Manage())
}
