package java

import "time"

type Object = interface{}

func NewObject() *Object {
	return nil
}

type Date = time.Time

func NewDate() *Date {
	return nil
}

type Float32 = float32

func NewFloat32() Float32 {
	return 0
}
