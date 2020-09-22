package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//create a custom datatype for queue
type logEntry struct {
	Type     string
	UniqueID int
	Data     string
}

// define path
var path = "../data.log"

//main function
func main() {
	for a := 0; a < 6; a++ {
		go Routine()
		time.Sleep(2 * time.Second)
	}

}

/*data generator*/

func datagen() string {
	// random data generation for queue

	var typ = [5]string{"A", "B", "C"}
	var uid = [3]int{1, 2, 3}
	var data = [9]string{"wazzup", "how to swim", "hotwheels101", "have a red bull", "bundesliga", "hola senorita", "viel spaß", "das is verrückt", "corona is a beer, really!!"}

	logdata1 := logEntry{Type: typ[rand.Intn(len(typ))], UniqueID: uid[rand.Intn(len(uid))], Data: data[rand.Intn(len(data))]}

	//fmt.Println(logdata)
	logdatastring := string(logdata1.Type) + ":" + strconv.Itoa(logdata1.UniqueID) + ":" + string(logdata1.Data) + "\n"

	return logdatastring
}
func Routine() {

	// open the file
	fmt.Println("Opening a file ")
	file, err := os.OpenFile(path, os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("error is", err)
	}
	// recursive close
	defer file.Close()
	// write and dequeue frim a queue
	var queue []string

	// write to file
	//select random data to create a string from arrays
	var logdatastring = datagen()
	queue = enqueue(queue, logdatastring)

	_, err = file.WriteString(queue[len(queue)-1])
	time.Sleep(2 * time.Second)
	queue = dequeue(queue)
	if err != nil {
		fmt.Println("error is", err)

	}
	fmt.Println("q is", queue)

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
