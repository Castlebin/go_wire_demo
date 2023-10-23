// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

// InitializeEvent 声明 injector 的函数签名
func InitializeEvent(msg string) Event {
	message := NewMessage(msg)
	greeter := NewGreeter(message)
	event := NewEvent(greeter)
	return event
}