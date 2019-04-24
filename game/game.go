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
	lastUpdate time.Time //ms of last update
	lastFrame  time.Time //ms of last frame
	running    bool
}

func (g *Game) frameInterval() time.Duration {
	return time.Duration(int64(time.Second) / int64(g.fps))
}

func (g *Game) updateInterval() time.Duration {
	return time.Duration(int64(time.Second) / int64(g.ups))
}

func (g *Game) update() {
	g.updates++
}

func (g *Game) doUpdates() {
	for g.running {
		now := time.Now()

		if g.lastUpdate.Add(g.updateInterval()).Before(now) {
			g.update()
			g.lastUpdate = now
		}
	}
}

func (g Game) doRefresh(win *pixelgl.Window) {
	for g.running {
		now := time.Now()

		if g.lastFrame.Add(g.frameInterval()).Before(now) {
			win.Update()
			g.lastFrame = now
		}
	}
}

func StartGame() {
	var g = Game{
		name:       "test game",
		ups:        30,
		fps:        60,
		updates:    0,
		lastUpdate: time.Now(),
		lastFrame:  time.Now(),
		running:    true,
	}

	win := window()

	go g.doUpdates()
	go g.doRefresh(win)

	time.Sleep(5 * time.Second)
	win.SetClosed(true) // according to my calculations this should close after 5 seconds?

	log.Printf("Game %v started running at #%v per second", g.name, g.ups)
}
