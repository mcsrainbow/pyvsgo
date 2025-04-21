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

# 使用pop删除键'age'
print("\nDeleting key 'age':")
removed_value = ordered_dict.pop("age", None)
print(f"Removed value: {removed_value}")

print("ordered_dict Keys after deletion:", list(ordered_dict.keys()))
print("ordered_dict Values after deletion:", list(ordered_dict.values()))

# 尝试使用pop删除不存在的键
print("\nAttempting to delete non-existent key 'gender':")
removed_value = ordered_dict.pop("gender", None)
result = removed_value is not None  # 如果键存在并被删除则为True
print(f"Delete result: {result}")
print(f"Removed value: {removed_value}")
