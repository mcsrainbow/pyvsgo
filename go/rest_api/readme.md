## Go REST framework

![](screenshot.jpg)

```bash
mkdir rest_api
mkdir rest_api/{config,controllers,models,routes}

cd rest_api
vim main.go
vim config/config.go
vim controllers/bookController.go
vim models/book.go
vim routes/bookRoutes.go

go mod tidy

go run main.go
```

```bash
# add books
curl -X POST http://localhost:8080/api/books \
     -H "Content-Type: application/json" \
     -d '{"title": "The Go Programming Language", "author": "Alan A. A. Donovan, Brian Kernighan"}'

curl -X POST http://localhost:8080/api/books \
     -H "Content-Type: application/json" \
     -d '{"title": "Rich Dad Poor Dad", "author": "Robert Kiyosaki, Sharon Lechter"}'

curl -X POST http://localhost:8080/api/books \
     -H "Content-Type: application/json" \
     -d '{"title": "How an Economy Grows and Why It Crashes", "author": "Peter Schiff"}'

curl -X POST http://localhost:8080/api/books \
     -H "Content-Type: application/json" \
     -d '{"title": "Nonviolent Communication: A Language of Life", "author": "Marshall Rosenberg"}'

curl -X POST http://localhost:8080/api/books \
     -H "Content-Type: application/json" \
     -d '{"title": "Influence: The Psychology of Persuasion", "author": "Robert Cialdini"}'

curl -X POST http://localhost:8080/api/books \
     -H "Content-Type: application/json" \
     -d '{"title": "Educated: A Memoir", "author": "Tara Westover"}'

curl -X POST http://localhost:8080/api/books \
     -H "Content-Type: application/json" \
     -d '{"title": "Pomodoro Technique Illustrated", "author": "Staffan Noteberg"}'

curl -X POST http://localhost:8080/api/books \
     -H "Content-Type: application/json" \
     -d '{"title": "PEAK: Secrets from the New Science of Expertise", "author": "K. Anders Ericsson, Robert Pool"}'

# get all books
curl -X GET http://localhost:8080/api/books

# get a book
curl -X GET http://localhost:8080/api/books/1

# update a book
curl -X PUT http://localhost:8080/api/books/1 \
     -H "Content-Type: application/json" \
     -d '{"title": "The Go Programming Language, Updated Edition", "author": "Alan A. A. Donovan, Brian Kernighan"}'

# delete a book
curl -X DELETE http://localhost:8080/api/books/8
```
