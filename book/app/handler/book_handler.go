package handler

import (
	"context"
	"errors"
	"sync"

	"github.com/surya-b21/library-management-app/book/app/model"
	"github.com/surya-b21/library-management-app/book/app/pb"
	"github.com/surya-b21/library-management-app/book/app/repo"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	repo *repo.BookRepository
	pb.UnimplementedBookServiceServer
}

func NewHandler(repo *repo.BookRepository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) Create(ctx context.Context, req *pb.BookRequest) (*pb.Book, error) {
	if req.Title == "" {
		return nil, status.Errorf(codes.InvalidArgument, "title cannot empty")
	}

	if req.Pages == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "pages cannot empty")
	}

	if req.AuthorId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "author cannot empty")
	}

	if req.CategoryId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "category cannot empty")
	}

	authorId := uuid.MustParse(req.AuthorId)
	categoryId := uuid.MustParse(req.CategoryId)
	book := model.BookAPI{
		Title:      &req.Title,
		Pages:      &req.Pages,
		AuthorID:   &authorId,
		CategoryID: &categoryId,
		Stock:      &req.Stock,
	}

	result := h.repo.Create(ctx, book)

	return &pb.Book{
		Id:         result.ID.String(),
		Title:      *result.Title,
		Pages:      *result.Pages,
		AuthorId:   result.AuthorID.String(),
		CategoryId: result.CategoryID.String(),
		Stock:      *result.Stock,
	}, nil
}

func (h *Handler) Get(ctx context.Context, req *pb.BookParam) (*pb.BookList, error) {
	result := h.repo.GetAll(ctx, req)

	bookList := make([]*pb.Book, len(result))

	for i, v := range result {
		book := pb.Book{
			Id:         v.ID.String(),
			Title:      *v.Title,
			Pages:      *v.Pages,
			AuthorId:   v.AuthorID.String(),
			CategoryId: v.CategoryID.String(),
			Stock:      *v.Stock,
		}

		bookList[i] = &book
	}

	return &pb.BookList{
		Book: bookList,
	}, nil
}

func (h *Handler) GetOne(ctx context.Context, req *pb.BookId) (*pb.Book, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot empty")
	}

	result := h.repo.Get(ctx, req.Id)

	return &pb.Book{
		Id:         result.ID.String(),
		Title:      *result.Title,
		Pages:      *result.Pages,
		AuthorId:   result.AuthorID.String(),
		CategoryId: result.CategoryID.String(),
		Stock:      *result.Stock,
	}, nil
}

func (h *Handler) Update(ctx context.Context, req *pb.Book) (*pb.Book, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot empty")
	}

	authorId := uuid.MustParse(req.AuthorId)
	categoryId := uuid.MustParse(req.CategoryId)
	bookApi := model.BookAPI{
		Title:      &req.Title,
		Pages:      &req.Pages,
		AuthorID:   &authorId,
		CategoryID: &categoryId,
		Stock:      &req.Stock,
	}

	result := h.repo.Update(ctx, bookApi, req.Id)

	return &pb.Book{
		Id:         result.ID.String(),
		Title:      *result.Title,
		Pages:      *result.Pages,
		AuthorId:   result.AuthorID.String(),
		CategoryId: result.CategoryID.String(),
		Stock:      *result.Stock,
	}, nil
}

func (h *Handler) Delete(ctx context.Context, req *pb.BookId) (*emptypb.Empty, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot empty")
	}

	h.repo.Delete(ctx, req.Id)

	return &emptypb.Empty{}, nil
}

func (h *Handler) Borrow(ctx context.Context, req *pb.BookId) (*emptypb.Empty, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot empty")
	}

	// handle race condition
	var mtx sync.Mutex

	mtx.Lock()
	response := h.repo.Borrow(ctx, req.Id)
	mtx.Unlock()

	if response != nil {
		return nil, errors.New(*response)
	}

	return &emptypb.Empty{}, nil
}

func (h *Handler) Return(ctx context.Context, req *pb.BookId) (*emptypb.Empty, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot empty")
	}

	var mtx sync.Mutex

	mtx.Lock()
	h.repo.Return(ctx, req.Id)
	mtx.Unlock()

	return &emptypb.Empty{}, nil
}
