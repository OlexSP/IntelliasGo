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

type transport interface {
	avgSpeedGetter
	fmt.Stringer
}
