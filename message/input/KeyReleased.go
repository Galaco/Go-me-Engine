package input

import (
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/galaco/bsp-viewer/engine/event"
)

type KeyReleased struct {
	event.Message
	Key glfw.Key
}
