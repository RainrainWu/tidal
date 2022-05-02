package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

)

var CONFIGS_DIR = os.Getenv("CONFIGS_DIR")


func get_config_paths() []string {

	config_files, err := ioutil.ReadDir(CONFIGS_DIR)
	if err != nil {
		log.Fatal(err)
	}

	config_paths := make([]string, 0)
	for _, config_file := range config_files {
		path := strings.Join([]string{CONFIGS_DIR, config_file.Name()}, "/")
		config_paths = append(config_paths, path)
	}

	return config_paths
}