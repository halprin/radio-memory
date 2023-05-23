package windows

import (
	"fmt"
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
	list       *widget.List
}

func (receiver MemoryList) Display() {
	log.Println("Starting the memory list window")

	receiver.window = app.NewWindow(app.Title("Memories"))
	receiver.theme = material.NewTheme(gofont.Collection())
	receiver.button = &widget.Clickable{}
	receiver.list = &widget.List{
		List: layout.List{
			Axis: layout.Vertical,
		},
	}

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
	}.Layout(context,
		layout.Rigid(material.Button(receiver.theme, receiver.button, "Moof!").Layout),
		layout.Rigid(func(context layout.Context) layout.Dimensions {
			themedList := material.List(receiver.theme, receiver.list)
			return themedList.Layout(context, 10, func(context layout.Context, index int) layout.Dimensions {
				listItem := material.H1(receiver.theme, fmt.Sprintf("List item %d", index))
				return listItem.Layout(context)
			})
		}))

	event.Frame(context.Ops)
}
