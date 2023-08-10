package postgres

import (
	"book/genproto/book_service"
	"book/models"
	"book/pkg/helper"
	"errors"
	"fmt"

	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type BookRepo struct {
	db *pgxpool.Pool
}

func NewBookRepo(db *pgxpool.Pool) *BookRepo {
	return &BookRepo{
		db: db,
	}
}

func (u *BookRepo) Create(ctx context.Context, req *book_service.CreateBook) (resp *book_service.BookPK, err error) {
	id := uuid.New().String()

	query := `
		INSERT INTO "book" (
			"id",
			"isbn",
			"title",
			"cover",
			"author",
			"published",
			"pages",
			"status",
			"created_at",
			"updated_at"
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())
	`

	_, err = u.db.Exec(
		ctx,
		query,
		id,
		req.Isbn,
		req.Title,
		req.Cover,
		req.Author,
		req.Published,
		req.Pages,
		req.Status,
	)
	if err != nil {
		return nil, err
	}

	return &book_service.BookPK{Id: id}, nil
}

func (u *BookRepo) GetByPKey(ctx context.Context, req *book_service.BookPK) (Book *book_service.Book, err error) {
	query := `
		SELECT
			"id",
			"isbn",
			"title",
			"cover",
			"author",
			"published",
			"pages",
			"status",
			"created_at",
			"updated_at"
		FROM "book"
		WHERE "id" = $1
	`

	var (
		id        sql.NullString
		isbn      sql.NullString
		title     sql.NullString
		cover     sql.NullString
		author    sql.NullString
		published sql.NullString
		pages     sql.NullInt32
		status    sql.NullInt32
		created   sql.NullString
		updated   sql.NullString
	)

	err = u.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&isbn,
		&title,
		&cover,
		&author,
		&published,
		&pages,
		&status,
		&created,
		&updated,
	)
	if err != nil {
		return Book, err
	}

	Book = &book_service.Book{
		Id:        id.String,
		Isbn:      isbn.String,
		Title:     title.String,
		Cover:     cover.String,
		Author:    author.String,
		Published: published.String,
		Pages:     int32(pages.Int32),
		Status:    int32(status.Int32),
		CreatedAt: created.String,
		UpdatedAt: updated.String,
	}

	return
}

func (u *BookRepo) GetBookByTitle(ctx context.Context, req *book_service.BookByTitle) (Book *book_service.Book, err error) {
	query := `
		SELECT
			"id",
			"isbn",
			"title",
			"cover",
			"author",
			"published",
			"pages",
			"status",
			"created_at",
			"updated_at"
		FROM "book"
		WHERE "title" ILIKE '%' || $1 || '%'
		LIMIT 1
	`

	row := u.db.QueryRow(ctx, query, req.Title)

	var (
		id        sql.NullString
		isbn      sql.NullString
		title     sql.NullString
		cover     sql.NullString
		author    sql.NullString
		published sql.NullString
		pages     sql.NullInt32
		status    sql.NullInt32
		created   sql.NullString
		updated   sql.NullString
	)

	err = row.Scan(
		&id,
		&isbn,
		&title,
		&cover,
		&author,
		&published,
		&pages,
		&status,
		&created,
		&updated,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Book not found
		}
		return nil, err
	}

	Book = &book_service.Book{
		Id:        id.String,
		Isbn:      isbn.String,
		Title:     title.String,
		Cover:     cover.String,
		Author:    author.String,
		Published: published.String,
		Pages:     int32(pages.Int32),
		Status:    int32(status.Int32),
		CreatedAt: created.String,
		UpdatedAt: updated.String,
	}

	return
}

func (u *BookRepo) GetAll(ctx context.Context, req *book_service.BookListRequest) (resp *book_service.BookListResponse, err error) {
	resp = &book_service.BookListResponse{}

	var (
		query  string
		limit  = ""
		offset = " OFFSET 0 "
		params = make(map[string]interface{})
		filter = " WHERE TRUE "
		sort   = " ORDER BY created_at DESC"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			"id",
			"isbn",
			"title",
			"cover",
			"author",
			"published",
			"pages",
			"status",
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS'),
			TO_CHAR(updated_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM "book"
	`

	if len(req.GetSearch()) > 0 {
		filter += " AND (title || ' ' || author) ILIKE '%' || '" + req.Search + "' || '%' "
	}
	if req.GetLimit() > 0 {
		limit = " LIMIT :limit"
		params["limit"] = req.Limit
	}
	if req.GetOffset() > 0 {
		offset = " OFFSET :offset"
		params["offset"] = req.Offset
	}

	query += filter + sort + offset + limit

	query, args := helper.ReplaceQueryParams(query, params)
	rows, err := u.db.Query(ctx, query, args...)
	if err != nil {
		return resp, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id        sql.NullString
			isbn      sql.NullString
			title     sql.NullString
			cover     sql.NullString
			author    sql.NullString
			published sql.NullString
			pages     sql.NullInt32
			status    sql.NullInt32
			created   sql.NullString
			updated   sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&isbn,
			&title,
			&cover,
			&author,
			&published,
			&pages,
			&status,
			&created,
			&updated,
		)
		if err != nil {
			return resp, err
		}

		resp.Books = append(resp.Books, &book_service.Book{
			Id:        id.String,
			Isbn:      isbn.String,
			Title:     title.String,
			Cover:     cover.String,
			Author:    author.String,
			Published: published.String,
			Pages:     int32(pages.Int32),
			Status:    int32(status.Int32),
			CreatedAt: created.String,
			UpdatedAt: updated.String,
		})
	}

	return resp, nil
}

func (u *BookRepo) Update(ctx context.Context, req *book_service.UpdateBook) (rowsAffected int64, err error) {
	query := `
		UPDATE "book"
		SET
			"title" = $1,
			"cover" = $2,
			"author" = $3,
			"published" = $4,
			"pages" = $5,
			"status" = $6,
			"updated_at" = NOW()
		WHERE "id" = $7
	`

	result, err := u.db.Exec(ctx, query,
		req.Title,
		req.Cover,
		req.Author,
		req.Published,
		req.Pages,
		req.Status,
		req.Id,
	)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (u *BookRepo) UpdatePatch(ctx context.Context, req *models.UpdatePatchRequest) (rowsAffected int64, err error) {

	var (
		set   = " SET "
		ind   = 0
		query string
	)

	if len(req.Fields) == 0 {
		err = errors.New("no updates provided")
		return
	}

	req.Fields["id"] = req.Id

	for key := range req.Fields {
		set += fmt.Sprintf(" %s = :%s ", key, key)
		if ind != len(req.Fields)-1 {
			set += ", "
		}
		ind++
	}

	query = `
		UPDATE
			"book"
	` + set + ` , updated_at = now()
		WHERE
			id = :id
	`

	query, args := helper.ReplaceQueryParams(query, req.Fields)

	result, err := u.db.Exec(ctx, query, args...)
	if err != nil {
		return
	}

	return result.RowsAffected(), err
}

func (u *BookRepo) Delete(ctx context.Context, req *book_service.BookPK) error {
	query := `DELETE FROM "book" WHERE "id" = $1`

	_, err := u.db.Exec(ctx, query, req.Id)
	if err != nil {
		return err
	}

	return nil
}
