package main

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type YS struct {
	A string `yaml:"a"`
}

func main() {
	buf, err := os.ReadFile("D:\\data\\program\\go\\GinHello\\test\\test.yaml")
	if err != nil {
		log.Fatalln("failed to read file", err.Error())
	}
	var ys YS
	err2 := yaml.Unmarshal(buf, &ys)
	if err2 != nil {
		log.Fatalln("failed to read file", err2.Error())
	}
	log.Println("a:", ys.A)
}
