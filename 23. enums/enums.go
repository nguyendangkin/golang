package main

import "fmt"

type LightState int

const (
	StateOff LightState = iota
	StateOn
	StateBroken
)

var stateName = map[LightState]string{
	StateOff:    "off",
	StateOn:     "on",
	StateBroken: "broken",
}

func (ls LightState) String() string {
	return stateName[ls]
}

func transition(s LightState) LightState {
	switch s {
	case StateOff:
		return StateOn
	case StateOn:
		return StateOff
	case StateBroken:
		return StateBroken
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}

func main() {
	current := transition(StateOff)
	fmt.Println(current) // should print "on"

	next := transition(current)
	fmt.Println(next) // should print "off"
}
