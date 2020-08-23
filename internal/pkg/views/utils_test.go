package views

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FixedPositionWithinTerminalBounds(t *testing.T) {
	// Simplest case, all dimensions are fixed values and within the terminal bounds
	WithTerminalSize(20, 40).
		AndView(5, 10, 10, 20).
		ExpectBounds(t, 5, 10, 15, 30)
}

func Test_FixedPositionOverflowingTerminalBounds(t *testing.T) {
	// All dimensions are fixed values and but would extend beyond the terminal bounds
	WithTerminalSize(10, 20).
		AndView(5, 10, 10, 20).
		ExpectBounds(t, 5, 10, 9, 19)
}

func Test_RelativeWidthAndHeight(t *testing.T) {
	// Fixed origin, relative width and height. Sufficient space for non-zero size
	WithTerminalSize(20, 40).
		AndView(5, 10, -10, -20).
		ExpectBounds(t, 5, 10, 10, 20)
}

func Test_RelativeWidthAndHeightOverflowingTerminalBounds(t *testing.T) {
	// Fixed origin, relative width and height. Would overflow so is shifted to smallest non-zero width that fits the bounds
	WithTerminalSize(10, 20).
		AndView(5, 10, -7, -12).
		ExpectBounds(t, 1, 6, 3, 8) // origin is shunted to ensure that the bounds don't describe a zero-width/height
}

func Test_RelativePosition(t *testing.T) {
	// Fixed origin, relative width and height. Sufficient space for non-zero size
	WithTerminalSize(20, 40).
		AndView(-10, -20, 5, 10).
		ExpectBounds(t, 10, 20, 15, 30)
}

func Test_RelativePositionOverflowingTerminalBounds(t *testing.T) {
	// Fixed origin, relative width and height. SWould overflow so is shift
	WithTerminalSize(10, 20).
		AndView(-10, -20, 5, 10).
		ExpectBounds(t, 0, 0, 5, 10)
}

type ViewTestData struct {
	terminalWidth  int
	terminalHeight int
	viewX          int
	viewY          int
	viewWidth      int
	viewHeight     int
}

func (m *ViewTestData) Size() (w int, h int) {
	return m.terminalWidth, m.terminalHeight
}

func WithTerminalSize(w int, h int) *ViewTestData {
	m := ViewTestData{
		terminalWidth:  w,
		terminalHeight: h,
	}
	return &m
}

func (m *ViewTestData) AndView(x int, y int, w int, h int) *ViewTestData {
	m.viewX = x
	m.viewY = y
	m.viewWidth = w
	m.viewHeight = h
	return m
}

func (m *ViewTestData) ExpectBounds(t *testing.T, x0 int, y0 int, x1 int, y1 int) {

	actualX0, actualY0, actualX1, actualY1 := getViewBounds(m, m.viewX, m.viewY, m.viewWidth, m.viewHeight)

	assert.Equal(t, x0, actualX0, "validate x0")
	assert.Equal(t, y0, actualY0, "validate y0")
	assert.Equal(t, x1, actualX1, "validate x1")
	assert.Equal(t, y1, actualY1, "validate y1")
}
