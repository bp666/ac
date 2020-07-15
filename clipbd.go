package ac

import (
	"syscall"
	"unsafe"
)

// Copy data to clipboard
func Copy(text string) {
	myOpenClipboard()
	defer myCloseClipboard()

	myEmptyClipboard()
	data := syscall.StringToUTF16(text)

	pm := myGlobalAlloc(len(data) * int(unsafe.Sizeof(data[0])))
	defer myGlobalFree(pm)

	block := myGlobalLock(pm)

	myLstrcpyW(block, unsafe.Pointer(&data[0]))

	myGlobalUnlock(pm)

	mySetClipboardData(pm)
}
