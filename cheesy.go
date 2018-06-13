package main

import (
	"fmt"
	"gui"
)

var (
	appGui = gui.Gui{}
)

func main() {
	appGui.MakeUI()
	fmt.Println("init !")
}
