package main

var currentState = "menu" // valeur initiale

func SetState(newState string) {
	currentState = newState
}

func GetState() string {
	return currentState
}
