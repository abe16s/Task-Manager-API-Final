### main.go
The entry point of the application.

### controllers/
Contains the library controller that handles user input and interactions.

### models/
Defines the data structures for books and members.

### services/
Implements the core library management functionality.

### docs/
Contains the documentation for the project.

### go.mod
Defines the module path and dependencies.

## Library Service
The library service (library_service.go) implements the core functionality of the library management system. It defines the LibraryManager interface and the Library struct.

### Library Service
The library service (library_service.go) implements the core functionality of the library management system. It defines the LibraryManager interface and the Library struct.

### LibraryManager Interface
The LibraryManager interface defines the following methods:

* AddBook(book models.Book)
* RemoveBook(bookID int)
* BorrowBook(bookID int, memberID int) error
* ReturnBook(bookID int, memberID int) error
* ListAvailableBooks() []models.Book
* ListBorrowedBooks(memberID int) []models.Book

### Library Struct
The Library struct contains two fields:

* Books map[int]*models.Book: A map of book IDs to book objects.
* Members map[int]*models.Member: A map of member IDs to member objects.
The Library struct implements the LibraryManager interface methods to manage the library's inventory and members.

## Library Controller
The library controller (library_controller.go) handles user input and interactions with the library service. It provides a console-based interface for users to manage books and memberships.

### StartConsole Function
The StartConsole function initiates the console-based interface and prompts the user for actions such as adding, removing, borrowing, and returning books, as well as listing available and borrowed books.

### continueUsing Function
The continueUsing function prompts the user to continue using the system or exit.

### addBook Function
The addBook function prompts the user to enter book details and adds the book to the library.

### removeBook Function
The removeBook function prompts the user to enter a book ID and removes the book from the library.

### borrowBook Function
The borrowBook function prompts the user to enter a book ID and member ID, and borrows the book for the specified member.

### returnBook Function
The returnBook function prompts the user to enter a book ID and member ID, and returns the book for the specified member.

### listAvailableBooks Function
The listAvailableBooks function lists all available books in the library.

### listBorrowedBooks Function
The listBorrowedBooks function prompts the user to enter a member ID and lists all books borrowed by the specified member.

## Models
The models define the data structures used in the library management system.

### Book
The Book struct is defined in models/book.go and contains the following fields:

* ID int
* Title string
* Author string
* Status string

### Member
The Member struct is defined in models/member.go and contains the following fields:

* ID int
* Name string
* BorrowedBooks []*Book LibraryManager Interface