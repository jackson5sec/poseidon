//go:build darwin
// +build darwin

package clipboard_monitor

/*
#cgo LDFLAGS: -framework AppKit -framework Foundation
#cgo CFLAGS: -x objective-c
#include "clipboard_monitor_darwin.h"
*/
import "C"

func CheckClipboard(oldCount int) (string, error) {
	contents := C.monitorClipboard(C.long(oldCount))
	return C.GoString(contents), nil
}

func GetClipboardCount() (int, error) {
	count := C.getClipboardCount()
	return int(count), nil
}
func GetFrontmostApp() (string, error) {
	return C.GoString(C.getFrontmostApp()), nil
}
func WaitForTime() {
	C.waitForTime()
}
