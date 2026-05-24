package grpc

import (
	"context"

	"tech-test-movies/movie-service/internal/core/domain"
	"tech-test-movies/movie-service/internal/core/ports"
	"tech-test-movies/pb"
)

// grpcServer implementa a interface MovieService
type grpcServer struct {
	pb.UnimplementedMovieServiceServer
	service ports.MovieService
}

func NewGrpcServer(service ports.MovieService) *grpcServer {
	return &grpcServer{
		service: service,
	}
}

func (s *grpcServer) GetMovie(ctx context.Context, req *pb.GetMovieRequest) (*pb.GetMovieResponse, error) {
	movie, err := s.service.GetMovie(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetMovieResponse{
		Movie: &pb.Movie{
			Id:    movie.ID,
			Title: movie.Title,
			Year:  movie.Year,
		},
	}, nil
}

func (s *grpcServer) ListMovies(ctx context.Context, req *pb.ListMoviesRequest) (*pb.ListMoviesResponse, error) {
	movies, err := s.service.ListMovies(ctx)
	if err != nil {
		return nil, err
	}

	var pbMovies []*pb.Movie
	for _, m := range movies {
		pbMovies = append(pbMovies, &pb.Movie{
			Id:    m.ID,
			Title: m.Title,
			Year:  m.Year,
		})
	}

	return &pb.ListMoviesResponse{
		Movies: pbMovies,
	}, nil
}

func (s *grpcServer) CreateMovie(ctx context.Context, req *pb.CreateMovieRequest) (*pb.CreateMovieResponse, error) {
	movie := &domain.Movie{
		ID:    req.GetId(),
		Title: req.GetTitle(),
		Year:  req.GetYear(),
	}

	err := s.service.CreateMovie(ctx, movie)
	if err != nil {
		return nil, err
	}

	return &pb.CreateMovieResponse{
		Movie: &pb.Movie{
			Id:    movie.ID,
			Title: movie.Title,
			Year:  movie.Year,
		},
	}, nil
}

func (s *grpcServer) DeleteMovie(ctx context.Context, req *pb.DeleteMovieRequest) (*pb.DeleteMovieResponse, error) {
	err := s.service.DeleteMovie(ctx, req.GetId())
	if err != nil {
		return &pb.DeleteMovieResponse{Success: false}, err
	}

	return &pb.DeleteMovieResponse{Success: true}, nil
}
