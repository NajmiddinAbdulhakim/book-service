package storage

import (
	"context"

	pb "github.com/NajmiddinAbdulhakim/ude/book-service/genproto"
)

type BookStorageI interface {
	Create(ctx context.Context, b *pb.CreateBookReq) (*pb.BookRes, error)
	GetById(ctx context.Context, id string) (*pb.BookRes, error)
	Update(ctx context.Context, b *pb.UpdateBookReq) (bool, error)
	Delete(ctx context.Context, id string) (bool, error)
}
