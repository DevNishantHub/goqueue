package main

import (
	"github.com/DevNishantHub/goqueue/broker"
	"github.com/DevNishantHub/goqueue/worker"
)

func main(){
	b := broker.New("tasks","localhost:6379")
	w := worker.New(b)
	w.Register("add",func(args []interface{}) (interface{}, error) {
		a := args[0].(float64)
		b := args[1].(float64)
		return a+b,nil
	})
	w.Register("greet", func(args []interface{}) (interface{}, error) {
		name := args[0].(string)
		return "hello " + name, nil
	})
	// multiply two numbers
	w.Register("multiply", func(args []interface{}) (interface{}, error) {
		a := args[0].(float64)
		b := args[1].(float64)
		return a * b, nil
	})

	// reverse a string
	w.Register("reverse", func(args []interface{}) (interface{}, error) {
		s := args[0].(string)
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes), nil
	})

	// check if number is even
	w.Register("iseven", func(args []interface{}) (interface{}, error) {
		n := args[0].(float64)
		return int(n)%2 == 0, nil
	})
	w.Run()
}
