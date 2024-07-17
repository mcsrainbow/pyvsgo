# 创建一个标准的 dict
ordered_dict = {
    "name": "John",
    "age": 30,
    "is_student": False,
    "scores": [90, 85, 92]
}

# 获取所有键和值
print("ordered_dict Keys:", list(ordered_dict.keys()))
print("ordered_dict Values:", list(ordered_dict.values()))

# 获取单个键的值
if "name" in ordered_dict:
    print("name:", ordered_dict["name"])

if "nationality" in ordered_dict:
    print("nationality:", ordered_dict["nationality"])
else:
    print("No such key: nationality")

# 循环打印所有键和值
print("All keys and values in ordered_dict:")
for key, value in ordered_dict.items():
    print(f"{key}: {value}")

