package fishing

const (
	ServiceID = "apex_demo.fishing"
)

type Fishing interface {
	ThrowPole(strength int32)
	PullPole(speed int32)
}
