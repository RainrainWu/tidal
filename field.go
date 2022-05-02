package main

import (
	"fmt"
	"math"
)

type Field struct {
	Gravitations []float64
}

func NewField(period int) Field {

	new_flow := Field{}
	for ptr := 0; ptr < period; ptr++ {
		new_flow.Gravitations = append(new_flow.Gravitations, float64(ptr) / float64(period))
	}
	return new_flow
}

func (f *Field) MapWithSin() *Field {

	for i, e := range f.Gravitations {
		f.Gravitations[i] = math.Sin(e * 2 * math.Pi)
	}
	return f
}

func (f *Field) ShiftToPositive() *Field {

	depth := float64(0.0)
	for i, e := range f.Gravitations {
        if i==0 || e < depth {
            depth = e
        }
    }
	for i, e := range f.Gravitations {
		f.Gravitations[i] = math.Abs(e + depth)
	}
	return f
}

func (f *Field) ScaleByVolume(volume int) *Field {

	var total_ripple float64 = 0
	for _, e := range f.Gravitations {
		total_ripple += e
	}

	var accumulate_ripple float64 = 0
	var accumulate_volume int = 0
	volume_float := float64(volume)
	for i, e := range f.Gravitations {
		accumulate_ripple += e

		target_volume := volume_float * (accumulate_ripple / total_ripple)
		delta_volume := math.Floor(target_volume - float64(accumulate_volume))
		f.Gravitations[i] = delta_volume
		accumulate_volume += int(delta_volume)
	}
	return f
}

func (f *Field) get_ripples() []int {

	gravitations := []int{}
	for _, e := range f.Gravitations {
		gravitations = append(gravitations, int(e))
	}
	fmt.Println(gravitations)
	return gravitations
}