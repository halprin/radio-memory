package windows

import (
	"fmt"
	"log"

	"github.com/halprin/radio-memory/radio"

	"gioui.org/widget"

	"gioui.org/font/gofont"
	"gioui.org/layout"
	"gioui.org/widget/material"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/op"
)

type MemoryList struct {
	radio    radio.Radio
	memories []radio.Memory

	window       *app.Window
	operations   op.Ops
	theme        *material.Theme
	reloadButton *widget.Clickable
	writeButton  *widget.Clickable
	list         *widget.List
}

func (receiver MemoryList) Display() {
	log.Println("Starting the memory list window")
	receiver.radio = radio.YaesuFtm500D{SdCardMemoryPath: "/Users/halprin/Desktop/MEMFTM500D.dat"}

	var err error
	receiver.memories, err = receiver.radio.ReadMemories()
	if err != nil {
		log.Fatalf("Couldn't read the memories!  %s", err.Error())
	}

	receiver.window = app.NewWindow(app.Title("Memories"))
	receiver.theme = material.NewTheme(gofont.Collection())
	receiver.reloadButton = &widget.Clickable{}
	receiver.writeButton = &widget.Clickable{}
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
		layout.Rigid(material.Button(receiver.theme, receiver.reloadButton, "Reload").Layout),
		layout.Rigid(material.Button(receiver.theme, receiver.writeButton, "Write").Layout),
		layout.Rigid(func(context layout.Context) layout.Dimensions {
			themedList := material.List(receiver.theme, receiver.list)
			return themedList.Layout(context, len(receiver.memories), func(context layout.Context, index int) layout.Dimensions {
				listItem := material.H3(receiver.theme, fmt.Sprintf("%f", receiver.memories[index].FrequencyRx))
				return listItem.Layout(context)
			})
		}))

	event.Frame(context.Ops)
}
