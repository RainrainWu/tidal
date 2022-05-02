package main

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v3"
)

type Orbital struct {
	Name string `yaml:"name"`
	Celestials map[string]Celestial `yaml:"celestials"`
}

func NewOrbital(config_path string) Orbital {

	new_orbital := Orbital{}
	new_orbital.load(config_path)

	return new_orbital
}

func (orbital *Orbital) load(config_path string) {

    file, err := ioutil.ReadFile(config_path)
    if err != nil {
        log.Fatal(err)
    }

	err = yaml.Unmarshal(file, &orbital)
    if err != nil {
        log.Fatal(err)
    }
}