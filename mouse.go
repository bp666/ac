package ac

import "unsafe"

// Point record x, y coordinates
type Point struct {
	x, y int32
}

type tagMouseInput struct {
	tp uint32
	mi mouseInput
}

type mouseInput struct {
	dx          int32
	dy          int32
	mouseData   int32
	dwFlags     uint32
	time        uint32
	dwExtraInfo uintptr
}

const (
	mouseEventFAbsolute       = 0x8000
	mouseEventFMove           = 0x0001
	mouseEventFMoveNocoalesce = 0x2000
	mouseEventFLeftDown       = 0x0002
	mouseEventFLeftUp         = 0x0004
	mouseEventFRightDown      = 0x0008
	mouseEventFRightUp        = 0x0010
	mouseEventFMiddleDown     = 0x0020
	mouseEventFMiddleUp       = 0x0040
	mouseEventFVirtualDesk    = 0x4000
	mouseEventFWheel          = 0x0800
	mouseEventFHwheel         = 0x1000
	mouseEventFXDown          = 0x0080
	mouseEventFXUp            = 0x0100
)

var (
	mdownInput tagMouseInput
	mupInput   tagMouseInput
)

// GetMousePos reutrn point(x, y)
func GetMousePos() Point {
	var pt Point
	mygetCursorPos(unsafe.Pointer(&pt))
	return pt
}

func setMousePos(x, y int32) {
	mySetCursorPos(x, y)
}

// Click is left mouse click
func Click(x, y int32) {
	setMousePos(x, y)
	mdownInput.mi.dwFlags = mouseEventFLeftDown
	mupInput.mi.dwFlags = mouseEventFLeftUp
	mySendInput(unsafe.Pointer(&mdownInput), int32(unsafe.Sizeof(mdownInput)))
	mySendInput(unsafe.Pointer(&mupInput), int32(unsafe.Sizeof(mupInput)))
}

// DClick is left mouse double-click
func DClick(x, y int32) {
	Click(x, y)
	Click(x, y)
}

// RClick is right mouse click
func RClick(x, y int32) {
	setMousePos(x, y)
	mdownInput.mi.dwFlags = mouseEventFRightDown
	mupInput.mi.dwFlags = mouseEventFRightUp
	mySendInput(unsafe.Pointer(&mdownInput), int32(unsafe.Sizeof(mdownInput)))
	mySendInput(unsafe.Pointer(&mupInput), int32(unsafe.Sizeof(mupInput)))
}

// MClick is middle mouse click
func MClick(x, y int32) {
	setMousePos(x, y)
	mdownInput.mi.dwFlags = mouseEventFMiddleDown
	mupInput.mi.dwFlags = mouseEventFMiddleUp
	mySendInput(unsafe.Pointer(&mdownInput), int32(unsafe.Sizeof(mdownInput)))
	mySendInput(unsafe.Pointer(&mupInput), int32(unsafe.Sizeof(mupInput)))
}

// Scroll is vertical scroll mouse.
// sign to down,
// unsign to up
func Scroll(value int32) {
	var input tagMouseInput
	input.mi.dwFlags = mouseEventFWheel
	input.mi.mouseData = value
	mySendInput(unsafe.Pointer(&input), int32(unsafe.Sizeof(input)))
}

// HScroll is horizontal scroll mouse.
// sign to left,
// unsign to right
func HScroll(value int32) {
	var input tagMouseInput
	input.mi.dwFlags = mouseEventFHwheel
	input.mi.mouseData = value
	mySendInput(unsafe.Pointer(&input), int32(unsafe.Sizeof(input)))
}
