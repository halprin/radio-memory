package windows

import (
	"log"

	"gioui.org/widget"

	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/widget/material"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/op"
)

type MemoryList struct {
	window     *app.Window
	operations op.Ops
	theme      *material.Theme
	button     *widget.Clickable
}

func (receiver MemoryList) Display() {
	log.Println("Starting the memory list window")

	receiver.window = app.NewWindow(app.Title("Memories"))
	receiver.theme = material.NewTheme(gofont.Collection())
	receiver.button = &widget.Clickable{}

	go receiver.eventLoop()
}

func (receiver MemoryList) eventLoop() {
	for event := range receiver.window.Events() {
		switch event := event.(type) {
		case system.FrameEvent:
			receiver.draw(event)
		case system.DestroyEvent:
			log.Println("Destroying the memory list window")
			if event.Err != nil {
				log.Fatal(event.Err)
			}
			return
		}
	}
}

func (receiver MemoryList) draw(event system.FrameEvent) {
	context := layout.NewContext(&receiver.operations, event)

	layout.Flex{
		Axis:    layout.Vertical,
		Spacing: layout.SpaceEnd,
	}.Layout(context, layout.Rigid(material.Button(receiver.theme, receiver.button, "Moof!").Layout))

	event.Frame(context.Ops)
}
