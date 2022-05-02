package main

type CelestialSeries string

const (
    Constant CelestialSeries = "constant"
    Periodic CelestialSeries = "periodic"
    Burst CelestialSeries = "burst"
)

type Celestial struct {
	Name string `yaml:"name"`
	Series CelestialSeries `yaml:"series"`

    // Constant
    Frequency float32 `yaml:"frequency"`

    // Periodic
    Period int `yaml:"period"`
    Volume int `yaml:"volume"`

    // Burst
    // Period int `yaml:"period"`
    // Volume int `yaml:"volume"`
    Interval int `yaml:"interval"`
}

func (c *Celestial) propagate_flows() {

    field := NewField(int(c.Period))
    field.MapWithSin().ShiftToPositive().ScaleByVolume(c.Volume)

    wave := NewWave(field.get_ripples())
    wave.launch()
}