import random  # 导入随机数模块: 用于生成随机的彩票号码
import sys     # 导入系统模块: 用于控制台输出和退出程序

# 定义一个名为 Lottery 的类: 代表彩票
class Lottery:
    # 红球的最大值: 33
    RED_BALL_MAX = 33  
    # 蓝球的最大值: 16
    BLUE_BALL_MAX = 16  
    # 红球的数量: 6
    RED_BALL_COUNT = 6   
    # 总共的球数: 7 (6 个红球 + 1 个蓝球)
    TOTAL_BALL_COUNT = 7  

    # 构造函数, 初始化彩票对象: user_balls 是用户选择的号码
    def __init__(self, user_balls):
        # 初始化空列表存储生成的红球号码
        self.red_balls = []  
        # 初始化变量存储生成的蓝球号码
        self.blue_ball = None  
        # 存储用户选择的号码
        self.user_balls = user_balls  

    # 验证用户输入的号码: 检查格式和范围是否正确
    def validate_user_balls(self):
        # 检查用户输入的号码数量是否为 7
        if len(self.user_balls) != self.TOTAL_BALL_COUNT:
            print("NOTICE: 您必须提供Exactly 7 个球.")  
            exit(1)  # 退出程序, 状态码 1
        # 检查每个号码是否是正整数
        for ball in self.user_balls:
            if not isinstance(ball, int) or ball <= 0:
                print(f"NOTICE: 球 {ball} 不是有效的正整数.")  
                exit(1)  # 退出程序, 状态码 1
        # 分离红球和蓝球号码
        red_balls = self.user_balls[:self.RED_BALL_COUNT]
        blue_ball = self.user_balls[self.RED_BALL_COUNT]
        # 检查每个红球号码是否在有效范围内 (1-33)
        if any(ball > self.RED_BALL_MAX for ball in red_balls):
            print("NOTICE: 其中一个红球大于允许的最大值 (33).")  
            exit(1)  # 退出程序, 状态码 1
        # 检查蓝球号码是否在有效范围内 (1-16)
        if blue_ball > self.BLUE_BALL_MAX:
            print(f"NOTICE: 蓝球 {blue_ball} 大于允许的最大值 (16).")  
            exit(1)  # 退出程序, 状态码 1

    # 生成随机的彩票号码
    def generate_lottery_balls(self):
        # 随机选择 6 个不重复的红球号码 (1-33)
        self.red_balls = random.sample(range(1, self.RED_BALL_MAX + 1), self.RED_BALL_COUNT)
        # 排序红球号码
        self.red_balls.sort()  
        # 随机选择一个蓝球号码 (1-16)
        self.blue_ball = random.randint(1, self.BLUE_BALL_MAX)  
        # 返回生成的彩票号码 (6 个红球 + 1 个蓝球)
        return self.red_balls + [self.blue_ball]

# 主函数: 程序入口
def main():
    # 用户选择的号码
    user_balls = [1, 5, 10, 15, 16, 26, 9]  
    # 创建彩票对象
    lottery = Lottery(user_balls)
    # 验证用户输入的号码
    lottery.validate_user_balls()

    # 初始化尝试次数
    attempts = 1

    # 不断生成随机彩票号码, 直到中奖
    while True:
        # 生成随机彩票号码
        generated_balls = lottery.generate_lottery_balls()
        # 检查是否中奖 (用户选择的号码与生成的号码相同)
        if generated_balls == user_balls:
            print(f"\nNOTICE: 您终于中了 500W 人民币, 购买了 {attempts} 张彩票!")  
            break
        # 输出尝试次数
        sys.stdout.write(f"\rINFO: 您已经尝试了 {attempts} 次.")  
        sys.stdout.flush()  # 刷新输出缓冲区
        attempts += 1  # 增加尝试次数

# 如果脚本被直接运行 (不是导入), 执行主函数
if __name__ == "__main__":
    main()