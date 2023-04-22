package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strconv"
	"time"
)

type ControlMessage struct {
	Target string
	Count  int
}

func main() {
	controlChannel := make(chan ControlMessage)
	workerCompleteChan := make(chan bool)
	pollChannel := make(chan chan bool)
	workerActive := false

	go func() {
		for {
			select {
			case respChan := <-pollChannel:
				respChan <- workerActive
			case msg := <-controlChannel:
				workerActive = true
				go doStuff(msg, workerCompleteChan)
			case status := <-workerCompleteChan:
				workerActive = status
			}
		}
	}()

	http.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		count, err := strconv.ParseInt(r.FormValue("count"), 10, 32)
		if err != nil {
			log.Printf("Error parsing count: %v", err)
			return
		}
		msg := ControlMessage{
			Target: r.FormValue("target"),
			Count:  int(count),
		}
		controlChannel <- msg
		fmt.Fprintf(w, "Control message issued for Target %s\n", html.EscapeString(r.FormValue("target")))
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		reqChan := make(chan bool)
		pollChannel <- reqChan
		timeout := time.NewTimer(5 * time.Second)
		select {
		case result := <-reqChan:
			if result {
				fmt.Fprintln(w, "ACTIVE")
			} else {
				fmt.Fprintln(w, "INACTIVE")
			}
		case <-timeout.C:
			fmt.Fprintln(w, "INACTIVE")
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func doStuff(msg ControlMessage, completeChan chan<- bool) {
	defer func() {
		completeChan <- false
	}()

	for i := 0; i < msg.Count; i++ {
		fmt.Printf("Target: %s, Count: %d\n", msg.Target, i)
		time.Sleep(time.Second)
	}

	completeChan <- true
}
