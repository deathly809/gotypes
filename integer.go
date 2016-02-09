package gotypes

type integer int64

func (a *integer) Type() Type {
	return Integer
}

func (a *integer) Array() []Value {
	return nil
}

func (a *integer) Real() float64 {
	return 0.0
}
func (a *integer) Integer() int64 {
	return int64(*a)
}
func (a *integer) Bool() bool {
	return false
}
func (a *integer) Text() string {
	return ""
}

func (a *integer) Unknown() interface{} {
	return nil
}
