package gotypes

type text string

func (a *text) Type() Type {
	return Text
}

func (a *text) Array() []Value {
	return nil
}

func (a *text) Real() float64 {
	return 0.0
}
func (a *text) Integer() int64 {
	return 0
}
func (a *text) Bool() bool {
	return false
}
func (a *text) Text() string {
	return string(*a)
}

func (a *text) Unknown() interface{} {
	return nil
}
