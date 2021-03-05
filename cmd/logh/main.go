package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/xgbutil"
	"github.com/wiedzmin/loggerhead/impl"
)

func main() {
	X, err := xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	idle, err := impl.GetXIdleTime(X)
	if err != nil {
		fmt.Println("error getting idle time")
	}
	fmt.Printf("idle time: %d\n", idle)

	active, err := impl.GetActiveWindowName(X)
	if err != nil {
		fmt.Println("error getting active window")
	}
	fmt.Printf("active window: %s\n", *active)

	desktop, err := impl.GetActiveDesktopMeta(X)
	if err != nil {
		fmt.Println("error getting active desktop meta")
	}
	fmt.Printf("active desktop: %d / %s\n", desktop.Index, desktop.Name)
	fmt.Println("==============================")

	windowNames, err := impl.GetWindowNames(X)
	if err != nil {
		fmt.Println("error getting window names")
	}
	for _, name := range windowNames {
		fmt.Printf("%s\n", name)
	}
}
