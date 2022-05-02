package main

import (
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	config_paths := get_config_paths()
	for _, config_path := range config_paths {
		orb := NewOrbital(config_path)
		for _, celestial := range orb.Celestials {
			celestial.propagate_flows()
		}
	}
}