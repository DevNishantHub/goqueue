package worker

import "github.com/DevNishantHub/goqueue/broker"

type Worker struct{
	broker 		*broker.Broker
	registry 	map[string]func(args []interface{})(interface{},error)
}

func new (b *broker.Broker) *Worker {
	return &Worker{
		broker: b,
		registry: make(map[string]func(args []interface{}) (interface{}, error)),
	}
}

func (w *Worker) Register(name string,fn func(args []interface{})(interface{},error)) {
	w.registry[name] = fn
}

