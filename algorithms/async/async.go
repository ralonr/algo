package async

type forget struct{}
type wait struct{}

// Dispatcher interface type for both sync and async func calls
type Dispatcher interface {
	Dispatch(func())
}

// Dispatch for async func call
func (fr *forget) Dispatch(f func()) {
	go f() // async call for real code
}

func Forget() Dispatcher {
	return &forget{}
}

// Dispatch for sync call. can be used for testing
func (wa *wait) Dispatch(f func()) {
	f()
}

func Wait() Dispatcher {
	return &wait{}
}
