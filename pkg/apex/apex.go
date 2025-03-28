package apex

// Apex is the simplest service registery. It works in these way:
//  1. all public data structure are under into pkg/dto
//     1.1 data structures should handle their dependencies carefully here.
//  2. all public interfaces are under pkg/service or pkg/xxx directories
//     2.1 interfaces should avoid dependency as much as possible
//  3. Apex is a struct combines interfaces. Internal service register
//     service implementation to Apex implementation.
//
// Currently, it doesn't provide "real" register function. Only to make
// code structure clean. As Apex is used as a registery, currently it is
// better to be in internal package. However, in future it would be better.
//
// And, Apex is always global and singleton. It should be enough for most cases.
type Apex struct {
}

// Register register an interface with a name, and report its dependent interfaces' name.
// Depend is a name instead of the interface type, because name and interface are defined
// at the same place.
func Register(itf any, name string, tags []string, depends []string) {

}

// GetByName returns the registered implementation for the name.
func GetByName(name string) any {
	return nil
}

// GetByType returns the registered implementation for the interface.
func GetByType[T any](itf T) T {
	return itf
}

// ListByTag returns the registered implementation for the name.
func ListByTag(tag string) []any {
	return nil
}

// ListByType returns a list of registered implementation for the interface.
func ListByType[T any](itf T) []T {
	return nil
}
