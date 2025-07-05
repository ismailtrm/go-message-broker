package broker

import "sync"

type Message struct {
	ID      string `json:"id"`      // message id
	Payload string `json:"payload"` // message payload
}

type Queue struct {
	mu      sync.Mutex // concurrency control
	storage []Message  // message storage
}

func (q *Queue) Enqueue(msg Message) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.storage = append(q.storage, msg)
}

func (q *Queue) Dequeue() (Message, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if len(q.storage) == 0 {
		return Message{}, false
	}
	msg := q.storage[0]
	q.storage = q.storage[1:] // remove the first message
	return msg, true
}

var GlobalQueue = &Queue{}
