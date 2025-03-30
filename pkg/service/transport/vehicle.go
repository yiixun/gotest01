package transport

const (
	ServiceID = "apex_demo.vehicle"
)

type Vehicle interface {
	Deliver()
}
