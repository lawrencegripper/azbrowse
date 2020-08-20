package views

import "github.com/stuartleeks/gocui"

func getViewBounds(g *gocui.Gui, x int, y int, width int, height int) (x0, y0, x1, y1 int) {
	x0 = x
	y0 = y
	sx, sy := g.Size()
	sx -= 2
	sy -= 2
	if x0 < 0 {
		x0 += sx
	}
	if y0 < 0 {
		y0 += sy
	}
	if width <= 0 {
		x1 = width + sx + x0 // width relative to terminal width
	} else {
		x1 = x0 + width // fixed width
	}
	if height <= 0 {
		y1 = height + sy + y0 // height relative to terminal height
	} else {
		y1 = y0 + height // fixed height
	}

	if x1 > sx {
		x1 = sx // don't allow x1 to exceed terminal width
	}
	if x1 <= x0 {
		x0 = x1 - 2 // ensure we don't end up with zero width view
	}
	if y1 <= y0 {
		y0 = y1 - 2 // ensure we don't end up with zero height view
	}
	return
}
