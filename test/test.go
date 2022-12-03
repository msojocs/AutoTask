package main

import (
	"log"
)

type YS struct {
	A string `yaml:"a"`
}

type FishType int

const (
	A FishType = iota
	B
	C
	D
)

func (f FishType) String() string {
	return [...]string{"A", "B", "C", "D"}[f]
}

func main() {
	var a FishType
	a = 9
	log.Println(a)

	//buf, err := os.ReadFile("D:\\data\\program\\go\\GinHello\\test\\test.yaml")
	//if err != nil {
	//	log.Fatalln("failed to read file", err.Error())
	//}
	//var ys YS
	//err2 := yaml.Unmarshal(buf, &ys)
	//if err2 != nil {
	//	log.Fatalln("failed to read file", err2.Error())
	//}
	//log.Println("a:", ys.A)
}
