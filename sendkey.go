package ac

import (
	"syscall"
	"unsafe"
)

// SendKey is send a string to the foreground
func SendKey(str string) {
	s := syscall.StringToUTF16(str)
	for _, c := range s {
		sendChar(c)
	}
}

func sendChar(char uint16) {
	var input tagKeyBdInput

	input.tp = inputKeyBoard
	input.ki.wScan = char

	input.ki.dwFlags = keyEventFUnicode
	mySendInput(unsafe.Pointer(&input), int32(unsafe.Sizeof(input)))

	input.ki.dwFlags = keyEventFUnicode | keyEventFKeyUp
	mySendInput(unsafe.Pointer(&input), int32(unsafe.Sizeof(input)))
}
