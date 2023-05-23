package windows

import (
	"gioui.org/app"
	"gioui.org/io/system"
	"log"
)

type MemoryList struct {
	window *app.Window
}

func (receiver MemoryList) Display() {
	log.Println("Starting the memory list window")

	receiver.window = app.NewWindow(app.Title("Memories"))

	go receiver.eventLoop()
}

func (receiver MemoryList) eventLoop() {
	for event := range receiver.window.Events() {
		switch event := event.(type) {
		case system.FrameEvent:
			receiver.draw()
			case system.DestroyEvent:
				log.Println("Destroying the memory list window")
				if event.Err != nil {
					log.Fatal(event.Err)
				}
				return
		}
	}
}

func (receiver MemoryList) draw() {

}
