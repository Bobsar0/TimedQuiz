/*Create a program that will read in a quiz provided via a CSV file and will then give the quiz to a user.
Keep track of how many questions they get right and how many they get incorrect.
Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.
*/
package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A CSV file in the format: 'question,answer'") //Specify the csv filename in the terminal when about to run program (-csv="problems.csv").Useful when youre giving users the binary file and not source code of course.
	flag.Parse()                                                                                     //read the flagged filename
	file, err := os.Open(*csvFilename)                                                               //Open file
	if err != nil {
		exit(fmt.Sprintf("Failed to open file: %s\n", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll() //read file
	if err != nil {
		exit(fmt.Sprintf("Failed to read/parse file: %s\n", lines))
	}
	correct := 0 //initialize counter to check number of correct answers
	problems := parseLines(lines)
	for i, p := range problems {
		fmt.Printf("Problem %d). %s = ", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			correct++
		}
	}
	fmt.Println("You answered ", correct, "out of", len(problems), "questions correctly")
	if correct == (len(problems)) {
		fmt.Println("WELL DONE!!!")
	}
} //end main

//parseLines returns a normal slice from a 2d slice file such as our problems.csv for easy iteration
func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i].q = line[0]
		ret[i].a = line[1]
	}
	return ret
}

//problem format
type problem struct {
	q string
	a string
}

//Used after error discovery in exiting programme
func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1) // status code 1 means something went wrong
}
