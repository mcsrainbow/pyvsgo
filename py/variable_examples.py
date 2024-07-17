# 常量, 全局变量
GLOBAL_CONST = "I am a global constant visible to external packages"

# 普通全局变量
global_var = "I am a global variables visible to external packages"

class MyClass:
    # 类变量, 类似于静态变量
    static_var = 10

    def __init__(self, instance_var):
        # 实例变量
        self.instance_var = instance_var

    def increment_class_var(self):
        MyClass.static_var += 1

    def increment_instance_var(self):
        self.instance_var += 1

def my_function():
    # 局部变量
    local_var = "I am a local variable"
    print("my_function() local_var:", local_var)

def main():
    # 使用示例
    print()
    print("GLOBAL_CONST:", GLOBAL_CONST)
    print()
    print("global_var:", global_var)

    # 动态变量, 可以在运行时改变类型
    dynamic_var = 5
    print("dynamic_var:", dynamic_var)
    dynamic_var = "Now I'm a string"
    print("dynamic_var:", dynamic_var)

    # 列表
    my_list = [1, 2, 3, 4, 5]

    # 字典
    my_dict = {"name": "John", "age": 30}

    # 字符串
    my_string = "Hello, World!"

    # 使用示例
    my_function()
    print("my_list:", my_list)
    print("my_dict:", my_dict)
    print("my_string:", my_string)

    # 创建类实例
    obj = MyClass(instance_var=1)

    print("obj.static_var:", obj.static_var)
    print("obj.instance_var:", obj.instance_var)
    # 增加类变量
    obj.increment_class_var()
    print("obj.increment_class_var() obj.static_var:", obj.static_var)
    # 增加实例变量
    obj.increment_instance_var()
    print("obj.increment_instance_var() obj.instance_var:", obj.instance_var)

if __name__ == "__main__":
    main()
