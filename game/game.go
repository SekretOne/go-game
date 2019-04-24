package game

import (
	"github.com/faiface/pixel/pixelgl"
	"log"
	"time"
)

type Game struct {
	name       string
	ups        int //updates per second
	updates    int
	fps        int
	lastUpdate int64 //ms of last update
	lastFrame  int64
}

func (g Game) update() {
	g.updates++
}

func doUpdates(win *pixelgl.Window) {
	for !win.Closed() {
		win.Update()
	}
}

func StartGame() {
	var g = Game{
		name:       "test game",
		ups:        30,
		fps:        60,
		updates:    0,
		lastUpdate: time.Now().UnixNano(),
		lastFrame:  time.Now().UnixNano(),
	}

	win := window()

	go doUpdates(win)

	time.Sleep(5 * time.Second)
	win.SetClosed(true) // according to my calculations this should close after 5 seconds?

	log.Printf("Game %v started running at #%v per second", g.name, g.ups)
}
