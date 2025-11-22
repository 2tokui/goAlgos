package main

import (
	"fmt"
	"time"

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

	ll.PrintEverything()

	fmt.Println("======== Linear Search ========")
	var myIntArray = []int{9, 2, 3, 1, 0, 4, 6, 8}
	iterations := 1_000_000
	for range iterations {
		myIntArray = append(myIntArray, 0)
	}
	myIntArray[999999] = 5

	startTime := time.Now()
	algo.LinearSearch(myIntArray, 5)
	elapsedTime := time.Since(startTime)

	fmt.Printf("Elapsed Time: %d ns\n", elapsedTime.Nanoseconds())

	fmt.Println("======== Binary Search ========")
	orderedCollection := []int{0, 3, 5, 6, 9, 11, 13, 15, 19, 40, 60}
	findAndPrint := func(guy int) {
		position := algo.BinarySearch(orderedCollection, guy)
		if position == -1 {
			fmt.Println(guy, "doesn't exist in orderedCollection")
		} else {
			fmt.Printf("Position of %d orderedCollection: %d\n", guy, position)
		}
	}
	findAndPrint(19)
	findAndPrint(0)
	findAndPrint(9)
	findAndPrint(60)
	findAndPrint(999)
}
