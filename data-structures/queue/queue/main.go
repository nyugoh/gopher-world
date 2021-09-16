package queue

import "fmt"

type ListNode struct {
	Data int
	Next *ListNode
}

type Queue struct {
	Front *ListNode
	Rear  *ListNode
}

func CreateQueue() *Queue {
	q := &Queue{nil, nil}
	return q
}

func (q *Queue) IsEmpty() bool {
	if q == nil {
		return true
	}
	return q.Front == nil
}

func (q *Queue) Enqueue(data int) {
	newNode := &ListNode{data, nil}

	if q.Rear != nil {
		q.Rear.Next = newNode
	}
	q.Rear = newNode
	if q.Front == nil {
		q.Front = q.Rear
	}
}

func (q *Queue) Dequeue() int {
	var data int
	if q.IsEmpty() {
		return data
	}
	data = q.Front.Data
	tempNode := q.Front
	q.Front = tempNode.Next
	tempNode = nil
	return data
}

func (q *Queue) DeleteQ() {
	for q.Front != nil {
		tempNode := q.Front
		q.Front = q.Front.Next
		fmt.Println(tempNode.Data)
		tempNode = nil
	}
}

func (q *Queue) Print() {
	tempNode := q.Front
	for tempNode != nil {
		fmt.Print(tempNode.Data)
		if tempNode.Next != nil {
			fmt.Print(" <- ")
		}
		tempNode = tempNode.Next
	}
	fmt.Println()
}

func (q *Queue) Size() (length int) {
	tempNode := q.Front
	for tempNode != nil {
		length += 1
		tempNode = tempNode.Next
	}
	return
}
