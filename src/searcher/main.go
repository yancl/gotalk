package main

import (
	"components"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var (
	User  = components.FakeSearcher("user")
	Topic = components.FakeSearcher("topic")
	Group = components.FakeSearcher("group")
	Voice = components.FakeSearcher("voice")
)

func SerializedSearch(query string) (results []components.Result) {
	results = append(results, User(query))
	results = append(results, Topic(query))
	results = append(results, Group(query))
	results = append(results, Voice(query))
	return
}

func ConcurrentSearch(query string) (results []components.Result) {
	c := make(chan components.Result)
	go func() { c <- User(query) }()
	go func() { c <- Topic(query) }()
	go func() { c <- Group(query) }()
	go func() { c <- Voice(query) }()

	for i := 0; i < 4; i++ {
		result := <-c
		results = append(results, result)
	}
	return
}

func ConcurrentTimeoutSearch(query string) (results []components.Result) {
	c := make(chan components.Result)
	go func() { c <- User(query) }()
	go func() { c <- Topic(query) }()
	go func() { c <- Group(query) }()
	go func() { c <- Voice(query) }()

	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 4; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("timed out")
			return
		}
	}
	return
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s cmd(serialized|concurrent|ctimeout)\n", os.Args[0])
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		usage()
	}

	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	cmd := os.Args[1]
	var results []components.Result
	if cmd == "serialized" {
		results = SerializedSearch("golang")
	} else if cmd == "concurrent" {
		results = ConcurrentSearch("golang")
	} else if cmd == "ctimeout" {
		results = ConcurrentTimeoutSearch("golang")
	} else {
		usage()
	}
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
