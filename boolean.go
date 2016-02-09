package gotypes

type boolean bool

func (a *boolean) Type() Type {
	return Boolean
}
func (a *boolean) Array() []Value {
	return nil
}

func (a *boolean) Real() float64 {
	return 0.0
}
func (a *boolean) Integer() int64 {
	return 0
}
func (a *boolean) Bool() bool {
	return bool(*a)
}
func (a *boolean) Text() string {
	return ""
}

func (a *boolean) Unknown() interface{} {
	return nil
}
