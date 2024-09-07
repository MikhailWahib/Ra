package object

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnvironment() *Environment {
	return &Environment{
		store: make(map[string]Object),
	}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer

	return env
}

func (e *Environment) Get(key string) (Object, bool) {
	val, ok := e.store[key]
	if !ok && e.outer != nil {
		val, ok = e.outer.Get(key)
	}
	return val, ok
}

func (e *Environment) Set(key string, val Object) Object {
	e.store[key] = val
	return val
}
