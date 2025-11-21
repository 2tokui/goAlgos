package main

import (
	"fmt"

	"github.com/2tokui/goAlgos/algo"
	"github.com/Zubayear/ryushin/priorityqueue"
)

func main() {
	fmt.Println("======== Stack ========")

	theStack := algo.NewStack()
	fmt.Println(theStack.Empty())
	theStack.Push(1)
	theStack.Push(false)
	theStack.Push("hello")

	fmt.Println(theStack.Peek())
	theStack.Pop()
	fmt.Println(theStack.Peek())
	theStack.Pop()
	theStack.Pop()
	theStack.Pop()
	theStack.Pop()
	fmt.Println(theStack.Peek())

	//for range(1000000000000) { // Don't do this
	//	theStack.Push("the bomb is ticking...")
	//}

	fmt.Println(theStack.Peek())

	fmt.Println("======== Queue ========")
	queue := algo.NewQueue()
	queue.Enqueue(1)
	queue.Enqueue("this guy")
	queue.Enqueue(false)
	queue.Enqueue(1337)

	fmt.Println("Length:", queue.Length())
	fmt.Println("IsEmpty:", queue.IsEmpty())

	fmt.Println("Peek:", queue.Peek())
	queue.Dequeue()
	fmt.Println("Peek:", queue.Peek())
	queue.Dequeue()
	fmt.Println("Peek:", queue.Peek())
	queue.Dequeue()
	fmt.Println("Peek:", queue.Peek())
	queue.Dequeue()
	fmt.Println("Peek:", queue.Peek())

	fmt.Println("Length:", queue.Length())
	fmt.Println("IsEmpty:", queue.IsEmpty())

	fmt.Println("======== Priority Queue ========")
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
	fmt.Println("======== Linked List ========")

	ll := algo.NewLinkedList()
	ll.Push("data")
	ll.Push('d')
	ll.Push(1)
	ll.Push(false)

	//ll.Pop(false)
	ll.Pop("data")

	ll.PrintEverything()
}
