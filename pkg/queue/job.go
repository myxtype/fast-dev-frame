package queue

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
)

// 队列任务
type QueueJob struct {
	Id   string          `json:"id"`
	Data json.RawMessage `json:"data"`
}

func NewDelayQueueJob(msg interface{}) (*QueueJob, error) {
	data, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return &QueueJob{
		Id:   uuid.NewV4().String(),
		Data: data,
	}, nil
}

func (j *QueueJob) Bytes() []byte {
	b, _ := json.Marshal(j)
	return b
}

func (j *QueueJob) Unmarshal(dst interface{}) error {
	return json.Unmarshal(j.Data, dst)
}
