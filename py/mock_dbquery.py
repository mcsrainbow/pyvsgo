from abc import ABC, abstractmethod

class Database(ABC):
    @abstractmethod
    def query(self, query: str) -> str:
        pass

class MockDatabase(Database):
    def query(self, query: str) -> str:
        return "mock result"

class RealDatabase(Database):
    def query(self, query: str) -> str:
        return "real result"

def execute_query(db: Database, query: str):
    try:
        result = db.query(query)
        print("Result:", result)
    except Exception as e:
        print("Error:", e)

if __name__ == "__main__":
    # 使用 MockDatabase
    mock_db = MockDatabase()
    execute_query(mock_db, "SELECT * FROM table")  # Output: mock result

    # 使用 RealDatabase
    real_db = RealDatabase()
    execute_query(real_db, "SELECT * FROM table")  # Output: real result
