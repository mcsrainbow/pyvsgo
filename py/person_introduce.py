class Person:
    def __init__(self, name: str, age: int, gender: str) -> None:
        self.name = name
        self.age = age
        self.gender = gender
    
    def introduce(self) -> str:
        return f"Hello! My name is {self.name}, I am {self.age} years old and I am {self.gender}."

if __name__ == "__main__":
    # create a Person instance
    user = Person("John", 25, "male")
    
    # call the introduce method
    print(user.introduce())
