import random
import sys
import time

class Lottery:
    RED_BALL_MAX = 33
    BLUE_BALL_MAX = 16
    RED_BALL_COUNT = 6
    TOTAL_BALL_COUNT = 7

    def __init__(self, user_balls):
        self.red_balls = []
        self.blue_ball = None
        self.user_balls = user_balls

    def validate_user_balls(self):
        if len(self.user_balls) != self.TOTAL_BALL_COUNT:
            print("NOTICE: You must provide exactly 7 balls.")
            exit(1)
        for ball in self.user_balls:
            if not isinstance(ball, int) or ball <= 0:
                print(f"NOTICE: Ball {ball} is not a valid positive integer.")
                exit(1)
        red_balls = self.user_balls[:self.RED_BALL_COUNT]
        blue_ball = self.user_balls[self.RED_BALL_COUNT]
        if any(ball > self.RED_BALL_MAX for ball in red_balls):
            print("NOTICE: One of the red balls is greater than the maximum allowed (33).")
            exit(1)
        if blue_ball > self.BLUE_BALL_MAX:
            print(f"NOTICE: Blue ball {blue_ball} is greater than the maximum allowed (16).")
            exit(1)

    def generate_lottery_balls(self):
        self.red_balls = random.sample(range(1, self.RED_BALL_MAX + 1), self.RED_BALL_COUNT)
        self.red_balls.sort()
        self.blue_ball = random.randint(1, self.BLUE_BALL_MAX)
        return self.red_balls + [self.blue_ball]

def main():
    user_balls = [1, 5, 10, 15, 16, 26, 9]  # 用户选择的球
    lottery = Lottery(user_balls)
    lottery.validate_user_balls()

    attempts = 1

    while True:
        generated_balls = lottery.generate_lottery_balls()
        if generated_balls == user_balls:
            print(f"\nNOTICE: You finally got 500W RMB after buying {attempts} tickets!")
            break
        sys.stdout.write(f"\rINFO: You have tried {attempts} times.")
        sys.stdout.flush()
        attempts += 1

if __name__ == "__main__":
    main()

