package queue

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
)

// 队列任务
type DelayQueueJob struct {
	Id   string          `json:"id"`
	Data json.RawMessage `json:"data"`
}

func NewDelayQueueJob(msg interface{}) (*DelayQueueJob, error) {
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return &DelayQueueJob{
		Id:   uuid.NewV4().String(),
		Data: data,
	}, nil
}

func (j *DelayQueueJob) Bytes() []byte {
	b, _ := json.Marshal(j)
	return b
}

func (j *DelayQueueJob) Unmarshal(dst interface{}) error {
	return json.Unmarshal(j.Data, dst)
}
