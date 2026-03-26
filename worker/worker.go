package worker

import (
	"fmt"

	"github.com/DevNishantHub/goqueue/broker"
	"github.com/DevNishantHub/goqueue/task"
)

type Worker struct{
	broker 		*broker.Broker
	registry 	map[string]func(args []interface{})(interface{},error)
}

func New (b *broker.Broker) *Worker {
	return &Worker{
		broker: b,
		registry: make(map[string]func(args []interface{}) (interface{}, error)),
	}
}

func (w *Worker) Register(name string,fn func(args []interface{})(interface{},error)) {
	w.registry[name] = fn
}

func (w *Worker)Run() {
	fmt.Println("worker started, listening for tasks...")
	for{
		t,err := w.broker.Dequeue()
		if err != nil {
			fmt.Println("Dequeue Error: ",err)
			continue
		}
		go func(t *task.Task) {
			fn,ok := w.registry[t.FuncName]
    		fmt.Printf("worker picked up task: %s | func: %s\n", t.Id, t.FuncName)
			if !ok {
				t.Status = task.StatusError
				t.Error = fmt.Sprintf("no handler for: %s", t.FuncName)
				w.broker.SetResult(t.Id, t)
				return
			}

			result,err := fn(t.Args)
			if err != nil {
				t.Status =task.StatusError
				t.Error = err.Error()
			} else {
				t.Status = task.StatusSuccess
				t.Result = result
			}
			if err := w.broker.SetResult(t.Id,t); err != nil {
				fmt.Printf("failed to set result: %v\n", err)
			}
			fmt.Printf("worker finished task: %s | result: %v\n", t.Id, t.Result)
		}(t)
	}
}

