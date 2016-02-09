package gotypes

type real float64

func (a *real) Type() Type {
	return Real
}

func (a *real) Array() []Value {
	return nil
}

func (a *real) Real() float64 {
	return float64(*a)
}

func (a *real) Integer() int64 {
	return 0
}

func (a *real) Bool() bool {
	return false
}

func (a *real) Text() string {
	return ""
}

func (a *real) Unknown() interface{} {
	return nil
}
