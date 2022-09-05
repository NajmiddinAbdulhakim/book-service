package service

import (
	"context"
	"log"

	pb "github.com/NajmiddinAbdulhakim/ude/book-service/genproto"
	"github.com/NajmiddinAbdulhakim/ude/book-service/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type BookService struct {
	repo storage.IStorage
}

func NewBookService(db *sqlx.DB) *BookService {
	return &BookService{
		repo: storage.NewStoragePg(db),
	}
}

func (s *BookService) Create(ctx context.Context, req *pb.CreateBookReq) (*pb.BookRes, error) {
	res, err := s.repo.Book().Create(ctx, req)
	if err != nil {
		log.Println(`Failed while creating book:`, err)
		return nil, status.Error(codes.InvalidArgument, `Failed while creating book`)
	}

	return res, nil
}

func (s *BookService) GetById(ctx context.Context, req *pb.BookByIdReq) (*pb.BookRes, error ) {
	res, err := s.repo.Book().GetById(ctx, req.Id)
	if err != nil {
		log.Println(`Failed while getting book by id:`, err)
		return nil, status.Error(codes.InvalidArgument, `Failed while getting book by id`)
	}
	return res, nil
}	

func (s *BookService) Update(ctx context.Context, req *pb.UpdateBookReq) (*pb.BoolRes, error) {
	res, err := s.repo.Book().Update(ctx, req)
	if err != nil { 
		log.Println(`Failed while updating book by id:`, err)
		return &pb.BoolRes{Success: res}, status.Error(codes.InvalidArgument, `Failed while updating book by id`)
	}

	return &pb.BoolRes{Success: res}, nil
}

func (s *BookService) Delete(ctx context.Context, req *pb.BookByIdReq) (*pb.BoolRes, error) {
	res, err := s.repo.Book().Delete(ctx, req.Id)
	if err != nil { 
		log.Println(`Failed while deleting book by id:`, err)
		return &pb.BoolRes{Success: res}, status.Error(codes.InvalidArgument, `Failed while deleting book by id`)
	}

	return &pb.BoolRes{Success: res}, nil
}
