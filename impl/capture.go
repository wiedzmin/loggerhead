package impl

import (
	"fmt"

	"github.com/BurntSushi/xgbutil"
	"github.com/wiedzmin/loggerhead/util/capture"
)

type Entry struct {
	ActiveWindow  string
	ActiveDesktop capture.Desktop
	Windows       []string
	IdleTime      uint32
}

func (e *Entry) Print() {
	fmt.Printf("idle time: %d\n", e.IdleTime)
	fmt.Printf("active window: %s\n", e.ActiveWindow)
	fmt.Printf("active desktop: %d / %s\n", e.ActiveDesktop.Index, e.ActiveDesktop.Name)
	fmt.Println("==============================")
	for _, name := range e.Windows {
		fmt.Printf("%s\n", name)
	}
}

func Capture() (*Entry, error) {
	var result Entry
	X, err := xgbutil.NewConn()
	if err != nil {
		return nil, err
	}

	idle, err := capture.GetXIdleTime(X)
	if err != nil {
		return nil, err
	}
	result.IdleTime = idle

	active, err := capture.GetActiveWindowName(X)
	if err != nil {
		return nil, err
	}
	result.ActiveWindow = *active

	desktop, err := capture.GetActiveDesktopMeta(X)
	if err != nil {
		return nil, err
	}
	result.ActiveDesktop = *desktop

	windowNames, err := capture.GetWindowNames(X)
	if err != nil {
		return nil, err
	}
	result.Windows = windowNames
	return &result, nil
}
