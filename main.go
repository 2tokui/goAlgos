package main

import (
	"fmt"
	"github.com/2tokui/goAlgos/algo"
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

	fmt.Println("======== Qeueue ========")
	queue := algo.NewQueue()
	queue.Enqueue(1)
	queue.Enqueue("this guy")
	queue.Enqueue(false)
	queue.Enqueue(1337)

	fmt.Println("Peek:", queue.Peek())
	queue.Dequeue()
	fmt.Println("Peek:", queue.Peek())
	queue.Dequeue()
	fmt.Println("Peek:", queue.Peek())
	queue.Dequeue()
	fmt.Println("Peek:", queue.Peek())
	queue.Dequeue()
	fmt.Println("Peek:", queue.Peek())
}

