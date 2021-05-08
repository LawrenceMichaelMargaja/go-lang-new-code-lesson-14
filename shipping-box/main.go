package main

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"sync"
)

type Product struct {
	Len int
	Wid int
	Hei int
}

type Box struct {
	Len int
	Wid int
	Hei int
}

// getBestBox returns the best box fitting for the products given.
// This solution uses concurrency to solve the problem.
// Proposed solution: Use concurrency to calculate the volume of both boxes and
// product.
func getBestBox(availableProducts []Product, boxes []Box) (*Box, error) {
	var productTotalVolume int
	var boxInfo []int

	var wg sync.WaitGroup
	var l = sync.Mutex{}

	// Get total volume
	for _, v := range availableProducts {
		wg.Add(1)
		go func(p Product, productTotalVolume *int) {
			defer wg.Done()
			t := p.Len * p.Wid * p.Hei
			l.Lock()
			*productTotalVolume += t
			l.Unlock()
		}(v, &productTotalVolume)
	}

	// Get total box volumes
	for i, v := range boxes {
		wg.Add(1)
		go func(i *int, b Box, boxInfo *[]int) {
			defer wg.Done()

			l.Lock()
			t := b.Len * b.Wid * b.Hei

			*boxInfo = append(*boxInfo, t)
			l.Unlock()
		}(&i, v, &boxInfo)
	}

	wg.Wait()

	// Placeholders set at 9999
	bestBoxVolumeSelectedIndex := 9999
	bestBoxVolumeSelectedVal := 9999

	for i, boxVolume := range boxInfo {
		// Check if box volume is greater than
		if boxVolume > productTotalVolume {
			// Check if what you currently have is the initial value
			if bestBoxVolumeSelectedIndex == 9999 {
				bestBoxVolumeSelectedIndex = i
				bestBoxVolumeSelectedVal = boxVolume
			} else {
				// Check if the box volume is greater than productTotal Volume and LESS than
				if boxVolume < bestBoxVolumeSelectedVal {
					bestBoxVolumeSelectedIndex = i
					bestBoxVolumeSelectedVal = boxVolume
				}
			}
		}
	}

	if bestBoxVolumeSelectedIndex == 9999 {
		return nil, errors.New("no proper box selected because none fits")
	}

	return &boxes[bestBoxVolumeSelectedIndex], nil
}


func init() {
	logConfig := zap.Config{
		OutputPaths: []string{"stdout", "/tmp/logs"},
		Encoding:    "json",
	}
	var err error
	Log, err := logConfig.Build()

	if err != nil {
		fmt.Println(Log)
	}
}

func main() {
	fmt.Println("=============== SHIPPING BOX PROBLEM ===============")

	products := []Product{
		{Len: 1, Wid: 1, Hei: 1},
		{Len: 4, Wid: 3, Hei: 1},
		{Len: 3, Wid: 3, Hei: 3},
	}

	boxes := []Box{
		{Len: 10, Wid: 3, Hei: 10},
		{Len: 10, Wid: 3, Hei: 10},
		{Len: 10, Wid: 3, Hei: 1},
	}

	result, err := getBestBox(products, boxes)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Here's your box", result)
		fmt.Println("Here's your box's total volume", result.Len * result.Wid * result.Hei)
	}
}
