package engine

import (
	"github.com/bradhe/stopwatch"
	"github.com/galaco/go-me-engine/engine/event"
	"github.com/galaco/go-me-engine/engine/interfaces"
	"runtime"
)

// Game engine
// Only 1 can be initialised
type engine struct {
	EventManager event.Manager
	Managers     []interfaces.IManager
	Running      bool

	entities   []interfaces.IEntity
	components []interfaces.IComponent
}

// Initialise the engine, and attached managers
func (engine *engine) Initialise() {

	for _, manager := range engine.Managers {
		manager.Register()
	}

}

// Run the engine
func (engine *engine) Run() {
	engine.Running = true

	// Begin the event manager thread in the background
	event.GetEventManager().RunConcurrent()
	// Anything that runs concurrently can start now
	for _, manager := range engine.Managers {
		manager.RunConcurrent()
	}

	dt := 0.0
	timer := stopwatch.Start()

	for engine.Running == true {
		dt = 1 / float64(timer.Milliseconds())
		//if dt < FRAMERATE {
		//	continue
		//}
		// Restart timer now so we can record loop time
		timer.Stop()
		timer = stopwatch.Start()

		for _, manager := range engine.Managers {
			manager.Update(dt)
		}

		for _, manager := range engine.Managers {
			manager.PostUpdate()
		}
	}
}

// Add a new manager to the engine
func (engine *engine) AddManager(manager interfaces.IManager) {
	engine.Managers = append(engine.Managers, manager)
}

func NewEngine() *engine {
	runtime.LockOSThread()
	return &engine{}
}
