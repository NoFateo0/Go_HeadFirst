package main

import (
	"Go_Head_first/go/src/car"
	"Go_Head_first/go/src/car/headlights"
	"Go_Head_first/go/src/car/stereo"
	"Go_Head_first/go/src/car/wheels"
)

func main() {
	car.OpenDoor()
	headlights.TurnOn()
	stereo.TurnOn()
	stereo.BoostBass()
	wheels.Steer()
	wheels.Accelerate()
}
