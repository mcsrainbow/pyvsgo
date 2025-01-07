class Person:
    def __init__(self, name: str, age: int, gender: str) -> None:
        self.name: str = name
        self.age: int = age
        self.gender: str = gender
    
    def introduce(self) -> str:
        return f"Hello! My name is {self.name}, I am {self.age} years old and I am {self.gender}."

if __name__ == "__main__":
    # create a Person instance
    user: Person = Person("John", 25, "male")
    
    # call the introduce method
    print(user.introduce())
