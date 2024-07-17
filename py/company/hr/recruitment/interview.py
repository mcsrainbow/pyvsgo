# interview module
def schedule_interview(candidate_name):
    print(f'''INFO: Scheduling interview for "{candidate_name}"''')

def conduct_interview(candidate_name):
    print(f'''INFO: Conducting interview for "{candidate_name}"''')

def feedback_interview(candidate_name, feedback):
    print(f'''INFO: Feedback for "{candidate_name}": {feedback}''')

# Example usage if this module is run directly
if __name__ == "__main__":
    candidate = "John Doe"
    schedule_interview(candidate)
    conduct_interview(candidate)
    feedback_interview(candidate, "Excellent")

