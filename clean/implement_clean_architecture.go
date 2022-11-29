package clean

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
)

/*
Clean architecture.
===================

Typically the data that crosses the boundaries is simple data structures.
Isolated, simple, data structures are passed across the boundaries.
The code in a layer, can only 'see' things that are defined in same layer or layer below it.
Code in the UseCases layer cannot see things defined in layer 4.

You talk inwards through simple data structures and talk outwards using interfaces.

(a) Entities
They encapsulate Enterprise wide business rules.
An entity can be an object with methods, or it can be a set of data structures and functions.
The entities could be used by different applications in the enterprise.

(b) Use Cases
They are application specific business rules.
They orchestrate the flow of data to and from the entities.
We do not expect changes in this layer to affect the entities.
We also do not expect this layer to be affected by changes to externalities such as the databases.

(c) Interface Adapters
They convert data from the format convenient for the use cases & entities, to the format convenient for the framework/drivers eg the Database.

(d) Frameworks and Drivers.
Generally you don’t write much code in this layer other than glue code that communicates to the next circle inwards.
This layer is where all the details go. The Web is a detail. The database is a detail.
Code in outer layers can use types/stuff defined in inner layers.
As an example, db can return types defined in the entities layer.

1. https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html
2. https://www.youtube.com/watch?v=C7MRkqP5NRI : Clean Architectures in Python
*/

// NOTE: In actual use, the four different components might/would be go packages each.

// 1. Entities
type Book struct {
	name   string
	author string
}

func newBook(name, author string) Book { return Book{name: name, author: author} }
func (b Book) String() string          { return fmt.Sprintf("Book{n: %s, a: %s}", b.name, b.author) }

// 2. Use Cases
func addBookUseCase(db dbInterface, name, author string) {
	b := newBook(name, author)
	db.save(b.String())
}

// 3. Interface Adapters
type dbInterface interface {
	save(content string)
	get(bookName string) Book
}

// 4. framework & drivers
func AddBookHandler(w http.ResponseWriter, _ *http.Request) {
	bookName := "A man of the people"
	author := "Chinua Achebe"

	db := database{path: "/tmp/clean_architecture_db.txt"}

	// Call use-case with simple data structures.
	// We cannot call, lower layers with types defined in upper layers.
	// As an example, we cannot call `addBookUseCase` with type defined in our web-framework.
	addBookUseCase(db, bookName, author)

	_, _ = io.WriteString(w, fmt.Sprintf("succesfully added book: %s \n", bookName))
}

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
