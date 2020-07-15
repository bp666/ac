package ac

import (
	"strings"
	"unsafe"
)

type tagKeyBdInput struct {
	tp uint32
	ki keyBdInput
}

type keyBdInput struct {
	wVk         uint16
	wScan       uint16
	dwFlags     uint32
	time        uint32
	dwExtraInfo uintptr
	Unused      [8]byte
}

const (
	keyEventFExtendedKey = 0x0001
	KeyEventFKeyUp       = 0x0002
	KeyEventFScanCode    = 0x0008
	KeyEventFUnicode     = 0x0004
)

var (
	kdownInput tagKeyBdInput
	kupInput   tagKeyBdInput

	keysMap map[string]int
)

func init() {
	kdownInput.tp = inputKeyBoard
	kupInput.tp = inputKeyBoard
	kupInput.ki.dwFlags = KeyEventFKeyUp

	keysMap = map[string]int{
		"back":  0x08,
		"tab":   0x09,
		"enter": 0x0D,
		"shift": 0x10,
		"ctrl":  0x11,
		"alt":   0x12,
		"prior": 0x21, // page up
		"next":  0x22, // page down
		"end":   0x23,
		"home":  0x24,
		"left":  0x25,
		"up":    0x26,
		"right": 0x27,
		"down":  0x28,
		"del":   0x2E,
		"f1":    0x70,
		"f2":    0x71,
		"f3":    0x72,
		"f4":    0x73,
		"f5":    0x74,
		"f6":    0x75,
		"f7":    0x76,
		"f8":    0x77,
		"f9":    0x78,
		"f10":   0x79,
		"f11":   0x7A,
		"f12":   0x7B,
	}
}

func keyDown(wvk uint16) {
	kdownInput.ki.wVk = wvk
	mySendInput(unsafe.Pointer(&kdownInput), int32(unsafe.Sizeof(kdownInput)))
}

func keyUp(wvk uint16) {
	kupInput.ki.wVk = wvk
	mySendInput(unsafe.Pointer(&kupInput), int32(unsafe.Sizeof(kupInput)))
}

// Key is simulate a single key
func Key(key string) {
	var code uint16

	switch len(key) {
	case 0:
		return
	case 1:
		// ascii
		code = uint16([]byte(key)[0])
		break
	default:
		key = strings.ToLower(key)
		value, ok := keysMap[key]
		if ok {
			code = uint16(value)
		} else {
			return
		}
	}
	keyDown(code)
	keyUp(code)
}

// HotKey is simulate multiple keys(Up to two keys).
// The special key should be the first,
// eg: ctrl+V(√)	V+ctrl(×).
func HotKey(keys ...string) {
	switch len(keys) {
	case 0:
		return
	case 1:
		Key(keys[0])
		return
	case 2:
		value, ok := keysMap[keys[0]]
		if ok {
			first := uint16(value)

			if len(keys[1]) == 1 {
				temp := strings.ToUpper(keys[1])
				second := uint16([]byte(temp)[0])
				keyDown(first)
				keyDown(second)
				keyUp(second)
				keyUp(first)
			}
		} else {
			return
		}
		break
	}
}
