package main

import (
	"fmt"
	"pyvsgo/go/company/hr/recruitment"
	"pyvsgo/go/company/hr/recruitment/interview"
)

func main() {
	fmt.Printf("INFO: recruitment.RecruitmentTeam: %s\n", recruitment.RecruitmentTeam)

	candidate := "Jane Doe"
	interview.ScheduleInterview(candidate)
	interview.ConductInterview(candidate)
	interview.FeedbackInterview(candidate, "Good")
}
