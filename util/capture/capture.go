package capture

import (
	"github.com/BurntSushi/xgb/screensaver"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/icccm"
)

func GetXIdleTime(x *xgbutil.XUtil) (uint32, error) {
	screensaver.Init(x.Conn())
	info, err := screensaver.QueryInfo(x.Conn(), xproto.Drawable(x.RootWin())).Reply()
	if err != nil {
		return 0, err
	}
	return info.MsSinceUserInput, nil
}

func getWindowName(x *xgbutil.XUtil, xid xproto.Window) (*string, error) {
	name, err := ewmh.WmNameGet(x, xid)
	if err != nil || len(name) == 0 {
		name, err = icccm.WmNameGet(x, xid)
		if err != nil || len(name) == 0 { // If we still can't find anything, give up.
			return nil, err
		}
	}
	return &name, nil
}

func GetWindowNames(x *xgbutil.XUtil) ([]string, error) {
	var result []string
	clientids, err := ewmh.ClientListGet(x)
	if err != nil {
		return nil, err
	}
	for _, clientid := range clientids {
		name, err := getWindowName(x, clientid)
		if err != nil {
			return nil, err
		}
		result = append(result, *name)
	}
	return result, nil
}

func GetActiveWindowName(x *xgbutil.XUtil) (*string, error) {
	active, err := ewmh.ActiveWindowGet(x)
	if err != nil {
		return nil, err
	}
	name, err := getWindowName(x, active)
	if err != nil {
		return nil, err
	}
	return name, nil
}

func GetActiveDesktopMeta(x *xgbutil.XUtil) (*Desktop, error) {
	workspace, err := ewmh.CurrentDesktopGet(x)
	if err != nil {
		return nil, err
	}

	workspaceNames, err := ewmh.DesktopNamesGet(x)
	if err != nil {
		return nil, err
	}

	return &Desktop{workspace, workspaceNames[workspace]}, nil
}
