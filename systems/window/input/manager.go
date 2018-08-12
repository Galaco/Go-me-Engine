package input

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/galaco/go-me-engine/engine/event"
	"github.com/galaco/go-me-engine/message/messages"
	"github.com/galaco/go-me-engine/message/messagetype"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/galaco/go-me-engine/engine/input"
)

type Manager struct {
	MouseCoordinates mgl64.Vec2
}

func (manager *Manager) Register(window *glfw.Window) {
	window.SetKeyCallback(manager.KeyCallback)
	window.SetCursorPosCallback(manager.MouseCallback)

	event.GetEventManager().Listen(messagetype.KeyDown, input.GetKeyboard())
	event.GetEventManager().Listen(messagetype.KeyReleased, input.GetKeyboard())
	event.GetEventManager().Listen(messagetype.MouseMove, input.GetMouse())
}

func (manager *Manager) Update(dt float64) {
	input.GetMouse().Update()
	glfw.PollEvents()
}

func (manager *Manager) Unregister() {

}

func (manager *Manager) KeyCallback(window *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	switch action {
	case glfw.Press:
		event.GetEventManager().Dispatch(messagetype.KeyDown, &messages.KeyDown{Key: key})
	case glfw.Repeat:
		event.GetEventManager().Dispatch(messagetype.KeyHeld, &messages.KeyHeld{Key: key})
	case glfw.Release:
		event.GetEventManager().Dispatch(messagetype.KeyReleased, &messages.KeyReleased{Key: key})
	}
}

func (manager *Manager) MouseCallback(window *glfw.Window, xpos float64, ypos float64) {
	event.GetEventManager().Dispatch(messagetype.MouseMove, &messages.MouseMove{
		X: manager.MouseCoordinates[0] - xpos,
		Y: manager.MouseCoordinates[1] - ypos,
	})
	manager.MouseCoordinates[0] = xpos
	manager.MouseCoordinates[1] = ypos
}