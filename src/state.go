package main

var currentState = "menu"

func SetState(newState string) {
	currentState = newState
}

func GetState() string {
	return currentState
}
