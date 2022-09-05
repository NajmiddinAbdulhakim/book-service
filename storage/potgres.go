package storage

import (
	"context"

	pb "github.com/NajmiddinAbdulhakim/ude/book-service/genproto"
	"github.com/jmoiron/sqlx"
)

type bookRepo struct {
	db *sqlx.DB
}

func NewBookRepo(db *sqlx.DB) *bookRepo {
	return &bookRepo{db}
}

func (r *bookRepo) Create(ctx context.Context, b *pb.BookReq) (*pb.BookRes, error) {
	var categoryId string

	query := `SELECT id FROM book_category WHERE category_name = $1`
	err := r.db.GetContext(ctx, &categoryId, query, b.Category)
	if err != nil {
		return nil, err
	}

	query = `INSERT INTO books (title, author_name, category_id)
	VALUES($1, $2, $3) 
	RETURNING id title, author_name`

	var book pb.BookRes
	err = r.db.QueryRowContext(
		ctx, query,
		b.Title, b.AuthorName, categoryId,
	).Scan(
		&book.Id,
		&book.Title,
		&book.AuthorName,
	)
	book.Category = b.Category

	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *bookRepo) GetById(ctx context.Context, id string) (*pb.BookRes, error) {
	query := `SELECT b.id, title, author_name, c.category_name  
	FROM books b
	INNER JOIN book_category c 
		ON b.category_id = c.id 
	WHERE b.id = $1`

	var book pb.BookRes
	err := r.db.GetContext(ctx, &book, query, id)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (r *bookRepo) Update(ctx context.Context, b *pb.UpdateBookReq) (bool, error) {
	var categoryId string

	query := `SELECT id FROM book_category WHERE category_name = $1`
	err := r.db.GetContext(ctx, &categoryId, query, b.CategoryId)
	if err != nil {
		return false, err
	}

	query = `UPDATE books 
	SET title = $1, author_name = $2, category_id = $3 
	WHERE id = $4`

	_, err = r.db.ExecContext(
		ctx, query,
		b.Title, b.AuthorName, categoryId, b.Id,
	)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *bookRepo) Delete(ctx context.Context, id string) (bool, error) {
	query := `DELETE FROM books 
	WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return false, err
	}

	return true, nil
}
