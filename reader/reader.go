package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// define a queue
var queue []string

// read the file, wait for input, depending on the writer give it to routine and clear the queue
func decesion(filename string) {
	f, err := os.Open(filename) //open file
	if err != nil {
		panic(err)
	}
	defer f.Close()         // close file once function is exited
	r := bufio.NewReader(f) // a reader object for opened file
	for {
		for line, prefix, err := r.ReadLine(); err != io.EOF; line, prefix, err = r.ReadLine() {

			if prefix {
				fmt.Println("big line") // if line is too big to read
			}

			queue := enqueue(queue, string(line)) // take each incoming write to the queue
			writer, content := logic(queue)       // split the Id of writer and also the retuen the content
			if writer == "A" {
				go readers(writer, content) //a routine given to writer A
				queue = dequeue(queue)
				fmt.Println(queue)
			} else if writer == "B" {
				go readers(writer, content)
				queue = dequeue(queue)
				fmt.Println(queue)
			} else if writer == "" {
				go readers(writer, content)
				queue = dequeue(queue)
				fmt.Println(queue)
			}

		}

	}

}

func enqueue(queue []string, element string) []string {
	queue = append(queue, element) // Simply append to enqueue.
	fmt.Println("Enqueued:", element)
	return queue
}

func dequeue(queue []string) []string {
	element := queue[0] // The first element is the one to be dequeued.
	fmt.Println("Dequeued:", element)
	return queue[1:] // Slice off the element once it is dequeued.
}
func logic(queue []string) (string, string) {
	// give the queue as input check whic writer it is give it to writer routine and deque
	var lastval string = queue[len(queue)-1]
	var firstval string = lastval[0:1]
	//fmt.Println(firstval)
	return firstval, lastval

}

func readers(writer string, content string) {
	fmt.Println("Reading contents of writer:", writer)
	return
}

func main() {
	decesion("../data.log")

}
