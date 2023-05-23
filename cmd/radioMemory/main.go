package main

import (
	"gioui.org/app"
	"github.com/halprin/radio-memory/windows"
)

func main() {
	windows.MemoryList{}.Display()
	app.Main()
}
