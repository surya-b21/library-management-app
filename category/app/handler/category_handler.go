package handler

import (
	"context"

	"github.com/surya-b21/library-management-app/category/app/model"
	"github.com/surya-b21/library-management-app/category/app/pb"
	"github.com/surya-b21/library-management-app/category/app/repo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	repo *repo.CategoryRepository
	pb.UnimplementedCategoryServiceServer
}

func NewHandler(repo *repo.CategoryRepository) *Handler {
	return &Handler{
		repo: repo,
	}
}

func (h *Handler) Create(ctx context.Context, req *pb.CategoryRequest) (*pb.Category, error) {
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name cannot empty")
	}

	category := model.CategoryAPI{
		Name: &req.Name,
	}

	result := h.repo.Create(ctx, category)

	return &pb.Category{
		Id:   result.ID.String(),
		Name: *result.Name,
	}, nil
}

func (h *Handler) Get(ctx context.Context, req *emptypb.Empty) (*pb.CategoryList, error) {
	result := h.repo.GetAll(ctx)

	categoryList := make([]*pb.Category, len(result))

	for i, v := range result {
		category := pb.Category{
			Id:   v.ID.String(),
			Name: *v.Name,
		}

		categoryList[i] = &category
	}

	return &pb.CategoryList{
		Category: categoryList,
	}, nil
}

func (h *Handler) GetOne(ctx context.Context, req *pb.CategoryId) (*pb.Category, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot empty")
	}

	result := h.repo.Get(ctx, req.Id)

	return &pb.Category{
		Id:   result.ID.String(),
		Name: *result.Name,
	}, nil
}

func (h *Handler) Update(ctx context.Context, req *pb.CategoryUpdate) (*pb.Category, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot empty")
	}

	category := model.CategoryAPI{
		Name: &req.Name,
	}

	result := h.repo.Create(ctx, category)

	return &pb.Category{
		Id:   result.ID.String(),
		Name: *result.Name,
	}, nil
}

func (h *Handler) Delete(ctx context.Context, req *pb.CategoryId) (*emptypb.Empty, error) {
	if req.Id == "" {
		return nil, status.Errorf(codes.InvalidArgument, "id cannot empty")
	}

	h.repo.Delete(ctx, req.Id)

	return &emptypb.Empty{}, nil
}
