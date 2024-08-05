package services

import (
	"fmt"
	"library_management/models"
)

type LibraryManager interface {
	AddBook(book models.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID int, memberID int) error
	ReturnBook(bookID int, memberID int) error
	ListAvailableBooks() []models.Book
	ListBorrowedBooks(memberID int) []models.Book
}

type Library struct {
	Books map[int]*models.Book
	Members map[int]*models.Member
}

func (l *Library) AddBook(book models.Book) {
	l.Books[book.ID] = &book
}

func (l *Library) RemoveBook(bookID int) {
	delete(l.Books, bookID)
}

func (l *Library) BorrowBook(bookID int, memberID int) error {
	book, b_ok := l.Books[bookID]
	member, m_ok := l.Members[memberID]

	if !b_ok {
		return fmt.Errorf("book not found with bookID %d", bookID)
	} else if !m_ok {
		return fmt.Errorf("member not found with memberID %d", memberID)
	} else if book.Status == "Borrowed" {
		return fmt.Errorf("the requested bookID %d is already borrowed", bookID)
	}
	member.BorrowedBooks = append(member.BorrowedBooks, book)
	book.Status = "Borrowed"
	return nil
}

func (l *Library) ReturnBook(bookID int, memberID int) error {
	member, m_ok := l.Members[memberID]
	if !m_ok {
		return fmt.Errorf("member not found with memberID %d", memberID)
	}
	for i, val := range member.BorrowedBooks {
		if val.ID == bookID {
			val.Status = "Available"
			member.BorrowedBooks = append(member.BorrowedBooks[:i], member.BorrowedBooks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("didn't find any borrowed book with id %d", bookID)
}

func (l *Library) ListAvailableBooks() []models.Book {
	availableBooks := make([]models.Book, 0)
	for _, book := range l.Books {
		if book.Status == "Available" {
			availableBooks = append(availableBooks, *book)
		}
	}
	return availableBooks
}

func (l *Library) ListBorrowedBooks(memberID int) []models.Book {
	member, m_ok := l.Members[memberID]
	borrowedBooks := make([]models.Book, 0)
	if !m_ok {
		return borrowedBooks
	}
	for _, book := range member.BorrowedBooks {
		borrowedBooks = append(borrowedBooks, *book)
	}
	return borrowedBooks
}