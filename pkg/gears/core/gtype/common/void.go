package common

type Void interface{}

func NewVoid() Void {
	return void{}
}

type void struct{}
