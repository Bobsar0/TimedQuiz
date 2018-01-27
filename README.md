# TimedQuiz
Timed Quiz created in Go.  The program does the following:

- Reads in a quiz provided via a CSV file and gives the quiz to a user.
- Keeps track of correct and incorrect answers.
- Regardless of whether the answer is correct or wrong the next question is asked immediately afterwards.

- A timer is added whose default time limit should be 10 seconds, but is also be customizable via a flag. (-limit=10).
- Quiz stops as soon as the time limit has exceeded - even if the program is currently waiting on an answer from the end user.
- At the end of the quiz the program outputs the total number of questions correct and how many questions there were in total. Questions given invalid answers or unanswered are considered incorrect.

Inspiration from Jon Calhoun.
