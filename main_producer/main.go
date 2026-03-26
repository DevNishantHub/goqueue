package main

import (
	"fmt"
	"time"

	"github.com/DevNishantHub/goqueue/broker"
	"github.com/DevNishantHub/goqueue/task"
)

func main(){
	b := broker.New("tasks","localhost:6379")
	t1 := task.New("add",10,32)
	t2 := task.New("greet", "Nishant")
	t3 := task.New("multiply",10,20)
	t4 := task.New("reverse","Hello")
	t5 := task.New("iseven",45)
	tasks := []*task.Task{t1, t2, t3, t4, t5}
	if err := b.Enqueue(t1); err != nil {
    	fmt.Printf("failed to enqueue t1: %v\n", err)
	}
	if err := b.Enqueue(t2); err != nil {
		fmt.Printf("failed to enqueue t2: %v\n", err)
	}
	if err := b.Enqueue(t3); err != nil {
		fmt.Printf("failed to enqueue t3: %v\n", err)
	}
	if err := b.Enqueue(t4); err != nil {
		fmt.Printf("failed to enqueue t4: %v\n", err)
	}
	if err := b.Enqueue(t5); err != nil {
		fmt.Printf("failed to enqueue t5: %v\n", err)
	}
	for _,t := range tasks{
		time.Sleep(6*time.Second)
		result, err := b.GetResult(t.Id)
		if err != nil {
			fmt.Printf("failed to get result: %v\n", err)
		} else {
			fmt.Printf("task: %s | status: %s | result: %v\n", result.Id, result.Status, result.Result)
		}
	}
}
