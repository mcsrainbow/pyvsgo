# 使用 for 循环遍历 list
my_list = [1, 2, 3, 4, 5]
print("Using 'for' loop to iterate through list:")
for item in my_list:
    print(item)

# 使用 while 循环遍历 list
print("\nUsing 'while' loop to iterate through list:")
index = 0
while index < len(my_list):
    print(my_list[index])
    index += 1

# 使用 for 循环遍历 dict
my_dict = {
    "name": "John",
    "age": 30,
    "is_student": False,
    "scores": [90, 85, 92]
}
print("\nUsing 'for' loop to iterate through dict:")
for key, value in my_dict.items():
    print(f"Key: {key}, Value: {value}")

# 使用 while 循环遍历 dict
print("\nUsing 'while' loop to iterate through dict:")
keys = list(my_dict.keys())
index = 0
while index < len(keys):
    key = keys[index]
    print(f"Key: {key}, Value: {my_dict[key]}")
    index += 1

# Python 3.10 及以上版本的 match-case 语句
def match_case_example(value):
    match value:
        case 1:
            return "One"
        case 2:
            return "Two"
        case 3:
            return "Three"
        case _:
            return "Other"

print("\nUsing 'match-case' statement:")
print(match_case_example(1))
print(match_case_example(4))

# 死循环示例
print("\nInfinite loop example:")
while True:
    print("This is an infinite loop. Use 'break' to exit.")
    break  # 添加 break 以防止实际运行时进入死循环

# 使用 break 退出循环的示例
print("\nUsing 'break' to exit the loop at 'number 5':")
for i in range(10):
    if i == 5:
        break  # 等于5时退出循环
    print(i)

# 使用 pass 在循环中的示例
print("\nUsing 'pass' in the loop at 'even numbers':")
for i in range(10):
    if i % 2 == 0:
        pass  # 偶数时什么都不做
    else:
        print(i)  # 奇数时打印

