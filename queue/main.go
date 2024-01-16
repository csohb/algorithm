package main

import "fmt"

type Queue []interface{}

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(data interface{}) {
	*q = append(*q, data)
}

func (q *Queue) Dequeue() interface{} {
	if q.isEmpty() {
		return nil
	} else {
		data := (*q)[0]
		*q = (*q)[1:]
		return data
	}
}

func main() {
	q := Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	fmt.Printf("%v\n", q)

	q.Dequeue()
	fmt.Printf("%v\n", q)
}
