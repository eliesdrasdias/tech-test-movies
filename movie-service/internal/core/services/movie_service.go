package services

import (
	"context"
	"tech-test-movies/movie-service/internal/core/domain"
	"tech-test-movies/movie-service/internal/core/ports"
)

// movieService implementa a interface MovieService
type movieService struct {
	repo ports.MovieRepository
}

// NewMovieService cria um novo MovieService
func NewMovieService(repo ports.MovieRepository) ports.MovieService {
	return &movieService{
		repo: repo,
	}
}

func (s *movieService) GetMovie(ctx context.Context, id int64) (*domain.Movie, error) {
	return s.repo.Get(ctx, id)
}

func (s *movieService) ListMovies(ctx context.Context) ([]*domain.Movie, error) {
	return s.repo.List(ctx)
}

func (s *movieService) CreateMovie(ctx context.Context, movie *domain.Movie) error {
	if movie.Title == "" {
		return context.DeadlineExceeded // Simulando um erro de validação
	}
	return s.repo.Create(ctx, movie)
}

func (s *movieService) DeleteMovie(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
