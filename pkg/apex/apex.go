package apex

import (
	"errors"
	"reflect"
	"sync"
	"time"
)

var (
	_app *Apex
)

// Apex is a microservice pattern IoC. It takes idea from microservices:
// 1. Services are like modules/beans, providing functionality through exposed interface.
// 2. Services all have identidy, so as modules/beans
// 3. Services handles their own dependencies, with some help from framework. So as Apex
// do, it helps service to boot after dependencies are ready, but its service's choice to
// not wait.
// 4. Services could go down by accident, so the graceful shutdown in order is not granteed,
// every module/bean should handle errors carefully when the dependent may cause error.
// 4. Services bootstrap concurrently, so are the module/bean initialized.
//
// It works in these way:
//  1. all public data structure are under into pkg/dto
//     1.1 data structures should handle their dependencies carefully here.
//  2. all public interfaces are under pkg/service or pkg/xxx directories
//     2.1 interfaces should avoid dependency as much as possible
//  3. Get depends from Apex and wire manually in Constructor when needed.
//     And ask Apex to initialized constructors.
//
// And, Apex is always global and singleton. It should be enough for most cases.
type Apex struct {
	services  sync.Map
	waitings  map[string][]chan int // all chan waiting for the object to be initialized
	initerCnt int                   // todo: use atomic
	callbacks []func()
	sync.RWMutex
}

type Opt struct {
	V any
}

// Get returns the registered implementation for the name.
func Get(name string) any {
	// during initialization, it runs in go routine
	for range 100 {
		if obj, ok := _app.services.Load(name); ok {
			return obj
		}
		// wait on signal
		obj := blockGet(name)
		if obj != nil {
			return obj
		}
	}

	return nil
}

// Wire is a help function for Get to reduce code
func Wire[T any](name string, p *T) {
	if obj := Get(name); obj != nil {
		*p = obj.(T)
	}
	p = nil
}

func blockGet(name string) any {
	_app.Lock()
	// After Lock, check again.
	if obj, ok := _app.services.Load(name); ok {
		return obj
	}
	// Prepare the waiting channel. Newly registered service after this time point, will deiliver message to channel.
	chan0 := make(chan int)
	if s, ok := _app.waitings[name]; ok {
		_app.waitings[name] = append(s, chan0)
	} else {
		_app.waitings[name] = []chan int{chan0}
	}
	_app.Unlock()
	// Wait on the chan
	<-chan0

	return nil
}

// WireOpt is used to wire an service might not exist. It is used in the callback of InitOptional.
func WireOpt[T any](name string, p *T) {
	if obj, ok := _app.services.Load(name); ok {
		*p = obj.(T)
	}
	p = nil
}

// Bootstrap start the apex service
func Bootstrap() {
	_app = new(Apex)
	_app.waitings = make(map[string][]chan int)
}

// Register an instance to a name and interface, tags and depends
func Register(name string, impl any) error {
	_app.Lock()
	defer _app.Unlock()

	if _, ok := _app.services.Load(name); ok {
		return errors.New("Name Conflict: " + name)
	}
	_app.services.Store(name, impl)

	// Trigger notifications
	waitingChans, ok := _app.waitings[name]
	if !ok {
		return nil // nothing to do
	}
	for _, c := range waitingChans {
		c <- 1   // signal the waiting chan
		close(c) // ensure every wait on chan happen once only
	}

	// clear the list
	delete(_app.waitings, name)

	return nil
}

// Bootup calculate the initialize order for all services and create them
func Bootup(constructors ...func()) {
	// as in microservice env, all services are concurrently running
	for _, c := range constructors {
		go func() {
			_app.Lock()
			_app.initerCnt += 1
			_app.Unlock()
			defer func() {
				_app.Lock()
				_app.initerCnt -= 1
				_app.Unlock()
			}()

			// TODO: catch all errors here
			c()
		}()
	}
}

// WireOptions is used to wire the fields which might be nil. The callbacks are invoked after all the
// constructors are completed, and before WaitUtilAllUp returns.
func WireOptions(callbacks ...func()) {
	_app.Lock()
	defer _app.Unlock()

	_app.callbacks = append(_app.callbacks, callbacks...)
}

func execCallbacks() {
	// as in microservice env, all services are concurrently running
	for _, c := range _app.callbacks {
		go func() {
			// TODO: catch all errors here
			c()
		}()
	}
}

// WaitUtilAllUp block the main thread until all constructors are executed and return.
// TODO: tracking details about which constructors are waiting for whom. Help engineers to discovery
// dependency problems quickly. And also to break infinit waitings.
func WaitUtilAllUp() {
	for range 1000 {
		cnt := func() int {
			_app.Lock()
			defer _app.Unlock()

			return _app.initerCnt
		}()

		// Constructors complete
		if cnt == 0 {
			execCallbacks()
			return
		}
		time.Sleep(15 * time.Millisecond)
	}
}

// Name get the type of interface using reflect and treat it as registered service name
func Name(itf any) string {
	t := reflect.TypeOf(itf)
	return t.PkgPath() + t.Name()
}
