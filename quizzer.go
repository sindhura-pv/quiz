package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"log"
	"io"
	"flag"
	"time"
)

//var userInput chan string

func main(){

	limit := flag.Int("limit", 100000, "limit per question")
	flag.Parse()

	fmt.Print(*limit)
	f, err := os.Open("problems.csv")
	if err != nil{
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(f)
	i := 1
	count := 0
	q := 0
	userInput := make(chan string, 1)

	go readInput(userInput)
	fl := false

	for{
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Question %d) %s is:", i, record[0])
		select{
			case m:= <-userInput:
				if m == string(record[1]) {
					count++
			}
			case <-time.After(time.Duration(*limit)*time.Second):
				fmt.Println("\n Time is over!")
				fl = true
		}

		q++
		i++

		if fl{
			break
		}
	}

	fmt.Printf("You answered %d out of %d questions correctly. ", count, q)
}

func readInput(input chan string){

	for{
		var s string
		fmt.Scanln(&s)
		input<-s
	}
}
