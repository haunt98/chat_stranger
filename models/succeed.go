package models

// Wrap anything return with JSON
type WrapSucceed struct {
	Succeed bool
	Value   interface{}
}
