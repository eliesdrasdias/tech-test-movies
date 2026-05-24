package ports

import (
	"context"
	"tech-test-movies/movie-service/internal/core/domain"
)

type MovieService interface {
	GetMovie(ctx context.Context, id int64) (*domain.Movie, error)
	ListMovies(ctx context.Context) ([]*domain.Movie, error)
	CreateMovie(ctx context.Context, movie *domain.Movie) error
	DeleteMovie(ctx context.Context, id int64) error
}
