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

	fmt.Println("======== Interpolation Search ========")
	dataArray := []int{2, 4, 8, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096, 8192}
	idx := algo.InterpolationSearch(dataArray, 4096)
	fmt.Printf("%d is at %d\n", 4096, idx)

	fmt.Println("======== Bubble Sort ========")
	intArray := []int{5, 2, 1, 3, 9, 10, 11, 98, 1, 4, 12, 10, 30, 50, 40}
	fmt.Println("intArray:", intArray)
	algo.BubbleSort(intArray)
	fmt.Println("intArray:", intArray)

	fmt.Println("======== Selection Sort ========")
	intArrayAgain := []int{5, 2, 1, 3, 9, 10, 11, 98, 1, 4, 12, 10, 30, 50, 40}
	fmt.Println("intArrayAgain:", intArrayAgain)
	algo.SelectionSort(intArrayAgain)
	fmt.Println("intArrayAgain:", intArrayAgain)

	fmt.Println("======== Insertion Sort ========")
	intInsertionArray := []int{5, 2, 1, 3, 9, 10, 11, 98, 1, 4, 12, 10, 30, 50, 40}
	fmt.Println("intInsertionArray:", intInsertionArray)
	algo.InsertionSort(intInsertionArray)
	fmt.Println("intInsertionArray:", intInsertionArray)

	fmt.Println("======== Recursion Example ========")
	path := algo.RandomWalk(algo.Vector{X:1, Y:1}, 40)
	for _, vec := range path {
		fmt.Println(vec)
	}
	fmt.Println("Factorial: !5 ->", algo.Factorial(5))

	fmt.Println("======== Merge Sort ========")
	mergeThis := []int{5, 2, 1, 3, 9, 10, 11, 98, 1, 4, 12, 10, 30, 50, 40}
	fmt.Println(mergeThis)
	algo.MergeSort(mergeThis)
	fmt.Println(mergeThis)
}

