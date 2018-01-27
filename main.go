package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A CSV file in the format: 'question,answer'") //Specify the csv filename in the terminal when about to run program (-csv="problems.csv").Useful when youre giving users the binary file and not source code of course.
	timeLimit := flag.Int("limit", 6, "the time limit for the quiz in seconds")                      // Allows the user to customize time limit
	flag.Parse()                                                                                     //read the flagged filename
	//Open and Read file
	file, err := os.Open(*csvFilename) //Open file
	if err != nil {
		exit(fmt.Sprintf("Failed to open file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll() //read file
	if err != nil {
		exit(fmt.Sprintf("Failed to read/parse file: %s\n", lines))
	}

	correct := 0 //initialize counter to check number of correct answers
	answerCount := 0
	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	//loop through each []problem in the [][]problems file
	for i, p := range problems {
		fmt.Printf("Problem %d). %s = ", i+1, p.que)
		answerCh := make(chan string) //channel to hold answer
		go func() {                   //goroutine to handle when an answer is entered
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer //whenever an answer is obtained, it is sent to the answer channel.
			answerCount++
		}() //call the goroutine

		select {
		case <-timer.C: // If we get a message from the timer channel (i.e once the timeLimit is reached), do the following:
			fmt.Println("\nSORRY, You are out of time!\nYou answered ", correct, "out of", len(problems), "questions correctly")
			return //leave the for loop (terminate Program)
		case answer := <-answerCh: //if we get a message from the answer channel
			if answer == p.ans { //if answer matches that in our csv file
				correct++
			}
			if answerCount == len(problems) {
				fmt.Println("\nYou answered ", correct, "out of", len(problems), "questions correctly")
				return
			}
		}
	}
	if correct == (len(problems)) {
		fmt.Println("WELL DONE!!!")
	}
} //end main

//parseLines returns a normal slice from a 2d slice file such as our problems.csv for easy iteration
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i].que = line[0]
		ret[i].ans = line[1]
	}
	return ret
}

//problem format
type problem struct {
	que string
	ans string
}

//Used after error discovery in exiting programme
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1) // status code 1 means something went wrong
}
