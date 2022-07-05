package main

import "fmt"

type truck struct {
	capacity int
	avgSpeed float64
	height   int
}

func (t truck) getAvgSpeed() float64 {
	return t.avgSpeed
}

func (t truck) String() string {
	return "truck"
}

type train struct {
	capacity int
	avgSpeed float64
	carts    int
}

func (t train) getAvgSpeed() float64 {
	return t.avgSpeed
}
func (t train) String() string {
	return "train"
}

type pickup struct {
	capacity int
	avgSpeed float64
	gasTank  int
}

func (p pickup) getAvgSpeed() float64 {
	return p.avgSpeed
}
func (p pickup) String() string {
	return "pickup"
}

type avgSpeedGetter interface {
	getAvgSpeed() float64
}

func minInRoad(asg avgSpeedGetter, pointA, pointB int) float64 {
	distance := pointB - pointA
	kmPerMinute := asg.getAvgSpeed() / 60
	roadTime := float64(distance) / kmPerMinute
	return roadTime
}

type transportNameGetter interface {
	getTransportName() string
}

type transport interface {
	avgSpeedGetter
	fmt.Stringer
}

func main() {
	pointA, pointB := 5, 30

	transports := []transport{
		truck{avgSpeed: 70},
		train{avgSpeed: 100},
		pickup{avgSpeed: 80},
	}

	for _, t := range transports {
		mirTrack := minInRoad(t, pointA, pointB)
		fmt.Printf("\nMinutes in road by %s : %f", t.String(), mirTrack)
	}

}
