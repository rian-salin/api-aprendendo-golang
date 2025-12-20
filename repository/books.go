package repository

import (
    "Api-Aula1/models"
    "database/sql"
)

type BooksRepo struct {
    db *sql.DB
}


func NewBooksRepo(db *sql.DB) *BooksRepo {
    return &BooksRepo{db}
}

func (repository BooksRepo) Create(book models.Book) (uint64, error) {
    statement, err := repository.db.Prepare(
        "INSERT INTO treehousedb.books (title, author, gender, user_id) VALUES (?, ?, ?, ?)",
    )
    if err != nil {
        return 0, err
    }
    defer statement.Close()

    result, err := statement.Exec(book.Title, book.Author, book.Gender, book.UserID)
    if err != nil {
        return 0, err
    }

    lastID, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return uint64(lastID), nil
}

func (repository BooksRepo) FetchByUserID(userID uint64) ([]models.Book, error) {
    rows, err := repository.db.Query("SELECT id, title, author, gender, user_id FROM treehousedb.books WHERE user_id = ?", userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var books []models.Book

    for rows.Next() {
        var book models.Book
        if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Gender, &book.UserID); err != nil {
            return nil, err
        }
        books = append(books, book)
    }

    return books, nil
}

func (repository BooksRepo) Delete(bookID uint64) error {
    statement, err := repository.db.Prepare("DELETE FROM books WHERE id = ?")
    if err != nil {
        return err
    }
    defer statement.Close()

    if _, err = statement.Exec(bookID); err != nil {
        return err
    }

    return nil
}

func (repository BooksRepo) Update(bookID uint64, book models.Book) error {
    statement, err := repository.db.Prepare("UPDATE books SET title = ?, author = ?, gender = ? WHERE id = ?")
    if err != nil {
        return err
    }
    defer statement.Close()

    if _, err = statement.Exec(book.Title, book.Author, book.Gender, bookID); err != nil {
        return err
    }

    return nil
}