package gotypes

type array []Value

func (a *array) Type() Type {
	return Array
}
func (a *array) Array() []Value {
	return *a
}

func (a *array) Real() float64 {
	return 0.0
}
func (a *array) Integer() int64 {
	return 0
}
func (a *array) Bool() bool {
	return false
}
func (a *array) Text() string {
	return "array"
}

func (a *array) Unknown() interface{} {
	return nil
}
