package main

import (
	"errors"
	"fmt"
)

func minInRoad(asg avgSpeedGetter, pointA, pointB int) (float64, error) {
	avgSpeed := asg.getAvgSpeed()

	if pointA > pointB {
		return 0, errors.New("wrong data input")
	}

	if err := validateSpeed(avgSpeed); err != nil {
		return 0, fmt.Errorf("failed to calculate minutesInRoad: %w", err)
	}

	distance := pointB - pointA
	kmPerMinute := asg.getAvgSpeed() / 60
	roadTime := float64(distance) / kmPerMinute
	return roadTime, nil
}

var slowTransportErr error = errors.New("it is too slow")

func validateSpeed(avgSpeed float64) error {
	if avgSpeed < 80 {
		return slowTransportErr
	}
	return nil
}

func main() {
	pointA, pointB := 5, 30

	transports := []transport{
		truck{avgSpeed: 70},
		train{avgSpeed: 100},
		pickup{avgSpeed: 80},
	}
	// truck:
	// 		is not an option
	//		  additional wrap
	//			failed to calculate
	//				it is too slow
	for _, t := range transports {
		mir, err := minInRoad(t, pointA, pointB)
		if err != nil { // if err, adding wrap
			err = fmt.Errorf("additional wrap: %w", err)
		}

		if errors.Is(err, slowTransportErr) {
			fmt.Printf("\n%s: is not an option. %v", t, err)
			continue
		}

		if err != nil {
			fmt.Printf("All is bad: %v.\n", err)
			return
		}

		fmt.Printf("\nMinutes in road by %-6v : %.2f", t, mir)
	}

}
