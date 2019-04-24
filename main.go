package main

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"test/game"
	"time"
)

const yamlFile = "config.yaml"

type configuration struct {
	Properties struct {
		First  string `yaml:"first"`
		Second string `yaml:"second"`
	}
	Third []string `yaml:"test"`
}

func (c *configuration) getConf() *configuration {

	yamlFile, err := ioutil.ReadFile(yamlFile)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {
	var c configuration
	c.getConf()

	fmt.Println("I think this worked maybe?")
	fmt.Println(c)

	pixelgl.Run(game.StartGame)

	time.Sleep(5 * time.Second)
}
