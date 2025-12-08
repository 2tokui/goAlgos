package main

import "time"

func ConcurrentProcedure(kokoChan chan string) string {
	go func() {
		time.Sleep(time.Second * 1)
		kokoChan <- "haiii :3"
	}()
	return <-kokoChan
}
