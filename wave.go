package main

import (
	"fmt"
	"time"
)

type Wave struct {
	Gravitations []int
}

func NewWave(gravitations []int) Wave {

	return Wave{
		Gravitations: gravitations,
	}
}

func (w *Wave) launch() {

	for _, e := range w.Gravitations {

		fmt.Printf("splash %d\n", e)
		time.Sleep(1 * time.Second)
	}
}
