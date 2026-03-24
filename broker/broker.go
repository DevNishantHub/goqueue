package broker

import (
	"context"
	"fmt"
	"time"
	"github.com/DevNishantHub/goqueue/task"
	"github.com/redis/go-redis/v9"
)

type broker struct{
	client 		*redis.Client
	queueName 	string
}

func New(queuename string, redisAddr string) *broker{
	return &broker{
		client: redis.NewClient(&redis.Options{
			Addr: redisAddr,
			Password: "",
			DB: 0,
			Protocol: 2,
		}),
		queueName: queuename ,
	}
}

func(b *broker) Enqueue(t *task.Task) error{

	json,err := task.Marshal(t)
	if err != nil {
		return fmt.Errorf("%w",err)
	}

	pushResult,err := b.client.LPush(context.Background(),b.queueName,json).Result()
	if err != nil{
		return fmt.Errorf("%w",err)
	}

	fmt.Println(pushResult)
	return nil
}

func(b *broker) Dequeue() (*task.Task,error) {
	popResult,err := b.client.BRPop(context.Background(),30*time.Second,b.queueName).Result()
	if err != nil {
		if err == redis.Nil{
			return nil , fmt.Errorf("Time Reached")
		} else{
			return nil,fmt.Errorf("%w",err)
		}
	}
	result,err := task.Unmarshal(popResult[1])
	if err != nil {
		return nil,err
	}
	return result,nil
}
func (b *broker) SetResult(taskId string,t *task.Task) (error) {
	key := fmt.Sprintf("result:%s",taskId)
	json,err := task.Marshal(t)
	if err != nil {
		return err
	}
	_,err = b.client.Set(context.Background(),key,json,time.Hour).Result()
	return err
}
func (b *broker) GetResult(taskId string)(*task.Task,error){
	key := fmt.Sprintf("result:%s",taskId)
	value,err := b.client.Get(context.Background(),key).Result()
	if err != nil {
		return nil,fmt.Errorf("Key Not found")
	}
	json,err := task.Unmarshal(value)
	if err != nil {
		return nil,fmt.Errorf("error Unmarshaling the json")
	}
	return json,nil
}

