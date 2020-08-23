package views

type Measurable interface {
	Size() (x int, y int)
}

func getViewBounds(g Measurable, x int, y int, width int, height int) (x0, y0, x1, y1 int) {
	x0 = x
	y0 = y
	sx, sy := g.Size()
	if x0 < 0 {
		x0 += sx
	}
	if y0 < 0 {
		y0 += sy
	}
	if width <= 0 {
		x1 = width + sx // width relative to terminal width
	} else {
		x1 = x0 + width // fixed width
	}
	if height <= 0 {
		y1 = height + sy // height relative to terminal height
	} else {
		y1 = y0 + height // fixed height
	}

	if x1 >= sx {
		x1 = sx - 1 // don't allow x1 to exceed zero-based terminal width
	}
	if y1 >= sy {
		y1 = sy - 1 // don't allow y1 to exceed zero-based  terminal height
	}
	if x1-x0 < 2 {
		x0 = x1 - 2 // ensure we don't end up with zero width view
	}
	if y1-y0 < 2 {
		y0 = y1 - 2 // ensure we don't end up with zero height view
	}
	return
}
