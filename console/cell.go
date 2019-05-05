package console

import "github.com/DWethmar/gorogue/color"

// Cell represents a cell in the console
type Cell struct {
	Foreground color.Color
	Background color.Color

	Char int
}

var emptyCell = Cell{
	Foreground: color.Black,
}

var redCell = Cell{
	Foreground: color.Red,
}
