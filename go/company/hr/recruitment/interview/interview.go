package interview

import (
	"fmt"
	"strconv"
)

func ScheduleInterview(candidateName string) {
	fmt.Printf("Scheduling interview for %s\n", strconv.Quote(candidateName))
}

func ConductInterview(candidateName string) {
	fmt.Printf("Conducting interview for %s\n", strconv.Quote(candidateName))
}

func FeedbackInterview(candidateName, feedback string) {
	fmt.Printf("Feedback for %s: %s\n", strconv.Quote(candidateName), feedback)
}
