import company.hr.recruitment as recruitment
import company.hr.recruitment.interview as interview

print(f"INFO: recruitment.recruitment_team: {recruitment.recruitment_team}")

interview.schedule_interview("Jane Doe")
interview.conduct_interview("Jane Doe")
interview.feedback_interview("Jane Doe", "Good")
