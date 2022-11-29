package myio

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
)

/*
My IO architecture.
===================

(a) io.
Has all the functions that do I/O.
Interfaces are not defined here since it's against Go-ism, instead they are defined in (b)

(b) use cases.
They are application specific business rules.
They orchestrate the flow of data to and from the entities.
They call into (a).
They also declare interfaces for calling into (a).
Note: these interfaces are declared on the consumer side as opposed to producer side.

(c) Handlers
These are one line and call into (b)
*/

// NOTE: In actual use, the three different components might/would be go packages each.

// 1. IO
type database struct {
	path string
}

func (d database) save(content string) {
	err := os.WriteFile(d.path, []byte(content), fs.ModeAppend)
	_ = err
}

func (d database) get(bookName string) Book {
	// fetch from file & marshall to a book type.
	return Book{}
}

// 2. Use Cases
type Book struct {
	name   string
	author string
}

func newBook(name, author string) Book { return Book{name: name, author: author} }
func (b Book) String() string          { return fmt.Sprintf("Book{n: %s, a: %s}", b.name, b.author) }

type dbInterface interface {
	save(content string)
	get(bookName string) Book
}

func addBookUseCase(db dbInterface, name, author string) {
	b := newBook(name, author)
	db.save(b.String())
}

// 3. Handlers
func AddBookHandler(w http.ResponseWriter, _ *http.Request) {
	bookName := "Africa Kills Her Sun"
	author := "Ken Saro Wiwa"

	db := database{path: "/tmp/clean_architecture_db.txt"}

	// Call use-case with simple data structures.
	// We cannot call, lower layers with types defined in upper layers.
	// As an example, we cannot call `addBookUseCase` with type defined in our web-framework.
	addBookUseCase(db, bookName, author)

	_, _ = io.WriteString(w, fmt.Sprintf("succesfully added book: %s \n", bookName))
}
