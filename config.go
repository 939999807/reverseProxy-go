package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Port   string
	Remote string
	Ip     string
	Dns    []string
}

func parseConfig() Config {
	cfg := Config{}
	yamlFile, err := ioutil.ReadFile("conf/config.yaml")
	if err != nil {
		log.Printf("conf/config.yaml not found err #%v ", err)
	}
	log.Println(string(yamlFile))
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return cfg
}
