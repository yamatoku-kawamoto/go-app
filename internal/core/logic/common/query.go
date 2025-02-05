package common

type Query interface {
	Validate() error
}
