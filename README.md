# Go-Bookstore

Go-Bookstore is a simple RESTful API written in Go that connects to a MySQL database to store and retrieve data. It is a simple application that can be used as a starting point for more complex Go applications. It is written in a way that is easy to understand and extend.

## Getting Started

To get started with Go-Bookstore, you need to have Go and MySQL installed on your system. You also need to have the following Go packages installed:

* [Gorilla Mux](https://github.com/gorilla/mux)
* [Go MySQL Driver](https://github.com/jinzhu/gorm/dialects/mysql)
* [GORM](https://github.com/jinzhu/gorm/)

You can install these packages using the following commands:

```
go get -u github.com/gorilla/mux
go get -u github.com/go-sql-driver/mysql
go get -u github.com/jinzhu/gorm
```

Once you have installed the required packages, you can clone the Go-Bookstore repository and run the application using the following commands:

```
git clone https://github.com/suhel-kap/go-bookstore.git
cd go-bookstore
go run main.go
```

This will start the Go-Bookstore API on port 8080.

## API Endpoints

Go-Bookstore provides the following API endpoints:

### GET /book

This endpoint returns a list of all the books in the database.

### GET /book/{id}

This endpoint returns the details of a specific book. The `id` parameter specifies the ID of the book.

### POST /book

This endpoint allows you to add a new book to the database. The details of the book are passed as JSON in the request body.

### PUT /book/{id}

This endpoint allows you to update the details of a specific book. The `id` parameter specifies the ID of the book. The new details of the book are passed as JSON in the request body.

### DELETE /book/{id}

This endpoint allows you to delete a specific book. The `id` parameter specifies the ID of the book.

## Database Configuration

Go-Bookstore uses a MySQL database to store and retrieve data. The database configuration is specified in the `config.go` file. By default, the application connects to a MySQL database running on `localhost`. You can modify the database configuration by changing the values of the url in the `Connect` function.

## License

Go-Bookstore is licensed under the MIT License. See the `LICENSE` file for more information.

## Acknowledgements

Go-Bookstore uses the following open source packages:

* [Gorilla Mux](https://github.com/gorilla/mux)
* [Go MySQL Driver](https://github.com/jinzhu/gorm/dialects/mysql)
* [GORM](https://github.com/jinzhu/gorm/)