package gotypes

type unknown struct {
	value interface{}
}

func (a *unknown) Type() Type {
	return Unknown
}

func (a *unknown) Array() []Value {
	return nil
}

func (a *unknown) Real() float64 {
	return 0.0
}

func (a *unknown) Integer() int64 {
	return 0
}

func (a *unknown) Bool() bool {
	return false
}
func (a *unknown) Text() string {
	return ""
}

func (a *unknown) Unknown() interface{} {
	return a.value
}
