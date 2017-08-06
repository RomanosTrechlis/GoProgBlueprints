package meander

// Facade has a single method Public that
// transforms an internal structure to an exported structure.
type Facade interface {
	Public() interface{}
}

// Public method takes an interface and transform it to another
// if it implements the Facade.Public method.
func Public(o interface{}) interface{} {
	if p, ok := o.(Facade); ok {
		return p.Public()
	}
	return o
}
