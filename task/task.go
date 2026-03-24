package task

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/google/uuid"
)

type Status string

const(
	StatusPending 	Status = "pending"
	StatusSuccess	Status = "success"
	StatusError		Status = "error"
)

type Task struct{
	Id 			string 			`json:"id"`
	FuncName	string 			`json:"funcname"`
	Args		[]interface{} 	`json:"args"`
	Status		Status 			`json:"status"`
	Result		interface{} 	`json:"result,omitempty"`
	Error		string 			`json:"error,omitempty"`
	CreatedAt	time.Time		`json:"created_at"`
}

func new(funcName string, args ...interface{}) *Task {
	return &Task{
		Id: uuid.NewString(),
		FuncName: funcName,
		Args: args,
		Status: StatusPending,
		CreatedAt: time.Now(),
	}
}

func Marshal(t *Task)(string, error){
	b,err := json.Marshal(t)
	if err != nil {
		return "",fmt.Errorf("%w",err);
	}
	return string(b), nil
}

func Unmarshal(data string) (*Task, error) {
	var t Task
	if err := json.Unmarshal([]byte(data), &t); err != nil {
		return nil, fmt.Errorf("unmarshal task: %w", err)
	}
	return &t, nil
}
