package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type MyInterface interface {
	getValue() any
	getIndex() uint
}

type returnValue struct {
	value       any
	whichThread uint
}

func (rV *returnValue) getValue() (string, error) {
	str, ok := (*rV).value.(string)
	if ok {
		return str, nil
	}
	return "", errors.New("error type, the value not string")
}

func (rV *returnValue) getIndex() uint {
	return (*rV).whichThread
}

func wait(NumberOfThreads int, ch chan returnValue, threadCounter uint) {
	timer := rand.Intn(NumberOfThreads + 1)
	time.Sleep(time.Duration(timer) * time.Second)
	ch <- returnValue{fmt.Sprintf("'Wait for %v'", timer), uint(threadCounter)}
}

func printThread(channel chan returnValue) {
	rV := <-channel

	str, err := rV.getValue()
	if err != nil {
		panic(fmt.Sprintf("error: %v", err))
	}
	fmt.Println(str)

	fmt.Println(rV.getIndex())
}

func main() {
	NumberOfThreads := uint(rand.Intn(10) + 1)
	fmt.Println("Start")

	channel := make(chan returnValue)

	for i := range int(NumberOfThreads) {
		go wait(i, channel, uint(i))
	}

	for range int(NumberOfThreads) {
		printThread(channel)
	}

	fmt.Println("Done")
}
