package algo

import (
	"fmt"

	"github.com/Zubayear/ryushin/priorityqueue"
)

type PriorityQueue struct {
	Collection []any
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{}
}

func Tester() {
	pqueue := priorityqueue.NewBinaryHeap[string]()
	pqueue.Add("dead")
	pqueue.Add("tehehehe")
	pqueue.Add("art")
	pqueue.Add("beef")
	for {
		if pqueue.Size() != 0 {
			el, _ := pqueue.Poll()
			fmt.Println(el)
			continue
		}
		break
	}
}

