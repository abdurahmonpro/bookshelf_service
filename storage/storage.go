package storage

import (
	"book/genproto/book_service"
	"book/models"

	"context"
)

type StorageI interface {
	CloseDB()
	Book() BookRepoI
}

type BookRepoI interface {
	Create(context.Context, *book_service.CreateBook) (*book_service.BookPK, error)
	GetByPKey(context.Context, *book_service.BookPK) (*book_service.Book, error)
	GetAll(context.Context, *book_service.BookListRequest) (*book_service.BookListResponse, error)
	Update(context.Context, *book_service.UpdateBook) (int64, error)
	UpdatePatch(context.Context, *models.UpdatePatchRequest) (int64, error)
	Delete(context.Context, *book_service.BookPK) error
	GetBookByTitle(context.Context, *book_service.BookByTitle) (*book_service.Book, error)
}
