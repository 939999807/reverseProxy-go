package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Port  string
	Route map[string]string
}

func parseConfig() Config {
	cfg := Config{}
	//dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	//if err != nil {
	//	log.Fatalf("filepath not found err #%v ", err)
	//}
	yamlFile, err := ioutil.ReadFile(dir + "/conf/config.yaml")
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
