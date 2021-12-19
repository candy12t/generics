package main

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

var fizzbuzz = `
fizz: FIZZ
buzz: BUZZ
`

type FizzBuzz struct {
	Fizz string `yaml:"fizz"`
	Buzz string `yaml:"buzz"`
}

var hogefuga = `
hoge: HOGE
fuga: FUGA
`

type HogeFuga struct {
	Hoge string `yaml:"hoge"`
	Fuga string `yaml:"fuga"`
}

type YAML interface {
	FizzBuzz | HogeFuga
}

func main() {
	fb := ParseYaml[FizzBuzz](fizzbuzz)
	fmt.Printf("FizzBuzz: %+v\n", fb)
	hf := ParseYaml[HogeFuga](hogefuga)
	fmt.Printf("HogeFuga: %+v\n", hf)
}

func ParseYaml[V YAML](data string) *V {
	v := new(V)
	yaml.Unmarshal([]byte(data), v)
	return v
}
