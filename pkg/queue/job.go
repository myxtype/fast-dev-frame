package queue

import (
	"encoding/json"
	"frame/pkg/randstr"
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
		Id:   randstr.Hex(8),
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
