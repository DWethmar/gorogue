package console

import (
	"fmt"
	"math"

	"github.com/DWethmar/gorogue/color"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var tileSize = 50

type Console struct {
	Title  string
	Height int
	Width  int
	x      int
	y      int
	buffer [][]Cell
}

func New(width, height int, title string) (*Console, error) {
	buf := make([][]Cell, width)
	for x := range buf {
		buf[x] = make([]Cell, height)
		for y := range buf[x] {
			var col = color.Green
			if (x+y)%2 == 0 {
				col = color.Red
			}
			buf[x][y] = Cell{Background: col}
		}
	}

	return &Console{
		Title:  title,
		Height: height,
		Width:  width,
		buffer: buf,
	}, nil
}

func (c *Console) elapsedFPS() float64 {
	e := 1.0 / math.Min(float64(ebiten.FPS), ebiten.CurrentFPS())
	if e > math.MaxFloat64 {
		e = 0
	}
	return e
}

func (con *Console) draw(screen *ebiten.Image, timeElapsed float64, offsetX, offsetY int) {

	for x := range con.buffer {

		for y := range con.buffer[x] {

			fmt.Printf("%f x: %d y: %d color: %+v\n", timeElapsed, x, y, con.buffer[x][y].Background)

			if con.buffer[x][y].Background.A == 0 {
				continue
			}

			fmt.Printf("%f\n", float64((offsetX+con.x+x)*tileSize))

			ebitenutil.DrawRect(
				screen,
				float64((offsetX+con.x+x)*tileSize),
				float64((offsetY+con.y+y)*tileSize),
				float64(tileSize),
				float64(tileSize),
				con.buffer[x][y].Background,
			)
		}
	}
}

func (con *Console) update(screen *ebiten.Image) error {
	timeElapsed := con.elapsedFPS()
	con.draw(screen, timeElapsed, 0, 0)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("FPS: %0.2f", ebiten.CurrentFPS()))

	return nil
}

// Start will open the console window with the given scale.
func (c *Console) Start(scale float64) error {
	return ebiten.Run(c.update, c.Width*tileSize, c.Height*tileSize, scale, c.Title)
}
