package mylib

import (
	"reflect"
	"sync"
)

var mu = &sync.Mutex{}

type Queue struct {
	Vals []interface{}
}

func Init() Queue {
	return Queue{make([]interface{}, 0)}
}

// Checks if inserted item is of same type as existing Queue values
// panics if not
func (q *Queue) typeCheck(item interface{}) {
	if len(q.Vals) > 0 && reflect.TypeOf(q.Vals[0]) != reflect.TypeOf(item) {
		panic("Attempting to add more than one Type to Queue")
	}

}

// Place item at back of queue
func (q *Queue) Enqueue(item interface{}) {
	for {
		mu.Lock()
		defer mu.Unlock()
		q.typeCheck(item)
		q.Vals = append(q.Vals, item)
		return
	}
}

// Return next item in queue & pop front
func (q *Queue) Dequeue() (x *interface{}) {
	if len(q.Vals) > 0 {
		for {
			mu.Lock()
			defer mu.Unlock()
			x, q.Vals = &q.Vals[0], q.Vals[1:]
			return
		}
	}
	return nil
}

// Push item to front of queue
func (q *Queue) Jump(item interface{}) {
	q.typeCheck(item)
	for {
		mu.Lock()
		defer mu.Unlock()
		q.Vals = append([]interface{}{item}, q.Vals...)
		return
	}
}

// Emplace item in Queue
func (q *Queue) CutIn(item interface{}, place int) {
	q.typeCheck(item)
	for {
		mu.Lock()
		defer mu.Unlock()
		q.Vals = append(q.Vals[:place], append([]interface{}{item}, q.Vals[place:]))
		return
	}
}

// Tests whether the Queue is empty
func (q *Queue) IsEmpty(item interface{}) bool {
	if len(q.Vals) == 0 {
		return true
	}
	return false
}
