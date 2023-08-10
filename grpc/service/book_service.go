package service

import (
	"book/config"
	"book/genproto/book_service"
	"book/grpc/client"
	"book/models"
	"book/pkg/logger"
	"book/storage"

	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookService struct {
	cfg      config.Config
	log      logger.LoggerI
	strg     storage.StorageI
	services client.ServiceManagerI
	book_service.UnimplementedBookServiceServer
}

func NewBookService(cfg config.Config, log logger.LoggerI, strg storage.StorageI, srvs client.ServiceManagerI) *BookService {
	return &BookService{
		cfg:      cfg,
		log:      log,
		strg:     strg,
		services: srvs,
	}
}

func (i *BookService) Create(ctx context.Context, req *book_service.CreateBook) (resp *book_service.Book, err error) {

	i.log.Info("---CreateBook------>", logger.Any("req", req))

	pKey, err := i.strg.Book().Create(ctx, req)
	if err != nil {
		i.log.Error("!!!CreateBook->Book->Create--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = i.strg.Book().GetByPKey(ctx, pKey)
	if err != nil {
		i.log.Error("!!!GetByPKeyBook->Book->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *BookService) GetByID(ctx context.Context, req *book_service.BookPK) (resp *book_service.Book, err error) {

	i.log.Info("---GetBookByID------>", logger.Any("req", req))

	resp, err = i.strg.Book().GetByPKey(ctx, req)
	if err != nil {
		i.log.Error("!!!GetBookByID->Book->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *BookService) GetBookByTitle(ctx context.Context, req *book_service.BookByTitle) (resp *book_service.Book, err error) {

	i.log.Info("---GetBookByID------>", logger.Any("req", req))

	resp, err = i.strg.Book().GetBookByTitle(ctx, req)
	if err != nil {
		i.log.Error("!!!GetBookByID->Book->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *BookService) GetList(ctx context.Context, req *book_service.BookListRequest) (resp *book_service.BookListResponse, err error) {

	i.log.Info("---GetBooks------>", logger.Any("req", req))

	resp, err = i.strg.Book().GetAll(ctx, req)
	if err != nil {
		i.log.Error("!!!GetBooks->Book->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return
}

func (i *BookService) Update(ctx context.Context, req *book_service.UpdateBook) (resp *book_service.Book, err error) {

	i.log.Info("---UpdateBook------>", logger.Any("req", req))

	rowsAffected, err := i.strg.Book().Update(ctx, req)

	if err != nil {
		i.log.Error("!!!UpdateBook--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Book().GetByPKey(ctx, &book_service.BookPK{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetBook->Book->Get--->", logger.Error(err))
		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *BookService) UpdatePatch(ctx context.Context, req *book_service.UpdatePatchBook) (resp *book_service.Book, err error) {

	i.log.Info("---UpdatePatchBook------>", logger.Any("req", req))

	updatePatchModel := models.UpdatePatchRequest{
		Id:     req.GetId(),
		Fields: req.GetFields().AsMap(),
	}

	rowsAffected, err := i.strg.Book().UpdatePatch(ctx, &updatePatchModel)

	if err != nil {
		i.log.Error("!!!UpdatePatchBook--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if rowsAffected <= 0 {
		return nil, status.Error(codes.InvalidArgument, "no rows were affected")
	}

	resp, err = i.strg.Book().GetByPKey(ctx, &book_service.BookPK{Id: req.Id})
	if err != nil {
		i.log.Error("!!!GetBook->Book->Get--->", logger.Error(err))

		return nil, status.Error(codes.NotFound, err.Error())
	}

	return resp, err
}

func (i *BookService) Delete(ctx context.Context, req *book_service.BookPK) (resp *empty.Empty, err error) {

	i.log.Info("---DeleteBook------>", logger.Any("req", req))

	err = i.strg.Book().Delete(ctx, req)
	if err != nil {
		i.log.Error("!!!DeleteBook->Book->Get--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return &empty.Empty{}, nil
}
