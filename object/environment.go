package object

type Environment struct {
	store map[string]Object
}

func NewEnvironment() *Environment {
	return &Environment{
		store: make(map[string]Object),
	}
}

func (e *Environment) Get(key string) (Object, bool) {
	val, ok := e.store[key]
	return val, ok
}

func (e *Environment) Set(key string, val Object) Object {
	e.store[key] = val
	return val
}
