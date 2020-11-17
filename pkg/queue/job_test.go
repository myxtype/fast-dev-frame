package queue

import "testing"

func TestNewDelayQueueJob(t *testing.T) {
	_, err := NewDelayQueueJob("test")
	if err != nil {
		t.Error(err)
	}
}

func TestDelayQueueJob_Bytes(t *testing.T) {
	job, err := NewDelayQueueJob(1000)
	if err != nil {
		t.Error(err)
	}

	t.Log(string(job.Bytes()))
}

func TestDelayQueueJob_Unmarshal(t *testing.T) {
	job, err := NewDelayQueueJob(10000)
	if err != nil {
		t.Error(err)
	}

	var id int64
	if err := job.Unmarshal(&id); err != nil {
		t.Error(err)
	}
	if id != 10000 {
		t.Error("Error 10000")
	}
}

func TestDelayQueueJob_Unmarshal2(t *testing.T) {
	job, err := NewDelayQueueJob(struct {
		Id int64
	}{
		Id: 1000,
	})
	if err != nil {
		t.Error(err)
	}

	var msg struct {
		Id int64
	}
	if err := job.Unmarshal(&msg); err != nil {
		t.Error(err)
	}
	if msg.Id != 1000 {
		t.Error("Error 1000")
	}
}
