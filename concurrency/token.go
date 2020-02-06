package co

// Coordinator creates a coordinator to handle handoff case.
type Coordinator struct {
	chans []chan struct{}
}

// NewCoordinator creates a  Coordinator.
func NewCoordinator(n int) *Coordinator {
	chans := make([]chan struct{}, n)
	for i := 0; i < n; i++ {
		chans[i] = make(chan struct{}, 1)
	}

	return &Coordinator{
		chans: chans,
	}
}

// Accquire accquires the token.
func (coo *Coordinator) Accquire(i int) {
	<-coo.chans[i]
}

// ReleaseTo releases the token ans passes it to i.
func (coo *Coordinator) ReleaseTo(i int) {
	coo.chans[i] <- struct{}{}
}
