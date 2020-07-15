package ac

import (
	"syscall"
	"unsafe"
)

// INPU TYPE:
// 0 mouse
// 1 keyboard
// 2 hardware
const (
	inputMouse    = 0
	inputKeyBoard = 1
	inputHardware = 2
)

const (
	gMemMoveable  = 2
	cfUnicodeText = 13
)

var (
	libuser32        = syscall.NewLazyDLL("user32.dll")
	sendInput        = libuser32.NewProc("SendInput")
	setCursorPos     = libuser32.NewProc("SetCursorPos")
	getCursorPos     = libuser32.NewProc("GetCursorPos")
	openClipboard    = libuser32.NewProc("OpenClipboard")
	emptyClipboard   = libuser32.NewProc("EmptyClipboard")
	setClipboardData = libuser32.NewProc("SetClipboardData")
	getClipboardData = libuser32.NewProc("GetClipboardData")
	closeClipboard   = libuser32.NewProc("CloseClipboard")

	kernel32     = syscall.NewLazyDLL("kernel32.dll")
	globalAlloc  = kernel32.NewProc("GlobalAlloc")
	globalFree   = kernel32.NewProc("GlobalFree")
	globalLock   = kernel32.NewProc("GlobalLock")
	globalUnlock = kernel32.NewProc("GlobalUnlock")
	lstrcpyW     = kernel32.NewProc("lstrcpyW")
)

func mySendInput(pInputs unsafe.Pointer, cbSize int32) {
	sendInput.Call(uintptr(1), uintptr(pInputs), uintptr(cbSize))
}

func mySetCursorPos(x, y int32) {
	setCursorPos.Call(uintptr(x), uintptr(y))
}

func mygetCursorPos(lpPoint unsafe.Pointer) {
	getCursorPos.Call(uintptr(lpPoint))
}

func myOpenClipboard() {
	openClipboard.Call(0)
}

func myEmptyClipboard() {
	emptyClipboard.Call()
}

func mySetClipboardData(handle uintptr) {
	setClipboardData.Call(uintptr(cfUnicodeText), handle)
}

func myGetClipboardData() uintptr {
	handle, _, _ := getClipboardData.Call(uintptr(cfUnicodeText))
	return handle
}

func myCloseClipboard() {
	openClipboard.Call()
}

func myGlobalAlloc(dataSize int) uintptr {
	pMem, _, _ := globalAlloc.Call(gMemMoveable, uintptr(dataSize))
	return pMem
}

func myGlobalFree(pMem uintptr) {
	globalFree.Call(pMem)
}

func myGlobalLock(pMem uintptr) uintptr {
	pMemBlock, _, _ := globalLock.Call(pMem)
	return pMemBlock
}

func myGlobalUnlock(pMem uintptr) {
	globalUnlock.Call(pMem)
}

func myLstrcpyW(pWstr uintptr, pCWstr unsafe.Pointer) {
	lstrcpyW.Call(pWstr, uintptr(pCWstr))
}
