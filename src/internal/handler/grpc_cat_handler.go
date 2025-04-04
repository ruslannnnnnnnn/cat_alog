package handler

import (
	pb "cat_alog/src/api/grpc"
	"cat_alog/src/internal/model"
	"cat_alog/src/internal/service"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type GrpcCatHandler struct {
	catService service.CatService
	pb.UnimplementedCatServiceServer
}

func NewGrpcCatHandler(catService service.CatService) *GrpcCatHandler {
	return &GrpcCatHandler{catService: catService}
}

func (g *GrpcCatHandler) GetCatById(ctx context.Context, request *pb.GetCatByIdRequest) (*pb.GetCatByIdResponse, error) {
	cat, err := g.catService.GetById(request.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.GetCatByIdResponse{
		Cat: &pb.Cat{
			Id:          cat.Id,
			Name:        cat.Name,
			DateOfBirth: cat.DateOfBirth.String(),
			ImageUrl:    cat.ImageUrl,
		},
	}, nil
}

func (g *GrpcCatHandler) GetAllCats(ctx context.Context, request *pb.GetAllCatsRequest) (*pb.GetAllCatsResponse, error) {
	cats, err := g.catService.GetAll(request.GetPage(), request.GetPerPage())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	var pbCats []*pb.Cat
	for _, cat := range cats {
		pbCats = append(pbCats, &pb.Cat{Id: cat.Id, Name: cat.Name, DateOfBirth: cat.DateOfBirth.String(), ImageUrl: cat.ImageUrl})
	}
	return &pb.GetAllCatsResponse{Cats: pbCats}, nil
}

func (g *GrpcCatHandler) CreateCat(ctx context.Context, request *pb.CreateCatRequest) (*pb.CreateCatResponse, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate id")
	}
	dateOfBirth, err := time.Parse("2006-01-02", request.DateOfBirth)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	cat := model.Cat{
		Id:          id.String(),
		Name:        request.Name,
		DateOfBirth: dateOfBirth,
		ImageUrl:    request.ImageUrl,
	}
	err = g.catService.Create(&cat)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &pb.CreateCatResponse{Id: id.String()}, nil
}
